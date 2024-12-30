package main

import (
	"container/list"
	"fmt"
)

func main() {
	numbers := list.New()
	el := numbers.PushBack(1)

	numbers.PushFront(0)
	numbers.PushBack(2)
	//012

	for e := numbers.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("--------")
	numbers.Remove(el)
	for e := numbers.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
