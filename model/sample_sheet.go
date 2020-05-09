package model

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

// Sample 样本
type Sample struct {
	SampleName string
	LibName    string
	Doctor     string
	Hospital   string
}

// SampleXlsxPath 样本进度表路径
const SampleXlsxPath string = "Z:/Project/项目进度表/阅尔基因项目进度表2-20180509.xlsx"

// SampleSheetRead 读取样本进度表
func SampleSheetRead(sheetPath string) []Sample {
	xlFile, err := xlsx.OpenFile(sheetPath)
	var sampleList []Sample
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
				// fmt.Println(header[j])
				// fmt.Println(cell.String())
				rowMap[header[j]] = cell.String()
				// }
				sampleList = append(sampleList,
					Sample{
						SampleName: rowMap["样本编号"],
						LibName:    rowMap["文库编号"],
						Doctor:     rowMap["送检医生"],
						Hospital:   rowMap["送检单位"],
					})
			}
		}
	}
	return sampleList
}

var sampleList = SampleSheetRead(SampleXlsxPath)

// FindSampleWithSampleName 利用样本名寻找Sample
func FindSampleWithSampleName(sampleName string) *Sample {
	for _, sample := range sampleList {
		if strings.HasPrefix(sample.SampleName, sampleName) {
			return &sample
		}
	}

	return nil
}
