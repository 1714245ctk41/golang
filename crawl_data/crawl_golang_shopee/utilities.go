package main

import(
	"fmt"
	//"os"
	"io"
	"log"
	"net/http"
	"regexp"
	//"reflect"
	"strings"
	//"encoding/json"
	"strconv"
)
func learnRegex(){
	ari:= "https://shopee.vn/api/v2/category_list/get_all?level=2"
	dataSource := GetDataAPI(ari)
	re := regexp.MustCompile(`"display_name":"(.*?)".+catid":(\d+),"image":"(.*)","parent`)
  	dataSourceStrAr := re.FindAllSubmatch([]byte(dataSource), -1)
  	for _, v := range dataSourceStrAr {
  		fmt.Println("=====================================")

  		for _, v1 := range v {
  		fmt.Println("_______________________________")
  			fmt.Println(string(v1))
  			
  		}
  	}
}

func AllShopTopSale(){
	//var AllShopTopSaleId []string
	//var AllShopTopSaleShopInfo []ShopInfo
	// var AllShopTopSaleDetailShopInfo []DetailShopInfo
	var AllProductCat []ProductsCat
	
	dataForCategory := DataForCategory()
	
	for i, v := range dataForCategory {
		 productCat100 := DataForProductCat(strconv.Itoa(int(v.Catid)))
		 AllProductCat = append(AllProductCat, productCat100...)
		 fmt.Println(i)
	}
	fmt.Println()
		fmt.Println(AllProductCat)
		fmt.Println(len(AllProductCat))
	fmt.Println()
	// for i, v := range AllProductCat {
	// 	dataForShopInfo := DataForShopInfo(strconv.Itoa(int(v.Shopid)))
	// 	detailShopInfo := DataForDetailShopInfo(strconv.Itoa(int(dataForShopInfo.Shopid)) , dataForShopInfo.Username )
	// 	AllShopTopSaleDetailShopInfo = append(AllShopTopSaleDetailShopInfo, detailShopInfo)
	// 	fmt.Println(i, "____", len(AllShopTopSaleDetailShopInfo));
	// }
	// 	fmt.Println(AllShopTopSaleDetailShopInfo)
	// 	fmt.Println("_______________________________")
	// 	fmt.Println(len(AllShopTopSaleDetailShopInfo))
	// 	fmt.Println()
}




type DetailShopInfo struct{
	Id string `json:"id"`
	Userid string `json:"userid"`
	Shopid string `json:"shopid"`
	Name string `json:"name"`
	Username string `json:"username"`
	Phone string `json:"phone"`
	Region string `json:"region"`
	City string `json:"city"`
	Address string `json:"address"`
	District string `json:"district"`
}

func DataForDetailShopInfo(shopId string, userName string) DetailShopInfo{
	api := "https://shopee.vn/api/v4/shop/get_shop_detail?shopid=" + shopId + "&sort_sold_out=1&username="+ userName +""

	NameVsUserName := CustomData(api, `"name":(.*?)"(.*?)"`,  false)
	Phone := CustomData(api, `"phone":(.*?)"(.*?)"`,  false)
	Region	:= CustomData(api, `"region":(.*?)"(.*?)"`,  false)
	City	:= CustomData(api, `"city":(.*?)"(.*?)"`,  false)
	Address	:= CustomData(api, `"address":(.*?)"(.*?)"`,  false)
	District	:= CustomData(api, `"district":(.*?)"(.*?)"`,  false)
	Shopid:=  CustomData(api, `("shopid":(.*?)\d+)`,  false)[0][1]
	Userid:=  CustomData(api, `("userid":(.*?)\d+)`,  false)[0][1]
	Id := CustomData(api, `("id":(.*?)\d+)`,  false)[0][1]
	fmt.Println(len(NameVsUserName))
	fmt.Println(len(Phone))
	fmt.Println(len(Region))
	fmt.Println(len(City))
	fmt.Println(len(Address))
	fmt.Println(len(District))
	fmt.Println(len(Shopid))
	fmt.Println(len(Id))
	
	
	detailShopInfo := DetailShopInfo{
		Id :	BeautyString(Id),
		Userid :	BeautyString(Userid),
		Shopid :	BeautyString(Shopid),
		Name :	BeautyString(NameVsUserName[0][1]),
		Username :	BeautyString(NameVsUserName[1][1]),
		Phone :	BeautyString(Phone[0][1]),
		Region :	BeautyString(Region[0][1]),
		City :	BeautyString(City[0][1]),
		Address :	BeautyString(Address[0][1]),
		District :	BeautyString(District[0][1]),
	}

	return detailShopInfo
}
func BeautyString(s string) string{
		if strings.Contains(s, "null") {
			s = "null"
		}
		s = strings.ReplaceAll(s, `{`, "")
		s = strings.ReplaceAll(s, `}`, "")
		s = strings.ReplaceAll(s, `id`, "")
	return s
}

type ShopInfo struct{
	Shopid uint64 `json:"shopid"`
	Userid uint64 `json:"userid"`
	Username	string `json:"username"`
}
func DataForShopInfo(shopId string ) ShopInfo{
	api := "https://shopee.vn/api/v4/product/get_shop_info?shopid=" + shopId + ""
	Shopid,err :=  strconv.ParseUint(CustomData(api, `("shopid":\d+)`,  false)[0][1], 10, 64)
	Userid,err :=  strconv.ParseUint(CustomData(api, `("userid":\d+)`,  false)[0][1], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	Username :=  CustomData(api, `"username":"(.*?)"`,  true )

	shopInfo := ShopInfo{
		Shopid:Shopid ,
		Userid:Userid ,
		Username:Username[0][1] ,
	}
	return shopInfo
}

// products Category
type ProductsCat struct{
	Itemid uint64 `json:"itemid"`
	Shopid uint64 `json:"shopid"`
	NameProduct string `json:"nameproduct"`
	Price uint64 `json:"price"`
	Image string `json:"image"`
	Currency string `json:"currency"`
	Historical_sold uint64 `json`
	Shop_location string `json:"shop_location"`
}

func DataForProductCat(matchCatId string) []ProductsCat {
	api := "https://shopee.vn/api/v4/search/search_items?by=sales&limit=100&match_id=" + matchCatId + "&newest=140&order=desc&page_type=search&scenario=PAGE_SUB_CATEGORY_SEARCH&version=2"
	names := CustomData(api, `,"name":"(.*?)","label`,  true )
    itemIds := uniqueArray(CustomData(api, `("itemid":\d+)`,  false))
 	shopIds := uniqueArray(CustomData(api, `("shopid":\d+)`,  false))
  	prices := CustomData(api, `"price":(\d+)`,  false)
  	currencies := CustomData(api, `"currency":"(.*?)"`,  false)
  	historical_solds := CustomData(api, `"historical_sold":(\d+)`,  false)
  	images := CustomData(api, `"image":"(.*?)"`,  false)
  	shop_locations := CustomData(api, `"shop_location":(.*?)"(.*?)"`,  false)
  	ProductsCategories := []ProductsCat{}
  	

  	for i := 0; i < len(itemIds) ; i++ {
		  	
  			Itemid, err := strconv.ParseUint(itemIds[i][1], 10, 64)
			Shopid, err := strconv.ParseUint(shopIds[i][1], 10, 64)
			Price, err := strconv.ParseUint(prices[i][1], 10, 64)
			Historical_sold, err := strconv.ParseUint(historical_solds[i][1], 10, 64)
			if err != nil {
				fmt.Println("can't convert to Uint64")
			}

  			productCag := ProductsCat{
  			Itemid : Itemid,
			Shopid : Shopid,
			NameProduct : names[i][1] ,
			Price : Price,
			Image : images[i][1],
			Currency : currencies[i][1],
			Historical_sold : Historical_sold,
			Shop_location : shop_locations[i][1],
  			}

  		ProductsCategories = append(ProductsCategories, productCag)
  	}
  	return ProductsCategories

}

func uniqueArray(duplicateValue [][]string) [][]string{
	var itemIdsUnique [][]string
   	for i, v := range duplicateValue {
   		
   		if i%2 != 0 {
   			itemIdsUnique = append(itemIdsUnique, v)
   		} 
   	}
   	return itemIdsUnique
}

func CustomData(url string, pattern string, name bool ) [][]string{
	dataGet := GetDataAPI(url)
	re := regexp.MustCompile(pattern)
  	dataFinded := re.FindAllString(dataGet, -1)
  	if(name){
  		namesString := strings.Join(dataFinded,"")
	 	reg, _ := regexp.Compile(`\[([^\[\]]*)\]`)
	 	namesClean := reg.ReplaceAllString(namesString, "")
	 	namesArray := strings.Split(namesClean, `","`)
	 	var varray [][]string
	  	for _, v := range namesArray {
		 v = strings.ReplaceAll(v, `"`, "")
	  		varray = append(varray, strings.Split(v, ":")) 
	  	}
	  	return varray
  	}
  	var varray [][]string
  	for _, v := range dataFinded {
	 v = strings.ReplaceAll(v, `"`, "")
  		varray = append(varray, strings.Split(v, ":")) 
  	}
  	return varray
}

//Category
type Category struct{
	Display_name string `json:"display_name"`
	Catid uint64 `json:"catid"`
	Image string `json:"image"`

}

func DataForCategory() []Category{
	ari:= "https://shopee.vn/api/v2/category_list/get_all?level=2"
	dataSource := GetDataAPI(ari)
	re := regexp.MustCompile(`"display_name":"(.*?)".*\[([^\[\]]*)\]`)
  	dataSourceStrAr := re.FindAllString(dataSource, -1)
 	dataSourceStr := strings.Join(dataSourceStrAr,"")
  	dataSourceArray := strings.Split(dataSourceStr, `,{"main":`)
	dataSourceFather := []string{}
	dataSourceChild := []string{}
  	i := 0
  	for _, v := range dataSourceArray {
  		dataS := strings.Split(v, `,"sub":`)
  		for _, v1 := range dataS {
  			if i%2 == 0{
  			 dataSourceFather = append(dataSourceFather, v1)
  			}else{
  			 dataSourceChild = append(dataSourceChild, v1)
  			}
  			i++
  		}
  	}
  	dataSourceChildString := strings.Join(dataSourceChild, "]}[{")
	dataChildDisplayName := CleanData(dataSourceChildString, `"display_name":"(.*?)"`)
	dataChildImage := CleanData(dataSourceChildString, `"image":"(.*?)"`)
	dataChildDCatId := CleanData(dataSourceChildString, `"catid":\d+`)
	categories := []Category{}
	
	// for i := 0; i < len(dataChildDCatId) ; i++ {
	for i := 0; i < 2 ; i++ {
		Catid, err := strconv.ParseUint(dataChildDCatId[i][1], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		category := Category{
			Display_name: dataChildDisplayName[i][1],
			Catid: Catid,
			Image: dataChildImage[i][1],
		}
		categories = append(categories, category)
	}
	 //fmt.Println(len(categories))
	 //fmt.Println(categories)
	return categories
}
func CleanData(dataSource string, pattern string) [][]string{
	dataPattern := regexp.MustCompile(pattern)
  	dataFinded := dataPattern.FindAllString(dataSource, -1)
	var varray [][]string
  	for _, v := range dataFinded {
	 v = strings.ReplaceAll(v, `"`, "")
	//  v = strings.ReplaceAll(v, `{`, "")
	// v = strings.ReplaceAll(v, `}`, "")
  		varray = append(varray, strings.Split(v, ":")) 
  	}
  	return varray
}

func GetDataAPI (url string) string{
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	dataToStr := string(body[:])
	//var result map[string]interface{}
	return dataToStr

}



//Detail Product
// type DetailProduct struct{
// 	Itemid uint64 `json:"itemid"`
// 	Shopid uint64 `json:"shopid"`
// 	Catid uint64 `json:"catid"`
// 	UserId uint64 `json:"userid"`
// 	name string `json:"name"`
// 	Price_min uint64 `json:"price_min"`
// 	Price_max uint64 `json:"price_max"`
// 	Show_discount uint64 `json:"show_discount"`
// 	Description string `json:"description"`
// }

// func DataForDetailProduct(itemId string, shopId string){
// 	api := "https://shopee.vn/api/v4/item/get?itemid=" + itemId + "&shopid=" + shopId
// 	names := CustomData(api,`,"name":"(.*?)","ctime.*itemid":(\d+?),"s` , true )
//     itemIds := CustomData(api, `("itemid":\d+)`,  false)
//     Shopid := CustomData(api, `("shopid":\d+)`,  false)
//     Catid := CustomData(api, `("catid":\d+)`,  false)
//     Price_min := CustomData(api, `("price_min":\d+)`,  false)
//     Price_max := CustomData(api, `("price_max":\d+)`,  false)
//     Show_discount := CustomData(api, `("show_discount":\d+)`,  false)
//     Description := CustomData(api, `"description":"(.*?)"`,  false)
//     UserId := CustomData(api, `"(userid":\d+)`,  false)

//     itemid, err := strconv.ParseUint(itemIds[0][1], 10, 64)
//     shopid, err := strconv.ParseUint(Shopid[0][1], 10, 64)
//     catid, err := strconv.ParseUint(Catid[0][1], 10, 64)
//     userid, err := strconv.ParseUint(UserId[0][1], 10, 64)
//     price_min, err := strconv.ParseUint(Price_min[0][1], 10, 64)
//     price_max , err:= strconv.ParseUint(Price_max[0][1], 10, 64)
//     show_discount, err := strconv.ParseUint(Show_discount[0][1], 10, 64)

//     if err != nil {
//     	fmt.Println("can't change to uinit64")
//     }


//     DetailPro := DetailProduct{
//     Itemid :	 itemid,
// 	Shopid :		shopid ,
// 	Catid :		catid ,
// 	UserId :		userid ,
// 	name :		names[0][1] ,
// 	Price_min :		price_min ,
// 	Price_max :		price_max ,
// 	Show_discount :		show_discount ,
// 	Description :		Description[0][1] ,
//     }

// 	fmt.Println(DetailPro)
// }


