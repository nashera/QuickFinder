package search

import (
	"fmt"
	"os"
	"time"
)

// ResultItem 搜索结果
type ResultItem struct {
	name       string
	ResultType string
	Created    time.Time
	Modified   time.Time
	Accessed   time.Time
	FullPath   string
}

func createResultItem(info os.FileInfo) *ResultItem {
	var r ResultItem
	r = ResultItem{
		name: info.Name(),
	}
	return &r
}

// func String()

func walkFunc(path string, info os.FileInfo, err error) error {
	// fmt.Printf("%s \n", path)

	fmt.Println(createResultItem(info).name)
	return nil
}
