// Copyright (c) 2012, Robert Dinu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license which can be found in the LICENSE file.

// Package table implements a thread-safe table data structure.
package table

import "sync"

// Table represents a collection that associates a row
// key and a column key, with a single value.
type Table struct {
	len int
	m   map[interface{}]map[interface{}]interface{}
	mu  sync.RWMutex
}

// Clear removes all mappings from the table.
func (t *Table) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	for row, _ := range t.m {
		delete(t.m, row)
		t.len = 0
	}
}

// Column returns a map that associates the row key with the value for
// all the mappings in the table that have the given column key.
// If there are no such mappings in the table, returns an empty map.
func (t *Table) Column(columnKey interface{}) map[interface{}]interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	c := make(map[interface{}]interface{})
	for rowKey, column := range t.m {
		if value, ok := column[columnKey]; ok {
			c[rowKey] = value
		}
	}
	return c
}

// Contains returns true if a mapping with the given
// row and column keys exists in the table.
func (t *Table) Contains(rowKey, columnKey interface{}) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	_, y := t.m[rowKey][columnKey]
	return y
}

// ContainsColumn returns true if a mapping with
// the given column key exists in the table.
func (t *Table) ContainsColumn(columnKey interface{}) (y bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, column := range t.m {
		if _, y = column[columnKey]; y {
			break
		}
	}
	return
}

// ContainsRow returns true if a mapping with
// the given row key exists in the table.
func (t *Table) ContainsRow(rowKey interface{}) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	_, y := t.m[rowKey]
	return y
}

// ContainsValue returns true if a mapping with
// the given value exists in the table.
func (t *Table) ContainsValue(value interface{}) (y bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, column := range t.m {
		for _, v := range column {
			if v == value {
				y = true
				break
			}
		}
	}
	return
}

// Get returns the value associated with the given row and column
// keys, or nil if no such mapping exists in the table.
func (t *Table) Get(rowKey, columnKey interface{}) interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.m[rowKey][columnKey]
}

// IsEmpty returns true if no mappings exist in the table.
func (t *Table) IsEmpty() (y bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if len(t.m) == 0 {
		y = true
	}
	return
}

// Len returns the number of mappings in the table.
func (t *Table) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.len
}

// Put associates the given value with the given row key and column key.
// For already existing mappings, replaces the old value with the given
// value, and returns the value previously associated with the keys.
// Returns nil if the mapping did not exist.
func (t *Table) Put(rowKey, columnKey, value interface{}) interface{} {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.m[rowKey] == nil {
		t.m[rowKey] = make(map[interface{}]interface{})
	}
	prev := t.m[rowKey][columnKey]
	t.m[rowKey][columnKey] = value
	t.len++
	return prev
}

// Remove removes the mapping associated with the given keys and
// returns the value previously associated with the keys.
// Returns nil if the mapping does not exist.
func (t *Table) Remove(rowKey, columnKey interface{}) interface{} {
	t.mu.Lock()
	defer t.mu.Unlock()
	var prev interface{}
	if _, ok := t.m[rowKey]; ok {
		prev = t.m[rowKey][columnKey]
		delete(t.m[rowKey], columnKey)
		t.len--
	}
	return prev
}

// Row returns a map that associates the column key with the value
// for all the mappings in the table that have the given row key.
// If there are no such mappings in the table, returns an empty map.
func (t *Table) Row(rowKey interface{}) map[interface{}]interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	m := make(map[interface{}]interface{})
	if row, ok := t.m[rowKey]; ok {
		m = row
	}
	return m
}

// New creates and returns a new (empty) table.
func New() *Table {
	return &Table{m: make(map[interface{}]map[interface{}]interface{})}
}
