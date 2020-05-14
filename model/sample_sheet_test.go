package model

import (
	"fmt"
	"testing"
)

// func TestReadSheet(t *testing.T) {
// 	var sampleList = SampleSheetRead(SampleXlsxPath)
// 	// SampleSheetRead(SampleXlsxPath)
// 	for _, sample := range sampleList {
// 		fmt.Println(sample.SampleName)
// 		fmt.Println(sample.LibName)
// 		fmt.Println(sample.Doctor)
// 		fmt.Println(sample.Hospital)
// 	}
// }

func TestFindSampleName(t *testing.T) {
	sample := FindSampleWithSampleName("A2054")
	fmt.Println(sample.Doctor)
}
