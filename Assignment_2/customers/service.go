package customers

type Service interface {
	CreateCustomer(input CustomerInput) (Customer, error)
}

type service struct {
	repository Repository
}

func NewCustomerService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCustomer(input CustomerInput) (Customer, error) {
	customerInput := Customer{}

	customerInput.Customer_Name = input.Customer_Name

	newCustomer, err := s.repository.Save(customerInput)
	if err != nil {
		return newCustomer, err
	}

	return newCustomer, nil
}