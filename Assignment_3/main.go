package main

import "Hacktiv8/Golang/Assignment/Assignment_3/router"

func main() {
	r := router.StartApp()
	r.Run(":8080")
}
