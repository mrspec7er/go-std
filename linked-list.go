package main

import (
	"container/list"
	"fmt"
)

func main()  {
	queue := list.New()

	queue.PushBack("First")
	queue.PushBack(2)
	queue.PushBack(3.0)

	queue.PushFront("Init")
	queue.PushFront("Start")

	for i := queue.Front(); i != nil; i =  i.Next() {
		fmt.Println(i.Value)
	}
}