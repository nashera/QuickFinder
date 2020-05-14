package cache

import (
	"testing"
)

const ResultDBPath string = "D:/Project/QuickFinder/local_result.db"

func TestDB(t *testing.T) {
	// CreateDB()
	BuildLocalDB("Z:/G_Counseling/Report/子阅", ResultDBPath)
	UpdateLocalDB("Z:/G_Counseling/Report/子阅", ResultDBPath)
	// dc, _ := ConnectDB(ResultDBPath)
	// defer dc.CloseDB()
	// dc.CreateDB()
	// var results = dc.QueryResult("仁济")
	// fmt.Println(results)
	// for i, x := range results {
	// 	// fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i,*x, x)
	// 	fmt.Println(i)
	// 	fmt.Println(x)
	// 	for _, sample := range x.Samples {
	// 		fmt.Println(sample)
	// 	}
	// }

}
