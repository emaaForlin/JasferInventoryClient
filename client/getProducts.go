package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetProducts(addr string) ([]Product, error) {
	var prods []Product
	res, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	jsonBlob, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonBlob, &prods)
	if err != nil {
		return nil, err
	}
	return prods, nil
}
