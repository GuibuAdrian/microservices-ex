package domain

type Order struct {
	Id          uint64 `json:"id"`
	PersonName  string `json:"person_name"`
	ProductName string `json:"product_name"`
}
