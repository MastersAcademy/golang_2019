package main

import (
	"simpleAPI/httpd/handler"
	"simpleAPI/platform/order"

	"github.com/gin-gonic/gin"
)

func main() {
	orders := order.New()
	r := gin.Default()

	r.GET("/order", handler.AllOrdersGet(orders))
	r.POST("/order", handler.AllOrdersPost(orders))

	r.GET("/order/:id", handler.OrderGet(orders))

	r.DELETE("/order/:id", handler.OrderDelete(orders))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
