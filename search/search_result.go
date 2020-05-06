package search

import (
	"os"

	"github.com/nashera/QuickFinder/cache"
	"github.com/nashera/QuickFinder/model"
)

// func String()

func walkFunc(path string, info os.FileInfo, err error) error {

	var f = model.CreateResultItem(path, info)
	// var db = "./result.db"
	// if !cache.DBIsExisted(db) {
	// 	// cache.CreateDB(db)
	// 	fmt.Println(db)
	// }
	cache.InsertResult(f)
	// var c *cache.Context
	// c, err = cache.ConnectDB(db)
	// if err != nil {
	// 	print(err)
	// }
	// fmt.Println(c.DbPath)
	// c.CreateDB()
	// c.InsertResult(f)
	// fmt.Printf("%s \n", f.FullPath)
	// fmt.Println(f.Name)
	// fmt.Println(f.ResultType)

	return nil
}
