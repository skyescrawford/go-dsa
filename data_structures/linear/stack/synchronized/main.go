package main

import (
	"fmt"
	"os"
	"sync"
)

type HistoryStack struct {
	items []string
	mutex sync.Mutex
}

func NewHistory() *HistoryStack {
	return &HistoryStack{items: make([]string, 0)}
}

func (h *HistoryStack) Push(url string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.items = append(h.items, url)
}

func (h *HistoryStack) Pop() (string, bool) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if len(h.items) == 0 {
		return "", false
	}

	lastIndex := len(h.items) - 1
	lastItem := h.items[lastIndex]
	h.items = h.items[:lastIndex]

	return lastItem, true
}

func (h *HistoryStack) Display() {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	lastIndex := len(h.items) - 1
	for i := lastIndex; i >= 0; i-- {
		fmt.Printf("%s\n", h.items[i])
	}
}

func (h *HistoryStack) Size() int {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	return len(h.items)
}

func (h *HistoryStack) IsEmpty() bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	return len(h.items) == 0
}

func main() {
	history := NewHistory()

	history.Push("https://google.com")
	history.Push("https://google.com?query=stack-tutorial")
	history.Push("https://www.geeksforgeeks.org/dsa/introduction-to-stack-data-structure-and-algorithm-tutorials/")

	history.Display()

	prevUrl, success := history.Pop()
	if !success {
		fmt.Println("you are in the initial page")
		os.Exit(1)
	}

	fmt.Printf("previous url: %s\n", prevUrl)
	fmt.Printf("size: %d\n", history.Size())
}
