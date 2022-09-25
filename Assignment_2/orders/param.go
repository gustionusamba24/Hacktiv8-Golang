package orders

type OrderInput struct {
	Customer_ID int `json:"customer_id"`
	Items_Code  string `json:"items_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type GetOrderByID struct {
	Order_ID int `uri:"id"`
}

type OrderUpdate struct {
	ID          int
	Customer_ID int    `json:"customer_id"`
	Items_Code  string `json:"items_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}