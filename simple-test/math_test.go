package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 7
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
