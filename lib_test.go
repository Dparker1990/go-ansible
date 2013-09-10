package main

import "testing"

func TestTrimNewline(t *testing.T) {
	str := "foo\n"
	actual := trimNewline(str)

	if actual != "foo" {
		t.Errorf("Expected foo got: %s", actual)
	}
}
