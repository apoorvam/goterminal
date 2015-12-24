package goterminal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriter(t *testing.T) {
	b := &bytes.Buffer{}
	w := New(b)
	fmt.Fprint(w, "foo bar")
	w.Print()

	if b.String() != "foo bar" {
		t.Fatalf("want %q, got %q", "foo bar", b.String())
	}
}
