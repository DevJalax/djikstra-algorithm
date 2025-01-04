package testcase

import "testing"

func TestAdd(t *testing.T) {

	got := Add(10, 20)
	want := 30

	if got != want {

		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestSubtract(t *testing.T) {

	got := Subtract(10, 20)
	want := 10

	if got != want {

		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestMul(t *testing.T) {

	got := Mul(10, 20)
	want := 2000

	if got != want {

		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestDiv(t *testing.T) {

	got := Div(10, 20)
	want := 2

	if got != want {

		t.Errorf("got %q, wanted %q", got, want)
	}

}
