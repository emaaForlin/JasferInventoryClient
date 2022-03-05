package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryClient/client"
	"github.com/gin-gonic/gin"
)

func (cl *Client) EditPage(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(c.Writer, "ID must to be an integer", http.StatusBadRequest)
	}
	data, err := client.GetOneProduct(cl.addr, id, cl.apikey)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.HTML(http.StatusOK, "update.html", gin.H{
		"data": data,
	})
}

func (cl *Client) EditPost(c *gin.Context) {
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)
	pPrice, _ := strconv.Atoi(c.PostFormArray("prod-price")[0])
	var p = client.Product{
		ID:          id,
		Name:        c.PostFormArray("prod-name")[0],
		Description: c.PostFormArray("prod-description")[0],
		Price:       float32(pPrice),
		SKU:         c.PostFormArray("prod-sku")[0],
	}
	status, err := client.UpdateProduct(cl.addr, id, p, cl.apikey)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	if status == http.StatusOK {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
