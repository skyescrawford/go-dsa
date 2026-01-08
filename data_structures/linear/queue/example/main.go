package main

import (
	"fmt"
	"sync"
	"time"
)

type Customer struct {
	Id   int
	Name string
}

type TicketQueue struct {
	customers []*Customer
	mutex     sync.Mutex
}

func NewQueue() *TicketQueue {
	return &TicketQueue{customers: make([]*Customer, 0)}
}

func (q *TicketQueue) Enqueue(c *Customer) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.customers = append(q.customers, c)
}

func (q *TicketQueue) Dequeue() (*Customer, bool) {
	if len(q.customers) == 0 {
		return nil, false
	}
	customer := q.customers[0]
	q.customers = q.customers[1:]

	return customer, true
}

func (q *TicketQueue) StartTicketIssue() {
	for {
		q.mutex.Lock()
		if len(q.customers) == 0 {
			q.mutex.Unlock()
			time.Sleep(time.Second)
			continue
		}

		customer := *q.customers[0]
		q.customers = q.customers[1:]
		q.mutex.Unlock()

		fmt.Printf("Issuing ticket to customer %s (ID: %d)\n", customer.Name, customer.Id)

		time.Sleep(2 * time.Second)
	}
}

func main() {
	queue := NewQueue()

	queue.Enqueue(&Customer{Id: 1, Name: "Skyes"})
	queue.Enqueue(&Customer{Id: 2, Name: "Sykes"})
	queue.Enqueue(&Customer{Id: 3, Name: "Silver"})
	queue.Enqueue(&Customer{Id: 4, Name: "Oliver"})
	queue.Enqueue(&Customer{Id: 5, Name: "Ryouta"})

	queue.StartTicketIssue()
}
