package structs

import (
	"testing"
)

func TestStackSimple(t *testing.T) {
	stack := BuildStack[int]()
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	_ = stack.Pop()
	got := stack.Count()
	want := 2
	if got != want {
		t.Errorf("Expexted %d, but got %d", want, got)
	}
}
