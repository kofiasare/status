package main

import "testing"

var hsc = &HTTPStatusCodes{}

func TestLookUp(t *testing.T) {

	for _, d := range testData {
		code, err := hsc.LookUp(d.statusCode, false)

		if err != nil && err.Error() != d.err.Error() {
			t.Errorf("expected %s, got %s", err, d.err)
		}

		if err == nil && code.message != d.message {
			t.Errorf("expected %d => %s, got %s", d.statusCode, code.message, d.message)
		}
	}
}

func TestColor(t *testing.T) {
	for _, d := range testData {
		code, err := hsc.LookUp(d.statusCode, false)

		if err != nil && code.color != d.color {
			t.Errorf("expected color of %d == %v, got %v", d.statusCode, code.color, d.color)
		}

		if err == nil && code.color != d.color {
			t.Errorf("expected color of %d == %v, got %v", d.statusCode, code.color, d.color)
		}
	}
}
