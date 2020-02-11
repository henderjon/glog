package logger

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHere(t *testing.T) {
	expected := "location_test.go:12"

	e := Here()
	if diff := cmp.Diff(string(e), expected); diff != "" {
		t.Error("Encode(e); (-got +want)", diff)
	}
}
