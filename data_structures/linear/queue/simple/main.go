package main

import "fmt"

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{items: make([]interface{}, 0)}
}

// menambahkan data di akhir list
func (q *Queue) Enqueue(data interface{}) {
	q.items = append(q.items, data)
}

// mengambil data pertama di list
func (q *Queue) Dequeue() (interface{}, bool) {
	if len(q.items) == 0 {
		return nil, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// mengecek apakah queue lagi kosong ngga, kalo kosong ajak main 😅, awokawok
func (q *Queue) IsEmpty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}

// mengecek panjang queue
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := NewQueue()

	// queue.Enqueue(1)
	// queue.Enqueue("hello, workd")
	// queue.Enqueue(true)
	// queue.Enqueue(3.14)

	for _, val := range queue.items {
		fmt.Printf("%v ", val)
	}

	fmt.Println()

	data, success := queue.Dequeue()
	if !success {
		fmt.Println("ayam, ptok ptok")
	}

	fmt.Printf("data: %v\n", data)
	for _, val := range queue.items {
		fmt.Printf("%v ", val)
	}

	fmt.Println()

	fmt.Printf("is empty: %v\n", queue.IsEmpty())

	fmt.Printf("size: %d\n", queue.Size())
}
