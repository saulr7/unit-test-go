package services_test

import (
	"testing"
	"unit-test/services"
)

func TestMultiply(t *testing.T) {
	want := 6

	got := services.Multiply(2, 3)

	if got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}
