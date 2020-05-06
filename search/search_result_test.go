package search

import (
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {
	filepath.Walk("Z:/G_Counseling/Report/子阅/202005", walkFunc)
}
