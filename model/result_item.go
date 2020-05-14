package model

import (
	"os"
	"regexp"
	"strings"
	"time"
)

// Sample 样本
type Sample struct {
	SampleName string
	LibName    string
	Doctor     string
	Hospital   string
}

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
	var item ResultItem
	var resultType string
	if info.IsDir() {
		resultType = "Folder"
	} else {
		resultType = "File"
	}
	item = ResultItem{
		Name:       info.Name(),
		Modified:   info.ModTime(),
		ResultType: resultType,
		FullPath:   path,
		Samples:    getRelatedSamples(info),
	}
	// if item.Samples != nil {
	// 	fmt.Println(info.Name())
	// 	for _, sample := range item.Samples {
	// 		fmt.Println(sample)
	// 		// fmt.Println(sample.Doctor)
	// 	}

	// }
	return &item
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
