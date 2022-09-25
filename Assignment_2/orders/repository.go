package orders

import (
	"assignment_2/customers"
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ID int) (customers.Customer, error)
	CreateOrder(order Order) (Order, error)
	GetAllOrders() ([]Order, error)
	Update(order Order) (Order, error)
	GetOrderByID(ID int) (Order, error)
	Delete(order Order) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repository *repository) FindByID(ID int) (customers.Customer, error) {
	var customer customers.Customer

	// err := r.db.Where("ID = ?", id).Find(&customer).Error
	err := repository.db.First(&customer, "id = ?", ID).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (repository *repository) CreateOrder(order Order) (Order, error) {

	err := repository.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (repository *repository) GetAllOrders() ([]Order, error) {
	var order []Order

	err := repository.db.Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (repository *repository) GetOrderByID(ID int) (Order, error) {
	var order Order
	err := repository.db.Find(&order, ID).Error
	if err != nil {
		return order, err
	}

	return order, nil

}

func (repository *repository) Update(order Order) (Order, error) {

	err := repository.db.Debug().Save(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (repository *repository) Delete(order Order) (Order, error) {
	err := repository.db.Delete(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}