package mathy

import (
	"testing"
)

func TestClipInt(t *testing.T) {
	x := ClipInt(-5, -2, 2)
	if x != -2 {
		t.Errorf("Expected -2, got %v", x)
	}

	x = ClipInt(5, -2, 2)
	if x != 2 {
		t.Errorf("Expected 2, got %v", x)
	}

	x = ClipInt(1, -2, 2)
	if x != 1 {
		t.Errorf("Expected 1, got %v", x)
	}
}
