package model

import (
	"os"
	"time"
)

// Report 搜索结果
type Report struct {
	Name       string
	ResultType string
	Modified   time.Time
	FullPath   string
}

// CreateReport Report Constructor
func CreateReport(path string, info os.FileInfo) *ResultItem {
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
		FullPath:   path,
	}
	return &r
}
