	"github.com/mutagen-io/mutagen/pkg/stream"
	// problems are the problems encountered during transition operations.
		} else if entry.Kind == EntryKind_SymbolicLink {
		entryCopy := entry.Copy(true)
	} else if entry.Kind == EntryKind_SymbolicLink {
	preemptableTemporary := stream.NewPreemptableWriter(
		temporary,
		t.cancelled,
		transitionCopyPreemptionInterval,
	)
		if copyErr == stream.ErrWritePreempted {
	created := target.Copy(false)
		} else if entry.Kind == EntryKind_SymbolicLink {
	} else if target.Kind == EntryKind_SymbolicLink {
		// fails, then record the reduced entry and continue to the next
		// transition.
		// entry is nil, then this is a no-op.