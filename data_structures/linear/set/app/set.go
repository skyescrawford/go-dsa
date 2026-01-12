package main

import "fmt"

type Set struct {
	items map[string]struct{}
}

func New() *Set {
	items := make(map[string]struct{})
	return &Set{items: items}
}

func (s *Set) Add(name string) bool {
	_, ok := s.items[name]
	if !ok {
		s.items[name] = struct{}{}
		return true
	}
	return false
}

func (s *Set) Delete(name string) bool {
	_, ok := s.items[name]
	if !ok {
		return ok
	}
	delete(s.items, name)
	return ok
}

func (s *Set) Display() {
	for item := range s.items {
		fmt.Printf("%s ", item)
	}
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set) Clear() {
	s.items = make(map[string]struct{})
}
