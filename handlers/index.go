package handlers

import (
	"net/http"

	"github.com/emaaForlin/JasferInventoryClient/client"
	"github.com/gin-gonic/gin"
)

func (cl *Client) Index(c *gin.Context) {
	cl.l.Println("Index page")
	data, err := client.GetProducts(cl.addr)
	if err != nil {
		http.Error(c.Writer, "Error getting products", http.StatusInternalServerError)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": data,
	})
}
