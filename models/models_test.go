package models

import (
	"reflect"
	"testing"
)

func TestDog(t *testing.T) {

	want := &Dog{
		Name: "Doggy",
		Age:  1,
		Kind: Kind{Name: "Bulldog"},
	}
	got := &Dog{
		Name: "Doggy",
		Age:  1,
		Kind: Kind{Name: "Bulldog"},
	}

	if *got != *want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestDogReflect(t *testing.T) {

	want := &Dog{
		Name: "Doggy",
		Age:  1,
		Kind: Kind{Name: "Bulldog"},
	}
	got := &Dog{
		Name: "Doggy",
		Age:  1,
		Kind: Kind{Name: "Bulldog"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
