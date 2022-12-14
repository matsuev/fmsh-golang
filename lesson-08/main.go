package main

import (
	"fmsh-golang/lesson-08/csvstorage"
	"fmt"
	"log"
)

type DataIterator interface {
	Scan() bool
	LineMap() map[string]string
	LineArray() []string
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
		lineArray := store.LineMap()
		if lineArray["email"] == "ivan@example.org" {
			break
		}
		fmt.Printf("%#v\n", lineArray)
	}

	// hjerkhkjshgjkshekltjhewr
	fmt.Println("Hello, Ivan!!!")

	for store.Scan() {
		lineArray := store.LineMap()
		fmt.Printf("%#v\n", lineArray)
	}

}
