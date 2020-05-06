package cache

import (
	"os"
	"path/filepath"

	"github.com/nashera/QuickFinder/model"
)

// ResultDbPath 保存结果数据库
const ResultDbPath string = "D:/Project/QuickFinder/local_result.db"

func walkFunc(path string, info os.FileInfo, err error) error {

	var item = model.CreateResultItem(path, info)

	InsertResult(item)
	return nil
}

// BuildLocalDB 建立本地文件路径数据库
func BuildLocalDB(folderPath string) error {
	filepath.Walk(folderPath, walkFunc)
	return nil
}

// SearchLocalDB 搜索本地文件数据库
func SearchLocalDB(searchPattern string) []model.ResultItem {
	return nil

}
