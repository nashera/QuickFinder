package search

import (
	"fmt"
	"os"
	"time"
)



// func String()

func walkFunc(path string, info os.FileInfo, err error) error {
	// fmt.Printf("%s \n", path)

	fmt.Println(createResultItem(info).name)
	return nil
}
