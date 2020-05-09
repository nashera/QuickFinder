package model

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

// ResultItem 搜索结果
type ResultItem struct {
	Name       string
	ResultType string
	Modified   time.Time
	FullPath   string
	Samples    []*Sample
}

// CreateResultItem Constructor
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
		FullPath:   path,
		Samples:    getRelatedSamples(info),
	}
	if r.Samples != nil {
		fmt.Println(info.Name())
		for _, sample := range r.Samples {
			fmt.Println(sample.SampleName)
		}

	}
	return &r
}

// GetSampleName print resultitem
func getRelatedSamples(info os.FileInfo) []*Sample {

	re := regexp.MustCompile(`A\w{3,4}|Lib\w{3,4}`)
	if info.IsDir() {
		return nil
	}
	if !strings.HasSuffix(info.Name(), "docx") && !strings.HasSuffix(info.Name(), "pdf") {
		return nil
	}
	var matchResult = re.FindAllString(info.Name(), -1)
	if matchResult == nil {
		return nil
	}
	var samples []*Sample
	for _, r := range matchResult {
		samples = append(samples, FindSampleWithSampleName(r))
	}
	return samples
}
