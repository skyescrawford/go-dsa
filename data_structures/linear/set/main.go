package main

import "fmt"

// Set adalah linear data struktur yang menyimpan data unik, ibaratkan array yg tiap elemen nya ngga ada yg duplikat

type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	e := make(map[string]struct{})
	return &Set{elements: e}
}

func (s *Set) contains(data string) bool {
	_, ok := s.elements[data]
	return ok
}

func (s *Set) add(data string) {
	if s.contains(data) {
		return
	}
	s.elements[data] = struct{}{}
}

func (s *Set) delete(data string) {
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

	foo.add("foo")
	foo.add("bar")

	// foo.delete(1)

	fmt.Printf("size foo: %d\n", foo.size())
	// foo.clear()
	fmt.Printf("foo: %v\n", foo)
	fmt.Printf("is foo empty: %v\n", foo.isEmpty())

	bar.add("john")
	bar.add("doe")

	fmt.Printf("foo: %v\n", foo.elements)
	fmt.Printf("bar: %v\n", bar.elements)

	fmt.Printf("intersection: %v\n", foo.intersect(bar))
	fmt.Printf("union: %v\n", foo.union(bar))

	foo.clear()

	fmt.Printf("foo: %v\n", foo.elements)
}
