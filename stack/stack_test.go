// Copyright (c) 2012, Robert Dinu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license which can be found in the LICENSE file.

package stack

import (
	"testing"

	"github.com/noll/samling/test"
)

func TestClear(t *testing.T) {
	stack := New()
	stack.Push("one")
	stack.Push("two")
	stack.Push("three")
	stack.Push("four")
	stack.Clear()
	test.Verify(t, 1, 0, true, nil == stack.Peek())
	test.Verify(t, 2, 0, 0, stack.Len())
}

func TestPeek(t *testing.T) {
	stack := New()
	test.Verify(t, 1, 0, true, nil == stack.Peek())
}

func TestLenPopPush(t *testing.T) {
	stack := New()
	test.Verify(t, 1, 0, true, nil == stack.Pop())
	test.Verify(t, 2, 0, 0, stack.Len())
	stack.Push("one")
	test.Verify(t, 3, 0, 1, stack.Len())
	test.Verify(t, 4, 0, "one", stack.Pop())
	test.Verify(t, 5, 0, 0, stack.Len())
	stack.Push("two")
	stack.Push("three")
	test.Verify(t, 6, 0, 2, stack.Len())
	test.Verify(t, 7, 0, "three", stack.Pop())
	test.Verify(t, 8, 0, 1, stack.Len())
	test.Verify(t, 9, 0, "two", stack.Pop())
	test.Verify(t, 10, 0, 0, stack.Len())
	test.Verify(t, 11, 0, true, nil == stack.Pop())
}
