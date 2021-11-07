package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//"io"
//"log"
//"net/http"
//"regexp"
//"reflect"
//"strings"
//"encoding/json"
//"strconv"

func main() {
	// x := GetNumberProduct("https://www.vascara.com/phu-kien/mat-kinh-nu")
	// fmt.Println(x)
	GetDataVascara()
	// productDetailVascava("https://www.vascara.com/giay-cao-got/giay-sandal-ankle-strap-van-da-ky-da-bmn-0489-mau-den")

}

// working with json file
//productDetail
func readJSONDe(nameFile string) []ProductDetail {
	// Open our jsonFile
	jsonFile, err := os.Open(nameFile + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []ProductDetail
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

// class="price sale">(.*?)<\/span>
// DataPatternAPI(`<div class="item col-\d">(.+?)<\/ul>\s*<\/div>\s*<\/div>`,"https://www.vascara.com/giay")
// DataPatternAPI(``,"https://www.vascara.com/giay")
