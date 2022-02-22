package handlers

import (
	"math"
	"net/http"

	"github.com/emaaForlin/JasferInventoryClient/client"
	"github.com/emaaForlin/JasferInventoryClient/config"
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
	if config.IsFirstDayOfMonth() {
		conf, err := config.ReadConfig()
		if err != nil {
			cl.l.Println("No config file found.")
			return
		}
		for _, d := range data {
			cl.l.Printf("Updating price of product %s\nActual price: %f\n", d.Name, d.Price)
			d.Price = float32(math.Round(float64(d.Price + (d.Price * conf.MensualPerc))))
			cl.l.Printf("New price: %f\n", d.Price)
			client.UpdateProduct(cl.addr, d.ID, d)
		}
	}
}
