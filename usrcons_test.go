package usrcons

import "testing"

func TestPrintHeader(t *testing.T) {
	header := "Hello, world."
	want := "\nHello, world.\n-------------\n"
	if got := Header(&header); got != want {
		t.Errorf("Header() = %q, want %q", got, want)
	}
}
