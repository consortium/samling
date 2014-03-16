// Copyright (c) 2012, Robert Dinu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license which can be found in the LICENSE file.

package table

import (
	"testing"

	"github.com/noll/samling/test"
)

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	table := New()
	table.Put("a", "b", 1)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		table.Get("a", "b")
		b.StopTimer()
	}
}

func BenchmarkPut(b *testing.B) {
	b.StopTimer()
	table := New()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		table.Put(i, i, 1)
		b.StopTimer()
	}
}

func TestClear(t *testing.T) {
	table := New()
	table.Put("a", "b", 1)
	table.Put("b", "c", 1)
	table.Put("c", "d", 1)
	table.Clear()
	test.Verify(t, 1, 0, 0, table.Len())
}

func TestColumn(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, 0, len(table.Column("b")))
	table.Put("a", "b", 1)
	table.Put("b", "b", 2)
	table.Put("c", "b", 3)
	column := table.Column("b")
	test.Verify(t, 2, 0, 3, len(column))
	test.Verify(t, 3, 0, 1, column["a"])
	test.Verify(t, 4, 0, 2, column["b"])
	test.Verify(t, 5, 0, 3, column["c"])
	table.Put("b", "a", 1)
	column = table.Column("a")
	test.Verify(t, 6, 0, 1, len(column))
	test.Verify(t, 7, 0, 1, column["b"])
}

func TestContains(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, false, table.Contains("a", "b"))
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, true, table.Contains("a", "b"))
}

func TestContainsColumn(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, false, table.ContainsColumn("b"))
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, true, table.ContainsColumn("b"))
}

func TestContainsRow(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, false, table.ContainsRow("a"))
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, true, table.ContainsRow("a"))
}

func TestContainsValue(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, false, table.ContainsValue(1))
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, true, table.ContainsValue(1))
}

func TestGet(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, nil, table.Get("a", "b"))
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, 1, table.Get("a", "b"))
}

func TestIsEmpty(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, true, table.IsEmpty())
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, false, table.IsEmpty())
}

func TestLen(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, 0, table.Len())
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, 1, table.Len())
	table.Put("a", "c", 1)
	test.Verify(t, 3, 0, 2, table.Len())
	table.Put("a", "d", 1)
	test.Verify(t, 4, 0, 3, table.Len())
	table.Put("b", "a", 1)
	test.Verify(t, 5, 0, 4, table.Len())
	table.Put("b", "b", 1)
	test.Verify(t, 6, 0, 5, table.Len())
}

func TestPut(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, nil, table.Put("a", "b", 1))
	test.Verify(t, 2, 0, 1, table.Get("a", "b"))
	test.Verify(t, 3, 0, 1, table.Put("a", "b", 2))
	test.Verify(t, 4, 0, 2, table.Get("a", "b"))
}

func TestTableRemove(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, nil, table.Remove("a", "b"))
	table.Put("a", "b", 1)
	test.Verify(t, 2, 0, 1, table.Remove("a", "b"))
}

func TestRow(t *testing.T) {
	table := New()
	test.Verify(t, 1, 0, 0, len(table.Row("a")))
	table.Put("a", "b", 1)
	table.Put("a", "c", 2)
	table.Put("a", "d", 3)
	row := table.Row("a")
	test.Verify(t, 2, 0, 3, len(row))
	test.Verify(t, 3, 0, 1, row["b"])
	test.Verify(t, 4, 0, 2, row["c"])
	test.Verify(t, 5, 0, 3, row["d"])
	table.Put("b", "a", 1)
	row = table.Row("b")
	test.Verify(t, 6, 0, 1, len(row))
	test.Verify(t, 7, 0, 1, row["a"])
}
