package usecons

import (
	"reflect"
	"testing"
)

func TestPrintHeader(t *testing.T) {
	header := "Hello, world."
	want := "\nHello, world.\n-------------\n"
	if got := Header(&header); got != want {
		t.Errorf("Header() = %q, want %q", got, want)
	}
}

// Equal arrays
func Test1CheckCommaSeparatedIntArray(t *testing.T) {
	intSeparatedArray := "12,14,-56,1046"
	want := []int16{12, 14, -56, 1046}

	if got, status := CommaSeparatedToIntArray(intSeparatedArray); !(status && reflect.DeepEqual(got, want)) {
		t.Errorf("CommaSeparatedToIntArray() = %q, want %q", got, want)
	}
}

// Not qeual arrays
func Test2CheckCommaSeparatedIntArray(t *testing.T) {
	intSeparatedArray := "12,14,-56,1044"
	want := []int16{12, 14, -56, 1046}

	if got, status := CommaSeparatedToIntArray(intSeparatedArray); !(status && !reflect.DeepEqual(got, want)) {
		t.Errorf("CommaSeparatedToIntArray() = %q, want %q", got, want)
	}
}
