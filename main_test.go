package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting one")
	}

	test.Type(in, "Andy")
	if out.Text != "Hello Andy !" {
		t.Error("Incorrect initial greeting two")
	}
}
