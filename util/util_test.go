package util

import "testing"

func TestTrimNewline(t *testing.T) {
	str := "foo\n"
	actual := TrimNewline(str)

	if actual != "foo" {
		t.Errorf("Expected foo got: %s", actual)
	}
}
