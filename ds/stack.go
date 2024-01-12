package ds

import "reflect"

func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

type Stack[T any] struct {
	size uint64
	data []T
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, s.size)
}

func (s *Stack[T]) Peek() T {
	if len(s.data) == 0 {
		return *new(T)
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		return *new(T)
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element
}

func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
}

func (s *Stack[T]) Size() int {
	var size int = 0
	for _, item := range s.data {
		if !IsZeroOfUnderlyingType(item) {
			size++
		}
	}
	return size
}

func NewStack[T any](size uint64) Stack[T] {
	return Stack[T]{
		size: size,
		data: make([]T, size),
	}
}
