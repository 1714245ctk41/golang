package main

import (
	//  "encoding/json"
	// "os"
	"fmt"
)


func main() {
	//learnRegex()
	//DataForProductCat("11035583")
	// AllShopTopSale()

	//fmt.Println(DataForShopInfo("401401593"))
	DetailShopInfo := DataForDetailShopInfo("406556054", "veestore33")
	fmt.Println(DetailShopInfo)
	//DataForDetailProduct("7871796272", "89827191")
	// s := DataForProductCat("11036314")
	//DataForProductCat("11036928")
	 //fmt.Println(s)
	//DataForCategory()
	// s := DataForCategory()
	// b, err := json.Marshal(DetailShopInfo)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// 	os.Stdout.Write(b)

}
//reflect.TypeOf(element)



