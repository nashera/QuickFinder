package action

import (
	"testing"

	"github.com/nashera/QuickFinder/cache"
)

func TestViews(t *testing.T) {

	var results = cache.QueryResult("仁济")
	// fmt.Println(results)
	// for i, x := range results {
	// 	// fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i,*x, x)
	// 	fmt.Println(i)
	// 	fmt.Println(x)
	// 	for _, sample := range x.Samples {
	// 		fmt.Println(sample)
	// 	}
	// }
	ViewReports(results)
}
