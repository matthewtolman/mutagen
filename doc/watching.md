# Filesystem watching

Mutagen uses filesystem watching to know when it should re-scan and propagate
files. Unfortunately, the filesystem watching landscape is *extremely* varied in
terms of implementation, efficiency, and robustness. Almost every platform uses
a completely different mechanism, many of which are unreliable or non-scalable.
For example, some systems (namely macOS and Windows) provide native recursive
watching mechanisms that can monitor arbitrarily large directory hierarchies,
but their behavior when the location being watched is deleted or changed is
problematic or useless, and Windows only supports using a directory as the root
of such a recursive watch. Other systems (e.g. Linux and BSD systems) provide
mechanisms that require a watch descriptor or file descriptor to be open for
*every* file or directory being watched in a directory hierarchy, which can
quickly exhaust system quotas for directories that might be used in development
(e.g. imagine a synchronization root containing a `node_modules` directory).
Mutagen takes a pragmatic, hybrid approach to filesystem watching that attempts
to maximize responsiveness while avoiding exhaustion of system resources or
problematic behavior.

On systems that natively support recursive filesystem watching, a watch is
established on either the synchronization root itself (on macOS) or the parent
directory of the synchronization root (on Windows) and events are filtered to
only those originating from the synchronization root. Because these systems can
behave strangely if the root of a watch is deleted, a regular (but very cheap)
polling mechanism is used to ensure that the watch root hasn't been deleted or
recreated. If a change to the watch root is detected, the watch is
re-established.

On all other systems, a polling mechanism is used to avoid exhausting watch/file
descriptors.

Mutagen provides two different filesystem watching modes:

- **Portable** (Default): In this mode, Mutagen uses the most efficient watching
  implementation possible for an endpoint. If native recursive watching is
  available, it is used, otherwise poll-based watching is used.
- **Force Poll**: In this mode, Mutagen will always use its poll-based watching
  implementation, even on systems that support native recursive watching.

Active R&D is underway to improve the watching situation on systems without
native recursive watching, so please stand by. If you have any feedback on other
efficient (but portable and safe) filesystem watching designs, please feel free
to contact me or open an issue to discuss your proposal.

These modes can be specified on a per-session basis by passing the
`--watch-mode=<mode>` flag to the `create` command (where `<mode>` is `portable`
or `force-poll`) and on a default basis by including the following configuration
in `~/.mutagen.toml`:

    [watch]
    mode = "<mode>"

The polling interval (which defaults to 10 seconds) can be specified on a
per-session basis by passing the `--watch-polling-interval=<interval>` flag to
the `create` command (where `<interval>` is an integer value representing
seconds) and on a default basis by including the following configuration in
`~/.mutagen.toml`:

    [watch]
    pollingInterval = <interval>


## Windows caveats

Mutagen is currently unable to use native recursive filesystem watching on
Windows paths that are *direct* descendants of a drive (e.g. `C:\code`) due to
[rjeczalik/notify#148](https://github.com/rjeczalik/notify/issues/148). As soon
as this bug is resolved, this limitation will be lifted. In the mean time,
watching will fall back to polling for these case.