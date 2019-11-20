package toykv

type KeyValueStore struct {
	memtable MemTable
	sstable  SSTable
}

type MemTable interface {
	Get(key string) (Entry, bool)
	Set(key string, value string)
	FlushToSSTable() error
}

type SSTable interface {
	Get(key string) (Entry, bool)
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

func (kvs KeyValueStore) Get(key string) Entry {
	//TODO: use a bloom filter to catch
	// non-existent keys

	// get key in memtable first
	// if not present, it must be
	// in the SSTable
	memtable := kvs.memtable
	res, present := memtable.Get(key)
	if present {
		return res
	}

	// else, it's in the SSTable
	sstable := kvs.sstable
	res, present = sstable.Get(key)
	if present {
		return res
	}

	return Entry{}
}

func (kvs KeyValueStore) Set(key string, value string) {
	memtable := kvs.memtable
	memtable.Set(key, value)

}
