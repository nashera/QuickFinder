package model

import (
	"fmt"
	"strings"

	"github.com/nashera/QuickFinder/finderconfig"
	"github.com/tealeg/xlsx"
)

var (
	myconfig       = finderconfig.GetConfig()
	sampleXlsxPath = myconfig.SampleSheet
	sampleList     = SampleSheetRead(sampleXlsxPath)
)

// SampleSheetRead 读取样本进度表
func SampleSheetRead(sheetPath string) []*Sample {
	xlFile, err := xlsx.OpenFile(sheetPath)
	var sampleList []*Sample
	if err != nil {
		fmt.Println(err.Error())
	}
	var sampleSheet = xlFile.Sheet["Sheet1"]
	var header []string
	for i, row := range sampleSheet.Rows {
		if i == 0 {
			for _, cell := range row.Cells {
				header = append(header, cell.String())
			}
		} else {
			rowMap := make(map[string]string)
			for j, cell := range row.Cells {
				if j > len(header)-1 {
					break
				}
				rowMap[header[j]] = cell.String()
			}
			sampleList = append(sampleList,
				&Sample{
					SampleName: rowMap["样本编号"],
					LibName:    rowMap["文库编号"],
					Doctor:     rowMap["送检医生"],
					Hospital:   rowMap["送检单位"],
				})

		}
	}
	return sampleList
}

// FindSampleWithSampleName 利用样本名寻找Sample
func FindSampleWithSampleName(query string) *Sample {
	if strings.HasPrefix(query, "A") {
		for _, sample := range sampleList {
			if strings.HasPrefix(sample.SampleName, query) {
				return sample
			}
		}
	} else {
		for _, sample := range sampleList {
			if strings.HasPrefix(sample.LibName, query) {
				return sample
			}
		}
	}
	var sampleNA = &Sample{
		SampleName: "NA",
		LibName:    "NA",
		Doctor:     "NA",
		Hospital:   "NA",
	}

	return sampleNA
}
