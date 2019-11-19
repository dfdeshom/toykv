package toykv

type KeyValueStore struct {
	memtable MemTable
	sstable  SSTable
}

type MemTable interface {
	Get(key string) (Entry, error)
	Set(key string) error
	FlushToSSTable() error
}

type SSTable interface {
	Get(key string) (Entry, error)
}

type SSTableDiskFormat struct {
	KeySize   int //varint
	Key       string
	ValueSize int //varint
	Value     string
}

type Entry struct {
	Key    string
	Value  string
	IsLive bool
	Offset int
}
