package main

//package memtable

import (
	"fmt"

	kv "github.com/dfdeshom/toykv/keyvaluestore"
	"github.com/google/btree"
)

type MemTable struct {
	btree btree.BTree
}

func (m MemTable) Get(key string) (kv.Entry, error) {
	item := m.btree.Get(key.(Item))
	if item != nil {
		return item.(Entry).Value, nil
	}
	return nil, nil
}

// type Entry struct {
// 	Key    string
// 	Value  string
// 	IsLive bool
// 	Offset int
// }

func (a Entry) Less(b btree.Item) bool {
	return a.Key < b.(Entry).Key
}

type EntryIterator func(i Entry) bool

func main() {
	btreeDegree := 32
	tr := btree.New(btreeDegree)

	e := Entry{Key: "ba",
		Value: "1212",
	}

	f := Entry{Key: "ab",
		Value: "a21233",
	}

	tr.ReplaceOrInsert(e)
	tr.ReplaceOrInsert(f)
	fmt.Printf("%#v", tr)

	//it := btree.ItemIterator
	//out := make([]btree.Item, 1)
	tr.Ascend(func(i btree.Item) bool {
		fmt.Printf("\n%#v\n", i.(Entry))
		return true
	})
}
