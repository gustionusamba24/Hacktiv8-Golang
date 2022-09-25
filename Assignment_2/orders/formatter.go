package orders 

import "time"

type OrderFormatter struct {
	Order_ID int `json:"order id"`
	Customer_ID int `json:"customer id"`
	OrderedAt time.Time `json:"orderedAt"`
	Items ItemsFormatter `json:"items"`
}

type ItemsFormatter struct {
	Item_Code string `json:"item code"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}

func OrderFormat(order Order) OrderFormatter {
	orderFormat := OrderFormatter{}

	orderFormat.OrderedAt = order.OrderedAt
	orderFormat.Order_ID = order.ID
	orderFormat.Customer_ID = order.Customer_ID

	ItemsFormatter := ItemsFormatter{}
	ItemsFormatter.Item_Code = order.Item_Code
	ItemsFormatter.Description = order.Description
	ItemsFormatter.Quantity = order.Quantity

	orderFormat.Items = ItemsFormatter

	return orderFormat
}

func OrdersFormat(orders []Order) []OrderFormatter {
	ordersFormatter := []OrderFormatter{}

	for _, order := range orders {
		orderFormatter := OrderFormat(order)
		ordersFormatter = append(ordersFormatter, orderFormatter)
	}

	return ordersFormatter
}