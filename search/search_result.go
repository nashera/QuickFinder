package search

import (
	"fmt"
	"os"

	"github.com/nashera/QuickFinder/model"
)

// func String()

func walkFunc(path string, info os.FileInfo, err error) error {

	var f = model.CreateResultItem(path, info)
	fmt.Printf("%s \n", f.FullPath)
	fmt.Println(f.Name)
	fmt.Println(f.ResultType)
	return nil
}
