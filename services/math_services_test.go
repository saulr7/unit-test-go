package services

import "testing"

func TestAdd(t *testing.T) {

	want := 4
	got := Add(2, 2)

	if got != want {
		t.Errorf("Error: Expected %d, got %d", want, got)
	}
}

func TestAddMany(t *testing.T) {

	want := 6
	got := AddMany(1, 2, 3)

	if got != want {
		t.Errorf("Error: Expected %d, got %d", want, got)
	}

}
