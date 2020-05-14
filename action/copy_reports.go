package action

import (
	"fmt"
	"os"
	"strings"

	"github.com/nashera/QuickFinder/model"
	"github.com/termie/go-shutil"
)

// CopyToOutput 拷贝结果到输出目录
func CopyToOutput(items []*model.ResultItem, outputFolder string) error {
	fmt.Println(outputFolder)
	if _, staterr := os.Stat(outputFolder); os.IsNotExist(staterr) {
		os.MkdirAll(outputFolder, os.ModePerm)

	}
	for _, item := range items {
		if item.ResultType == "Folder" {
			continue
		}
		if strings.HasSuffix(item.Name, "pdf") || strings.HasSuffix(item.Name, "docx") {
			shutil.Copy(item.FullPath, outputFolder, false)
		}
	}
	return nil

}
