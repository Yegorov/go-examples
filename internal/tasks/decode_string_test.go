package tasks

import (
	"strings"
	"testing"
)

func TestDecodeStringSimple(t *testing.T) {
	got := DecodeString("3[a]2[bc]")
	want := "aaabcbc"
	if !strings.Contains(got, want) {
		t.Errorf("Expexted %q, but got %q", want, got)
	}
}
func TestDecodeStringNested(t *testing.T) {
	got := DecodeString("3[a2[c]]")
	want := "accaccacc"
	if !strings.Contains(got, want) {
		t.Errorf("Expexted %q, but got %q", want, got)
	}
}
