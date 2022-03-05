package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryClient/client"
	"github.com/emaaForlin/JasferInventoryClient/config"
	"github.com/gin-gonic/gin"
)

func (cl *Client) ConfigPage(c *gin.Context) {
	data, err := config.ReadConfig()
	if err != nil {
		data.MensualPerc = 0.0
	}
	data.MensualPerc = data.MensualPerc * 100
	c.HTML(http.StatusOK, "config.html", gin.H{
		"data": data,
	})
}

func (cl *Client) Config(c *gin.Context) {
	mPercentil, _ := strconv.Atoi(c.PostFormArray("mensual-percentile")[0])
	conf := config.Config{}
	conf.MensualPerc = float32(mPercentil) / 100
	config.WriteConfig(conf)
	conf, err := config.ReadConfig()
	if err != nil {
		cl.l.Println(err)
	}
	cl.l.Printf("Updating config => %#v", conf)
	fmt.Println(mPercentil)
	c.Redirect(http.StatusMovedPermanently, "/config")
}

func (cl *Client) UpdNow(c *gin.Context) {
	mPercentil, _ := strconv.ParseFloat(c.PostFormArray("mensual-percentile")[0], 64)
	data, err := client.GetProducts(cl.addr, cl.apikey)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	for _, d := range data {
		cl.l.Printf("Updating price of product %s\nActual price: %f\n", d.Name, d.Price)
		d.Price = float32(math.Round(float64(d.Price) + float64(d.Price*float32(mPercentil/100))))
		cl.l.Printf("New price: %f\n", d.Price)
		client.UpdateProduct(cl.addr, d.ID, d, cl.apikey)
	}
	c.Redirect(http.StatusMovedPermanently, "/config")
}
