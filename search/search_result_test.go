package search

import (
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {
	filepath.Walk("D:/project/QuickFinder", walkFunc)
}
