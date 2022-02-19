package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryClient/client"
	"github.com/gin-gonic/gin"
)

func (cl *Client) DeleteProduct(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(c.Writer, "ID must to be an integer", http.StatusBadRequest)
	}
	_, err = client.DeleteProduct(cl.addr, id)
	if err != nil {
		http.Error(c.Writer, "Something was wrong, maybe this ID does not exists", http.StatusBadRequest)
	}
}
