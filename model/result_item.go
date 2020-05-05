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

func CreateResultItem(info os.FileInfo) *ResultItem {
	var r ResultItem
	var resultType string
	if info.IsDir() {
		resultType = "Folder"
	}
	else {
		resultType = "File"
	}
	r = ResultItem{
		Name:       info.Name(),
		Modified:   info.ModTime(),
		ResultType: resultType,
	}
	return &r
}
