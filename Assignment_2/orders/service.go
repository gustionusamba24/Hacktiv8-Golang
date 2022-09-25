package orders

type Service interface {
	Create(input OrderInput) (Order, error)
	GetOrders() ([]Order, error)
	Update(inputID GetOrderByID, inputData OrderUpdate) (Order, error)
	GetOrderByID(inputID GetOrderByID) (Order, error)
	Delete(inputID int) (Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input OrderInput) (Order, error) {
	createOrder := Order{}

	createOrder.Customer_ID = input.Customer_id
	createOrder.Item_code = input.Items_code
	createOrder.Description = input.Description
	createOrder.Quantity = input.Quantity

	check, err := s.repository.FindByID(createOrder.Customer_ID)
	if check.ID == 0 {
		return createOrder, err
	}

	newOrder, err := s.repository.CreateOrder(createOrder)
	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (s *service) GetOrders() ([]Order, error) {

	orders, err := s.repository.GetAllOrders()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *service) GetOrderByID(inputID GetOrderByID) (Order, error) {
	order, err := s.repository.GetOrderByID(inputID.Order_ID)
	if err != nil {
		return order, err
	}

	return order, nil
}


func (s *service) Update(inputID GetOrderByID, inputData OrderUpdate) (Order, error) {

	order, err := s.repository.GetOrderByID(inputID.Order_ID)
	if err != nil {
		return order, err
	}

	order.Item_Code = inputData.Items_Code
	order.Description = inputData.Description
	order.Quantity = inputData.Quantity
	order.Customer_ID = inputData.Customer_ID

	updateOrder, err := s.repository.Update(order)
	if err != nil {
		return updateOrder, err
	}

	return updateOrder, nil
}

func (s *service) Delete(inputID int) (Order, error) {
	order, err := s.repository.GetOrderByID(inputID)
	if err != nil {
		return order, err
	}

	deleteOrder, err := s.repository.Delete(order)
	if err != nil {
		return deleteOrder, err
	}
	return deleteOrder, nil
}
