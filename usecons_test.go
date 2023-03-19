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
func TestCommaSeparatedIntToArrayCorrect(t *testing.T) {
	intSeparatedArray := "12,14,-56,1046"
	want := []int16{12, 14, -56, 1046}

	if got, status := CommaSeparatedIntToArray(intSeparatedArray); !(status && reflect.DeepEqual(got, want)) {
		t.Errorf("TestCommaSeparatedIntToArrayCorrect() = %q, want %q", got, want)
	}
}

// Not qeual arrays
func TestCommaSeparatedIntToArrayIncorrect(t *testing.T) {
	intSeparatedArray := "12,14,-56,1044"
	want := []int16{12, 14, -56, 1046}

	if got, status := CommaSeparatedIntToArray(intSeparatedArray); !(status && !reflect.DeepEqual(got, want)) {
		t.Errorf("TestCommaSeparatedIntToArrayIncorrect() = %q, want %q", got, want)
	}
}

// Not qeual arrays

func TestGenerateIntArray(t *testing.T) {
	arr := GenerateIntArray(10)
	threshold := 999

	for i := 0; i < len(arr); i++ {
		if arr[i] > threshold {
			t.Errorf("An array item (%v) with index %v is more than allowed threshold (%v)", arr[i], i, threshold)
		}
	}
}
