package main

import (
	"fmt"
	"lrucache_experiment/list"
	"time"
)

func main() {
	newList := list.NewList[string, string](10)
	fmt.Println(*newList)

	entry1 := &list.Entry[string, string]{
		Next:  nil,
		Back:  nil,
		Key:   "key",
		Value: "value",
		Bday:  time.Now(),
		TTL:   time.Duration(5 * 10000000000),
	}
	newList.Add(*entry1)

	fmt.Println(*newList)
	fmt.Println(newList.Len())
}
