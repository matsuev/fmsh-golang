package main

import (
	"fmsh-golang/lesson-08/csvstorage"
	"fmt"
	"log"
)

// DataIterator interface
type DataIterator interface {
	Scan() bool
	GetMap() map[string]string
	GetArray() []string
	GetHeaders() []string
}

func main() {
	println("CSV Data")

	var store DataIterator
	var err error

	store, err = csvstorage.Create("./list.csv", ";")
	if err != nil {
		log.Fatalln(err)
	}

	for store.Scan() {
		line := store.GetMap()
		fmt.Printf("%#v\n", line)
	}

}
