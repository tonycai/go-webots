package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	db, err := leveldb.OpenFile("../db/webots.db", nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()
	// put data
	//err = db.Put([]byte("fizz.html"), []byte("buzz"), nil)
	//err = db.Put([]byte("fizz2.html"), []byte("buzz2"), nil)
	//err = db.Put([]byte("fizz3.html"), []byte("buzz3"), nil)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	fmt.Println("\n")

	for ok := iter.Seek([]byte("fizz2")); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	fmt.Println("\n")

	for ok := iter.First(); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	iter.Release()
	err = iter.Error()
}
