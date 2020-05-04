package search

import (
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {
	filepath.Walk("Z:/G_Counseling/Report", walkFunc)
}
