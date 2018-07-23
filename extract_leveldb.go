package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
)

func main() {
	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	db, err := leveldb.OpenFile("~/geth/chaindata/", nil)
	iter := db.NewIterator(nil, nil)
	fmt.Print(err)
	for iter.Next() {
	// Remember that the contents of the returned slice should not be modified, and
	// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Println(key)
		fmt.Println(value)
	}
	iter.Release()
}
