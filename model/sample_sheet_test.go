package model

import (
	"fmt"
	"testing"
)

func TestReadSheet(t *testing.T) {
	var sampleList = SampleSheetRead(SampleXlsxPath)
	// SampleSheetRead(SampleXlsxPath)
	for _, sample := range sampleList {
		fmt.Println(sample.LibName)
		fmt.Println(sample.Doctor)
		fmt.Println(sample.Hospital)
	}
}
