package api

import (
	"testing"
)

func TestNewRouter(t *testing.T) {
	rt := NewRouter()
	if rt == nil {
		t.Fatal("Expected router, got nothing")
	}
}
