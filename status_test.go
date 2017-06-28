package main

import "testing"

var hsc = &HTTPStatusCodes{}

func TestLookUp(t *testing.T) {

	for _, d := range testData {
		sc, err := hsc.LookUp(d.statusCode, false)

		if err != nil && err.errorMessage() != d.errorMessage() {
			t.Errorf("expected %d => %s, got %s", d.statusCode, err.errorMessage(), d.errorMessage())
		}

		if err == nil && sc.message != d.message {
			t.Errorf("expected %d => %s, got %s", d.statusCode, sc.message, d.message)
		}
	}
}

func TestColor(t *testing.T) {
	for _, d := range testData {
		sc, err := hsc.LookUp(d.statusCode, false)

		if err != nil && sc.color != d.color {
			t.Errorf("expected color of %d == %v, got %v", d.statusCode, sc.color, d.color)
		}

		if err == nil && sc.color != d.color {
			t.Errorf("expected color of %d == %v, got %v", d.statusCode, sc.color, d.color)
		}
	}
}
