package handler

import (
	"net/http"
	"simpleAPI/platform/order"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type allOrdersPostRequest struct {
	ID       string `json:"id"`
	Products string `json:"products"`
}

func AllOrdersPost(o *order.AllOrders) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody := allOrdersPostRequest{}
		c.Bind(&reqBody)

		item := order.Order{
			ID:       uuid.New().String(),
			Products: reqBody.Products,
		}

		if (item.ID != "") || (item.Products != "") {
			o.Add(item)
		}

		c.JSON(http.StatusOK, item.ID)

		c.Status(http.StatusNoContent)
	}
}
