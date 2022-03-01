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

	// get the products
	data, err := client.GetProducts(cl.addr)
	if err != nil {
		http.Error(c.Writer, "Error getting products", http.StatusInternalServerError)
	}

	// read the config file
	conf, err := config.ReadConfig()
	if err != nil {
		cl.l.Println("[ERROR] No config file found.")
	}
	// update the prices if is the first day of the month

	if config.IsDayOfMonth(1) && !conf.UpdatedThisMonth {
		for _, d := range data {
			cl.l.Printf("[INFO] Updating price of product %s\nActual price: %f\n", d.Name, d.Price)
			d.Price = float32(math.Round(float64(d.Price + (d.Price * conf.MensualPerc))))
			cl.l.Printf("[INFO] New price: %f\n", d.Price)
			client.UpdateProduct(cl.addr, d.ID, d)
		}
		conf.UpdatedThisMonth = true
		conf.MensualPerc = conf.MensualPerc
		config.WriteConfig(conf)
		cl.l.Println("[INFO] Already updated this month.")
	}

	// read the config again to watch if the UpdatedThisMonth flag was change
	conf, err = config.ReadConfig()
	if err != nil {
		cl.l.Println("[ERROR] No config file found.")
	}

	if config.IsDayOfMonth(2) && conf.UpdatedThisMonth {
		conf.UpdatedThisMonth = false
		conf.MensualPerc = conf.MensualPerc
		config.WriteConfig(conf)
	}

	// update the data to show it
	data, err = client.GetProducts(cl.addr)
	if err != nil {
		http.Error(c.Writer, "Error getting products", http.StatusInternalServerError)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": data,
	})
}
