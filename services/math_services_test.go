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

func TestMultiply(t *testing.T) {

	table := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"2x1", 2, 1, 2},
		{"2x2", 2, 2, 4},
		{"2x3", 2, 3, 6},
		{"2x4", 2, 4, 8},
	}

	for _, v := range table {

		t.Run(v.name, func(t *testing.T) {

			got := multiply(v.a, v.b)

			if got != v.want {
				t.Errorf("Error: Expected %d, got %d", v.want, got)
			}

		})

	}

}
