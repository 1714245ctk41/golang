package utils

import (
	"crawl_data/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	//"os"
	// "io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	//"reflect"
	"strings"

	"github.com/gocolly/colly/v2"
	//"encoding/json"
	//"strconv"
)

var linkFile = "./pkg/utils/rabbitmq_go/vascara_cache/"

//get data total
func GetDataVascara() {
	// cate := GetCategoryContainer("https://www.vascara.com/")
	// _, cateCollecChild := GetCategory(cate)

	// producReview := GetProductReview(cateCollecChild)

	// data := GetProductDetail(producReview, cateCollecChild)

	// BestSaleHandle()

}
func removeDuplicate(data []model.ProductDetail) []model.ProductDetail {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i].ID == data[j].ID {
				data[j].ID = 0
			}
		}
	}
	newData := []model.ProductDetail{}
	for _, v := range data {
		if v.ID != 0 {
			newData = append(newData, v)
		}
	}
	return newData
}

//* bestsale handle
func BestSaleHandle() []string {

	titleGroup := []string{}
	collector := colly.NewCollector()
	collector.OnHTML(".list-product .product-item", func(element *colly.HTMLElement) {
		title := element.ChildText(".product-title a")
		titleGroup = append(titleGroup, title)

	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL.String())
	})

	collector.Visit("https://www.vascara.com/best-sale?src=bestsale-homepage")

	// for i, v := range data {
	// 	for _, x := range titleGroup {
	// 		if strings.Contains(v.Title, x) {
	// 			data[i].BestSale = true
	// 			break
	// 		}
	// 	}
	// }
	return titleGroup
}

//* get product detail
func GetProductDetail(productViews []model.ProductView, categoryChild []model.CategoryChild) []model.ProductDetail {
	data := []model.ProductDetail{}
	for _, v := range productViews {

		data = append(data, ProductDetailVascava(v.Detaillink))

	}

	// data = removeDuplicate(data)
	err := os.RemoveAll(linkFile + "cache")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(data))
	return data
}

func ProductDetailVascava(url string) model.ProductDetail {

	collector := colly.NewCollector(
		colly.CacheDir(linkFile + "cache"),
	)

	productDetailChild := model.ProductDetail{}

	collector.OnHTML(".page-content", func(element *colly.HTMLElement) {
		productDetailChild = model.ProductDetail{
			ProductCode: element.ChildAttr("#productCode", "value"),
			Title:       element.ChildText(".title-product"),
			Image:       element.ChildAttr(".item img", "src"),
			Discount:    element.ChildText(".product-info .percent-discount"),
			Currency:    element.ChildTexts(".price .currency")[0],
		}

		idStr := element.ChildAttr("#productId", "value")
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println(err)
		}
		productDetailChild.ID = uint(idInt)
		price := element.ChildTexts(".price .amount")[0]
		price = strings.Replace(price, ".", "", -1)
		priceInt, err := strconv.ParseInt(price, 10, 64)
		if err != nil {
			panic(err)
		}
		productDetailChild.Price = priceInt
		inforso := element.ChildTexts(".list-oppr span")
		inforStr := ""
		for i, v := range inforso {
			if i%2 != 0 {
				inforStr += v + "||"
			} else {
				inforStr += v + "|"
			}
		}
		inforStr = strings.Replace(inforStr, "\u0026", "và ", -1)
		productDetailChild.Infor = inforStr
		content := element.ChildTexts(".breadcrumb a")
		productDetailChild.CategoryName = content[len(content)-1]

	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL.String())
	})

	collector.Visit(url)
	collector.Wait()
	return productDetailChild

	// writeJSON(allFacts)

	// enc := json.NewEncoder(os.Stdout)
	// // fmt.Println(enc)
	// enc.SetIndent("", "_")
	// enc.Encode(productDetailChild)
}

//* get number Product
func GetNumberProduct(url string) int {
	collector := colly.NewCollector()
	numberPr := 0
	collector.OnHTML(".page-content .cate-view-more", func(element *colly.HTMLElement) {
		numberStr := element.ChildText(".viewmore-totalitem")
		numberStr = strings.Trim(numberStr, " ")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		numberPr = number
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL.String())
	})

	collector.Visit(url)
	return numberPr

}

//* get productView
func GetProductReview(categories []model.CategoryChild) []model.ProductView {
	data := []model.ProductView{}
	for _, v := range categories {
		numPro := GetNumberProduct(v.LinkCategory) + 21
		x := numPro % 21
		y := numPro / 21
		if x > 0 {
			y = y + 1
		}
		url := v.Link
		for i := 1; i <= y; i++ {
			if i > 1 {
				url = strings.Replace(url, "page="+strconv.Itoa(i-1), "page="+strconv.Itoa(i), 1)
			}
			data = append(data, productVascava(url)...)
			fmt.Println(url)

		}

	}
	return data
}

func productVascava(url string) []model.ProductView {
	productView := make([]model.ProductView, 0)
	getDataAPI := getDataAPI(url)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(getDataAPI)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".product-item").Each(func(i int, s *goquery.Selection) {
		image, ok := s.Find(".avatar img").First().Attr("src")
		detaillink, ok := s.Find(".product-title a").First().Attr("href")
		if !ok {
			fmt.Println("failed find data attr")
		}
		title := s.Find(".product-title a").First().Text()

		title = strings.Trim(title, "\n")
		title = strings.Trim(title, " ")

		productVieChild := model.ProductView{
			Image:      image,
			Price:      s.Find(".amount").First().Text(),
			Currency:   s.Find(".currency").First().Text(),
			Title:      title,
			Detaillink: detaillink,
		}
		productView = append(productView, productVieChild)
	})
	// writeJSON(allFacts)
	return productView

}

//* get category
func FindFaCategory(data []model.CategoryFa) []model.CategoryFa {
	dataCopy := []model.CategoryFa{}

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i].ID == data[j].ID {
				dataCopy = append(dataCopy, data[j])
				break
			}
		}
	}

	return dataCopy
}
func GetCategory(cate []string) ([]model.CategoryFa, []model.CategoryChild) {
	cateChildCollec := []model.CategoryChild{}
	cateCollec := []model.CategoryFa{}
	for _, v := range cate {
		vChild := strings.Split(v, "|")
		cateId, cateName, cateNameFa := GetCategoryIdNameNameFa(vChild[1])
		if strings.Compare(cateId, "") != 0 {

			cateIdInt, err := strconv.Atoi(cateId)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			linkAPI := "https://www.vascara.com/product/filterproduct?page=1&cate=" + cateId + "&viewmore=1&viewcol=3"
			cateChild := model.CategoryFa{
				ID:           uint(cateIdInt),
				Name:         cateName,
				LinkCategory: vChild[1],
				Link:         linkAPI,
			}
			cateCollec = append(cateCollec, cateChild)
			if cateNameFa != "" {
				for _, vc := range cateCollec {
					if strings.Compare(vc.Name, cateNameFa) == 0 {
						cateChild := model.CategoryChild{
							ID:           uint(cateIdInt),
							IdFa:         vc.ID,
							Name:         cateName,
							LinkCategory: vChild[1],
							Link:         linkAPI,
						}
						cateChildCollec = append(cateChildCollec, cateChild)
						break
					}

				}

			}
		}
	}
	cateFaCollec := FindFaCategory(cateCollec)
	for i, v := range cateFaCollec {
		cateChildCollecChild := []model.CategoryChild{}
		for _, vc := range cateChildCollec {
			if v.ID == vc.IdFa {
				cateChildCollecChild = append(cateChildCollecChild, vc)
			}
		}
		cateFaCollec[i].CategoryChilds = cateChildCollecChild

	}
	for _, v := range cateFaCollec {
		fmt.Println(v)
	}

	return cateFaCollec, cateChildCollec
}

func GetCategoryIdNameNameFa(link string) (string, string, string) {

	collector := colly.NewCollector()
	cateSour := ""
	cateName := ""
	cateNameFa := ""
	collector.OnHTML(".content-filter", func(element *colly.HTMLElement) {
		cateSour = element.ChildAttr("#hdn_cate_id", "value")
	})
	collector.OnHTML(".breadcrumb", func(element *colly.HTMLElement) {
		cateName = element.ChildText(".breadcrumb h1")
		cateNameFaAr := element.ChildTexts(".breadcrumb a")
		if len(cateNameFaAr) > 1 {
			cateNameFa = cateNameFaAr[1]
		} else {
			cateNameFa = ""

		}

		fmt.Println(cateName)
	})

	collector.OnRequest(func(request *colly.Request) {

		fmt.Println("Visiting ", request.URL.String())

	})
	collector.Visit(link)
	return cateSour, cateName, cateNameFa

}

func GetCategoryContainer(url string) []string {
	collector := colly.NewCollector()
	cate := []string{}
	collector.OnHTML(".main-menu li", func(element *colly.HTMLElement) {
		cateSour := element.ChildText("a") + "|" + element.ChildAttr("a", "href")
		cate = append(cate, cateSour)
	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL.String())

	})
	collector.Visit(url)
	// for _, v := range cate {
	// 	fmt.Println(v)
	// }
	return cate
}

func getDataAPI(url string) *strings.Reader {
	dataSource := GetDataAPI(url)
	data := model.Source{}
	jsonErr := json.Unmarshal(dataSource, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	newReader := strings.NewReader(data.Html)
	return newReader
}

// working with json file
//productDetail
func ReadJSONDe(nameFile string) []model.ProductDetail {
	// Open our jsonFile
	jsonFile, err := os.Open(linkFile + nameFile + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []model.ProductDetail
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

//productReview
func readJSONRe(nameFile string) []model.ProductView {
	// Open our jsonFile
	jsonFile, err := os.Open(linkFile + nameFile + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []model.ProductView
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

//category
func readJSON(nameFile string) []model.CategoryChild {
	// Open our jsonFile
	jsonFile, err := os.Open(linkFile + nameFile + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []model.CategoryChild
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

func writeJSON(data []model.ProductView, nameFile string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile(linkFile+nameFile+".json", file, 0644)
}
func writeJSONDetail(data []model.ProductDetail, nameFile string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile(linkFile+nameFile+".json", file, 0644)
}

func GetDataAPI(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	// dataToStr := string(body[:])

	//var result map[string]interface{}
	return body

}
