package search

import (
	"fmt"
	"os"

	"github.com/nashera/QuickFinder/model"
)

// func String()

func walkFunc(path string, info os.FileInfo, err error) error {
	// fmt.Printf("%s \n", path)
	f = model.CreateResultItem(info)
	fmt.Println(f.Name)
	fmt.Println(f.ResultType)
	return nil
}
