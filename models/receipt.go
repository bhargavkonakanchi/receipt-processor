package models

// Receipt structure matches the expected schema for a receipt
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item represents an item on the receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
