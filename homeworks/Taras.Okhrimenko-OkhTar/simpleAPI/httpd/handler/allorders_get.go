package handler

import (
	"net/http"
	"simpleAPI/platform/order"

	"github.com/gin-gonic/gin"
)

func AllOrdersGet(order *order.AllOrders) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := order.GetAll()
		c.JSON(http.StatusOK, res)
	}
}
