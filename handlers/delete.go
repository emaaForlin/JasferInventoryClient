package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryClient/client"
	"github.com/gin-gonic/gin"
)

type Form struct {
	Id []string `form:"delete[]"`
}

func (cl *Client) DeleteProduct(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	status, err := client.DeleteProduct(cl.addr, id, cl.apikey)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if status == http.StatusOK {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
