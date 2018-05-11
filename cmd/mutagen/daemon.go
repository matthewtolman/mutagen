package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/pkg/errors"

	"github.com/havoc-io/mutagen/pkg/agent"
	"github.com/havoc-io/mutagen/cmd"
	"github.com/havoc-io/mutagen/pkg/daemon"
	"github.com/havoc-io/mutagen/pkg/process"
	"github.com/havoc-io/mutagen/pkg/rpc"
	"github.com/havoc-io/mutagen/pkg/session"
	"github.com/havoc-io/mutagen/pkg/ssh"
)

var daemonUsage = `usage: mutagen daemon [-h|--help] [-s|--stop]

Controls the lifecycle of the daemon. The default behavior of this command is to
start the daemon in the background. The command is idempotent - a daemon
instance is only created if one doesn't already exist.
`

func daemonMain(arguments []string) error {
	// Parse command line arguments.
	var run bool
	var stop bool
	flagSet := cmd.NewFlagSet("daemon", daemonUsage, nil)
	flagSet.BoolVarP(&run, "run", "r", false, "run the daemon server")
	flagSet.BoolVarP(&stop, "stop", "s", false, "stop any running daemon server")
	flagSet.ParseOrDie(arguments)

	// Check that options are sane.
	if run && stop {
		return errors.New("-r/--run with -s/--stop doesn't make sense")
	}

	// If stopping is requested, try to send a termination request.
	if stop {
		daemonClient := rpc.NewClient(daemon.NewOpener())
		stream, err := daemonClient.Invoke(daemon.MethodTerminate)
		if err != nil {
			return errors.Wrap(err, "unable to invoke daemon termination")
		}
		stream.Close()
		return nil
	}

	// Unless running (non-backgrounding) is requested, then we need to restart
	// in the background.
	if !run {
		daemonProcess := &exec.Cmd{
			Path:        process.Current.ExecutablePath,
			Args:        []string{"mutagen", "daemon", "--run"},
			SysProcAttr: daemonProcessAttributes,
		}
		if err := daemonProcess.Start(); err != nil {
			return errors.Wrap(err, "unable to fork daemon")
		}
		return nil
	}

	// TODO: Do we eventually want to encapsulate the construction of the daemon
	// RPC server into the daemon package, much like we do with endpoints? It
	// becomes a bit difficult to do cleanly. Also, I want the ability to have
	// different processes host the daemon (e.g. a GUI). In those cases, we may
	// want to add additional services that wouldn't be present in the CLI
	// daemon. So I'll leave things the way they are for now, but I'd like to
	// keep thinking about this for the future. One easy thing we could do is
	// move the daemon lock into the daemon service (and add a corresponding
	// shutdown method to the daemon service).

	// Attempt to acquire the daemon lock and defer its release. If there is a
	// crash, the lock will be released by the OS automatically, but on Windows
	// this may only happen after some unspecified period of time (though it
	// does seem to be basically instant).
	lock, err := daemon.AcquireLock()
	if err != nil {
		return errors.Wrap(err, "unable to acquire daemon lock")
	}
	defer lock.Unlock()

	// Perform housekeeping.
	agent.Housekeep()
	session.HousekeepCaches()
	session.HousekeepStaging()

	// Create the RPC server.
	server := rpc.NewServer()

	// Create and register the daemon service.
	daemonService, daemonTermination := daemon.NewService()
	server.Register(daemonService)

	// Create and regsiter the SSH service.
	sshService := ssh.NewService()
	server.Register(sshService)

	// Create the and register the session service and defer its shutdown. We
	// want to do a clean shutdown because we don't want to information
	// generated during a synchronization cycle.
	sessionService, err := session.NewService(sshService)
	if err != nil {
		return errors.Wrap(err, "unable to create session service")
	}
	server.Register(sessionService)
	defer sessionService.Shutdown()

	// Create the daemon listener and defer its closure.
	listener, err := daemon.NewListener()
	if err != nil {
		return errors.Wrap(err, "unable to create daemon listener")
	}
	defer listener.Close()

	// Serve incoming connections in a separate Goroutine, watching for serving
	// failure.
	serverTermination := make(chan error, 1)
	go func() {
		serverTermination <- server.Serve(listener)
	}()

	// Wait for termination from a signal, the server, or the daemon server. We
	// treat daemon termination as a non-error.
	signalTermination := make(chan os.Signal, 1)
	signal.Notify(signalTermination, cmd.TerminationSignals...)
	select {
	case sig := <-signalTermination:
		return errors.Errorf("terminated by signal: %s", sig)
	case <-daemonTermination:
		return nil
	case err = <-serverTermination:
		return errors.Wrap(err, "premature server termination")
	}
}
