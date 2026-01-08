package main

import (
	"fmt"
	"sync"
)

type SynchronizedQueue struct {
	items []interface{}
	mutex sync.Mutex
}

func NewSyncQueue() *SynchronizedQueue {
	return &SynchronizedQueue{items: make([]interface{}, 0)}
}

func (q *SynchronizedQueue) Enqueue(data interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, data)
}

func (q *SynchronizedQueue) Dequeue() (interface{}, bool) {
	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *SynchronizedQueue) IsEmpty() (result bool) {
	if len(q.items) == 0 {
		result = true
	}
	return
}

func (q *SynchronizedQueue) Size() (length int) {
	if len(q.items) > 0 {
		length = len(q.items)
	}
	return
}

func (q *SynchronizedQueue) Display() {
	fmt.Println()
	for _, val := range q.items {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func main() {
	q := NewSyncQueue()
	q.Enqueue(3.14)
	q.Enqueue("hello world")
	q.Enqueue(true)
	q.Enqueue(111)

	q.Display()

	data, success := q.Dequeue()
	if !success {
		fmt.Println("real ayam 😅")
	}
	fmt.Printf("data: %v\n", data)

	q.Display()

	fmt.Printf("is empty: %v\n", q.IsEmpty())

	fmt.Printf("size: %d\n", q.Size())
}
