package main

import (
	"assignment_2/customers"
	"assignment_2/orders"
	"assignment_2/handler"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main(){
	dsn := "user:@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	customerRepository := customers.NewRepository(db)
	orderRepository := orders.NewRepository(db)

	customerService := customers.NewCustomerService(customerRepository)
	orderService := orders.NewService(orderRepository)

	customerHandler := handler.NewCustomerHandler(customerService)
	orderHandler := handler.NewOrderHandler(orderService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.POST("/order", orderHandler.CreateOrder)
	v1.POST("/customer", customerHandler.CreateCustomer)
	v1.GET("/orders", orderHandler.GetOrders)
	v1.PUT("/order/:id", orderHandler.UpdateOrder)
	v1.DELETE("/delete/:id", orderHandler.DeleteOrder)
	router.Run()
}