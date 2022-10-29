package handler

import (
	"Hacktiv8/Golang/Assignment/Assignment_3/weather"
	"fmt"
	"github.com/gin-gonic/gin"
	"text/template"
)

func GetStatusHandler(c *gin.Context) {
	statusValue := weather.GetStatusValue()

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(c.Writer, statusValue)
}
