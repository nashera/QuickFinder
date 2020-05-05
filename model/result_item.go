package model

import (
	"os"
	"time"
)

// ResultItem 搜索结果
type ResultItem struct {
	Name       string
	ResultType string
	Modified   time.Time
	FullPath   string
}

func CreateResultItem(path string, info os.FileInfo) *ResultItem {
	var r ResultItem
	var resultType string
	if info.IsDir() {
		resultType = "Folder"
	} else {
		resultType = "File"
	}
	r = ResultItem{
		Name:       info.Name(),
		Modified:   info.ModTime(),
		ResultType: resultType,
		FullPath, path
	}
	return &r
}

func String(result *ResultItem) error {
	return nil
}
