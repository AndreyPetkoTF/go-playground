package main

import "errors"

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{Elements: make(map[string]struct{})}
}

func (s *Set) Add(element string) {
	s.Elements[element] = struct{}{}
}

func (s *Set) Delete(element string) error {
	if !s.Exists(element) {
		return errors.New("no element")
	}

	delete(s.Elements, element)

	return nil
}

func (s *Set) Exists(element string) bool {
	_, exists := s.Elements[element]

	return exists
}

func main() {

}
