// Copyright (c) 2012, Robert Dinu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license which can be found in the LICENSE file.

// Package stack implements a thread-safe stack data structure.
package stack

import "sync"

// Stack represents a stack data structure.
type Stack struct {
	front *element
	len   int
	mu    sync.RWMutex
}

type element struct {
	next  *element
	value interface{}
}

// Clear removes all objects from the stack.
func (s *Stack) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.front = nil
	s.len = 0
}

// Len returns the number of elements in the stack.
func (s *Stack) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.len
}

// Peek returns the value of the element at the top
// of the stack without removing it.
func (s *Stack) Peek() interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.len < 1 {
		return nil
	}
	return s.front.value
}

// Pop removes and returns the value of the element at the top of the stack.
func (s *Stack) Pop() (value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.len < 1 {
		return
	}
	value = s.front.value
	s.front = s.front.next
	s.len--
	return
}

// Push inserts an element with the given value at the top of the stack.
func (s *Stack) Push(value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.front = &element{
		next:  s.front,
		value: value,
	}
	s.len++
}

// New creates and returns a new (empty) stack.
func New() *Stack {
	return new(Stack)
}
