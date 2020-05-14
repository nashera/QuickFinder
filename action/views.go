package action

import (
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/nashera/QuickFinder/model"
)

// ViewReports 拷贝结果到输出目录
func ViewReports(items []*model.ResultItem) error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "FullPath", "Sample", "Doctor", "Hospital"})
	// t.AppendRows([]table.Row{
	// 	{1, "Arya", "Stark", 3000},
	// 	{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	// })
	// t.AppendSeparator()
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	for i, item := range items {
		if item.ResultType == "Folder" {
			continue
		}
		if strings.HasSuffix(item.Name, "pdf") || strings.HasSuffix(item.Name, "docx") {
			for _, sample := range item.Samples {
				t.AppendRow(table.Row{
					i,
					item.Name,
					item.FullPath,
					sample.SampleName,
					sample.Doctor,
					sample.Hospital})

			}
			// shutil.Copy(item.FullPath, outputFolder, false)
		}
	}
	t.Render()
	return nil

}

// ViewSamples pretty view samples
func ViewSamples(samples []*model.Sample) error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "SampleName", "LibName", "Doctor", "Hospital"})
	// t.AppendRows([]table.Row{
	// 	{1, "Arya", "Stark", 3000},
	// 	{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	// })
	// t.AppendSeparator()
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	for i, sample := range samples {
		t.AppendRow(table.Row{
			i,
			sample.SampleName,
			sample.LibName,
			sample.Doctor,
			sample.Hospital})
	}
	t.Render()
	return nil
}
