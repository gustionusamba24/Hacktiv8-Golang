package customers

type CustomerFormatter struct {
	ID int `json:"customer id"`
	Customer_Name string `json:"customer name"`
}

func CustomerFormat(customer Customer) CustomerFormatter {
	customerFormat := CustomerFormatter{}

	customerFormat.ID = customer.ID
	customerFormat.Customer_Name = customer.Customer_Name

	return customerFormat
}