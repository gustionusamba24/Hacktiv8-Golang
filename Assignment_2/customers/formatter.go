package customers

type CustomerFormatter struct {
	ID int `json:"customer id"`
	Customer_Name string `json:"customer name"`
}

func CustomerFormat(customer Customer) CustomerFormatter {
	customerFormatter := CustomerFormatter{}

	customerFormatter.ID = customer.ID
	customerFormatter.Customer_Name = customer.Customer_Name

	return customerFormatter
}