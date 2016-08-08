	"os"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/metrics"
	Logger(value xlog.Logger) DatabaseOptions
	GetLogger() xlog.Logger
	MetricsScope(value xmetrics.Scope) DatabaseOptions
	GetMetricsScope() xmetrics.Scope
	// CommitLogFlushSize sets the commit log flush size and returns a new DatabaseOptions
	CommitLogFlushSize(value int) DatabaseOptions

	// GetCommitLogFlushSize returns the commit log flush size
	GetCommitLogFlushSize() int

	// CommitLogFlushInterval sets the commit log flush interval and returns a new DatabaseOptions
	CommitLogFlushInterval(value time.Duration) DatabaseOptions

	// GetCommitLogFlushInterval returns the commit log flush interval
	GetCommitLogFlushInterval() time.Duration

	// CommitLogBacklogQueueSize sets the commit log backlog queue size and returns a new DatabaseOptions
	CommitLogBacklogQueueSize(value int) DatabaseOptions

	// GetCommitLogBacklogQueueSize returns the commit log backlog queue size
	GetCommitLogBacklogQueueSize() int

	// CommitLogStrategy sets the commit log strategy and returns a new DatabaseOptions
	CommitLogStrategy(value CommitLogStrategy) DatabaseOptions

	// GetCommitLogStrategy returns the commit log strategy
	GetCommitLogStrategy() CommitLogStrategy

	// SegmentReaderPool sets the segment reader pool.
	SegmentReaderPool(value SegmentReaderPool) DatabaseOptions

	// GetSegmentReaderPool returns the segment reader pool.
	GetSegmentReaderPool() SegmentReaderPool

	// FileWriterOptions sets the file writer options.
	FileWriterOptions(value FileWriterOptions) DatabaseOptions

	// GetFileWriterOptions returns the file writer options.
	GetFileWriterOptions() FileWriterOptions


	// NewPersistenceManagerFn sets the function for creating a new persistence manager.
	NewPersistenceManagerFn(value NewPersistenceManagerFn) DatabaseOptions

	// GetNewPersistenceManagerFn returns the function for creating a new persistence manager.
	GetNewPersistenceManagerFn() NewPersistenceManagerFn

	// WriterBufferSize sets the buffer size for writing TSDB files.
	WriterBufferSize(value int) DatabaseOptions

	// GetWriterBufferSize returns the buffer size for writing TSDB files.
	GetWriterBufferSize() int

	// ReaderBufferSize sets the buffer size for reading TSDB files.
	ReaderBufferSize(value int) DatabaseOptions

	// GetReaderBufferSize returns the buffer size for reading TSDB files.
	GetReaderBufferSize() int
	Logger(value xlog.Logger) ClientOptions
	GetLogger() xlog.Logger
	MetricsScope(value xmetrics.Scope) ClientOptions
	GetMetricsScope() xmetrics.Scope
	// ClusterConnectConsistencyLevel sets the clusterConnectConsistencyLevel and returns a new ClientOptions
	ClusterConnectConsistencyLevel(value ConsistencyLevel) ClientOptions

	// GetClusterConnectConsistencyLevel returns the clusterConnectConsistencyLevel
	GetClusterConnectConsistencyLevel() ConsistencyLevel

// TopologyTypeOptions is a set of static topology type options
type TopologyTypeOptions interface {
	// ShardScheme sets the shardScheme and returns a new TopologyTypeOptions
	ShardScheme(value ShardScheme) TopologyTypeOptions
	// Replicas sets the replicas and returns a new TopologyTypeOptions
	Replicas(value int) TopologyTypeOptions
	// HostShardSets sets the hostShardSets and returns a new TopologyTypeOptions
	HostShardSets(value []HostShardSet) TopologyTypeOptions

// FileWriterOptions is a set of file writing options for a file writer
type FileWriterOptions interface {
	// NewFileMode sets the new file mode.
	NewFileMode(value os.FileMode) FileWriterOptions

	// GetNewFileMode returns the new file mode.
	GetNewFileMode() os.FileMode

	// NewDirectoryMode sets the new directory mode.
	NewDirectoryMode(value os.FileMode) FileWriterOptions

	// GetNewDirectoryMode returns the new directory mode.
	GetNewDirectoryMode() os.FileMode
}