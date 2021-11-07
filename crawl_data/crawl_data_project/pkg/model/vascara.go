package model

type ProductDetail struct {
	ID           uint   `json:"id" bson:"_id,omitempty"`
	ProductCode  string `json:"productcode"`
	Title        string `json:"title"`
	CategoryName string `json:"categoryname"`
	CategoryId   uint   `json:"categoryid"`
	Image        string `json:"image"`
	Discount     string `json:"discount"`
	Price        int64  `json:"price"`
	Currency     string `json:"currency"`
	BestSale     bool   `json:"bestsale"`
	Infor        string `json:"infor"`
	BaseModel
}

type CategoryChild struct {
	ID             uint            `json:"id"`
	IdFa           uint            `json:"idfa"`
	Name           string          `json:"title"`
	Link           string          `json:"link"`
	LinkCategory   string          `json:"linkcategory"`
	ProductDetails []ProductDetail `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
type CategoryFa struct {
	ID             uint            `json:"id"`
	Name           string          `json:"name"`
	Link           string          `json:"link"`
	LinkCategory   string          `json:"linkcategory"`
	CategoryChilds []CategoryChild `gorm:"foreignKey:IdFa;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"categorychild"`
	BaseModel
}

type ProductView struct {
	Image      string `json:"image"`
	Price      string `json:"price"`
	Currency   string `json:"currency"`
	Title      string `json:"title"`
	Detaillink string `json:"detaillink"`
}

//* get data from api type html
type Source struct {
	Html string `json:"html"`
}
