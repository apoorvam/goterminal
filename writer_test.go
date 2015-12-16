package goterminal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriter(t *testing.T) {
	w := New()
	b := &bytes.Buffer{}
	w.Out = b
	fmt.Fprint(w, "foo bar")
	w.Print()

	if b.String() != "foo bar" {
		t.Fatalf("want %q, got %q", "foo bar", b.String())
	}
}
