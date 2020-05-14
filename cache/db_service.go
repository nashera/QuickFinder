package cache

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nashera/QuickFinder/model"
)

// ResultDBPath 保存结果数据库

// func walkFunc(path string, info os.FileInfo, err error) error {

// 	var item = model.CreateResultItem(path, info)
// 	dc, _ := ConnectDB(ResultDBPath)

// 	dc.InsertResult(item)
// 	defer dc.CloseDB()
// 	return nil
// }

// BuildLocalDB 建立本地文件路径数据库
func BuildLocalDB(folderPath string, dbPath string) error {
	fmt.Println(folderPath)
	fmt.Println(dbPath)
	filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		dc, _ := ConnectDB(dbPath)
		dc.CreateDB()
		var item = model.CreateResultItem(path, info)
		dc.InsertResult(item)
		defer dc.CloseDB()
		return nil
	})
	return nil
}

// SearchLocalDB 搜索本地文件数据库
func SearchLocalDB(searchPattern string) []model.ResultItem {
	return nil

}

// UpdateLocalDB 更新本地数据库
func UpdateLocalDB(folderPath string, dbPath string) error {
	DeleteDB(dbPath)
	BuildLocalDB(folderPath, dbPath)
	return nil
}
