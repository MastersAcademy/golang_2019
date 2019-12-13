package handler

import (
	"net/http"
	"simpleAPI/platform/order"

	"github.com/gin-gonic/gin"
)

func OrderGet(order *order.AllOrders) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		res := order.GetById(id)

		c.JSON(http.StatusOK, res)
	}
}
