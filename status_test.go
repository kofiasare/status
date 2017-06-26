package main

import "testing"

var hsc = new(HTTPStatusCodess)

func TestLookUp(t *testing.T) {

	hsc := new(HTTPStatusCodess)

	for _, d := range testData {
		h, err := hsc.LookUp(d.statusCode, false)

		if err != nil && err.errorMessage() != d.errorMessage() {
			t.Errorf("expected %d => %s, got %s", d.statusCode, err.errorMessage(), d.errorMessage())
		}

		if err == nil && h.message != d.message {
			t.Errorf("expected %d => %s, got %s", d.statusCode, h.message, d.message)
		}
	}
}

func TestColor(t *testing.T) {
	for _, d := range testData {
		h, err := hsc.LookUp(d.statusCode, false)

		if err != nil && h.color != d.color {
			t.Errorf("expected color of %d == %v, got %v", d.statusCode, h.color, d.color)
		}

		if err == nil && h.color != d.color {
			t.Errorf("expected color of %d == %v, got %v", d.statusCode, h.color, d.color)
		}
	}
}
