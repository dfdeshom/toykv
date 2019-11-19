package main

import (
	"fmt"

	//kv "github.com/dfdeshom/toykv"
	"github.com/dfdeshom/toykv"
	"github.com/google/btree"
)

type Entry toykv.Entry

type MemTable struct {
	btree btree.BTree
}

func (m MemTable) Get(key string) (string, error) {
	k := Entry{Key: key}
	item := m.btree.Get(k)
	if item != nil {
		return item.(Entry).Value, nil
	}
	return "", nil
}

func (a Entry) Less(b btree.Item) bool {
	return a.Key < b.(Entry).Key
}

type EntryIterator func(i toykv.Entry) bool

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
	fmt.Printf("%#v", tr.Get(e))

	tr.Ascend(func(i btree.Item) bool {
		fmt.Printf("\n%#v\n", i.(Entry))
		return true
	})
}
