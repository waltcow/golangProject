package demo1

import (
	"testing"
)

func TestBusinessLogicWithMem(t *testing.T) {
	ht := NewMemHashTable()
	BusinessLogic(ht)
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("error on Get: %s", err)
	}
	if string(val) != "world" {
		t.Fatalf("expected 'world', got '%s'", string(val))
	}
}
