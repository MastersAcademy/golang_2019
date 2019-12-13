package handler

import (
	"simpleAPI/platform/order"

	"github.com/gin-gonic/gin"
)

func OrderDelete(order *order.AllOrders) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		order.DelById(id)
	}
}
