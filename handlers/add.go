package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryClient/client"
	"github.com/gin-gonic/gin"
)

func (cl *Client) Add(c *gin.Context) {
	cl.l.Println("Add page")
	c.HTML(http.StatusOK, "add.html", nil)
}

func (cl *Client) AddPost(c *gin.Context) {
	cl.l.Println("Add Page POST")
	pPrice, _ := strconv.Atoi(c.PostFormArray("prod-price")[0])
	var p = client.Product{
		Name:        c.PostFormArray("prod-name")[0],
		Description: c.PostFormArray("prod-description")[0],
		Price:       float32(pPrice),
		//SKU:         c.PostFormArray("prod-sku")[0],
	}
	cl.l.Printf("Adding %#v", p)
	status, err := client.AddProduct(cl.addr, p)
	cl.l.Println(status, err)
	cl.l.Println("Add page")
	c.HTML(http.StatusOK, "add.html", nil)
}
