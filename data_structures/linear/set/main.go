package main

import "fmt"

type Set struct {
	elements map[int]bool
}

func NewSet() *Set {
	e := make(map[int]bool)
	return &Set{elements: e}
}

func (s *Set) contains(data int) bool {
	_, ok := s.elements[data]
	return ok
}

func (s *Set) add(data int) {
	if s.contains(data) {
		return
	}
	s.elements[data] = true
}

func (s *Set) delete(data int) {
	if !s.contains(data) {
		return
	}
	delete(s.elements, data)
}

func (s *Set) intersect(t *Set) *Set {
	intersect := NewSet()

	for k := range s.elements {
		if t.contains(k) {
			intersect.add(k)
		}
	}

	return intersect
}

func (s *Set) union(t *Set) *Set {
	union := NewSet()
	
	for k := range s.elements {
		union.add(k)
	}

	for k := range t.elements {
		union.add(k)
	}

	return union
}

func (s *Set) size() int {
	return len(s.elements)
}

func (s *Set) clear() {
	s.elements = nil
}

func (s *Set) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}

func main() {
	foo := NewSet()
	bar := NewSet()

	foo.add(1)
	foo.add(1)
	foo.add(2)

	// foo.delete(1)

	fmt.Printf("size foo: %d\n", foo.size())
	// foo.clear()
	fmt.Printf("foo: %v\n", foo)
	fmt.Printf("is foo empty: %v\n", foo.isEmpty())

	bar.add(1)
	bar.add(3)
	bar.add(4)

	fmt.Printf("foo: %v\n", foo.elements)
	fmt.Printf("bar: %v\n", bar.elements)

	fmt.Printf("intersection: %v\n", foo.intersect(bar))
	fmt.Printf("union: %v\n", foo.union(bar))

	foo.clear()

	fmt.Printf("foo: %v\n", foo.elements)
}