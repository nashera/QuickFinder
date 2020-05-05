package search

import (
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {
	filepath.Walk("J:/整理/新生突变的数据分析与遗传咨询", walkFunc)
}
