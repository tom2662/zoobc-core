package constant

import "time"

const (
	// SnapshotGenerationTimeout maximum time, in seconds, allowed for a node to generate a snapshot
	// @iltoga reduce to 1 for testing locally
	MainchainSnapshotGenerationTimeout time.Duration = 10 * time.Minute // 10 minutes before including in spine block
	// MainchainSnapshotInterval interval in mainchain blocks between snapshots
	// @iltoga reduce to 5 for testing locally
	MainchainSnapshotInterval uint32 = 720 // 720 mainchain blocks (= MinRollbackHeight)
	// @iltoga reduce to 1 for testing locally
	SnapshotChunkSize int = int(100 * 1024) // 10 KB
	// DownloadSnapshotNumberOfRetries number of times to retry downloading failed snapshot file chunks from other peers
	DownloadSnapshotNumberOfRetries uint32 = 3
)