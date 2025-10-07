package model

type PageData struct {
	NavbarID     string
	ImageCatName string
	CatName      string
	Cats         []Kitten
}

type Kitten struct {
	CatReferenceName string `json:"catReferenceName"`
	CatTitleName     string `json:"catTitleName"`
	Price            string `json:"price"`
	Description      string `json:"description"`
	DOB              string `json:"dob"`
	Litter           string `json:"litter"`
	Age              string `json:"age"`
	Color            string `json:"color"`
}
