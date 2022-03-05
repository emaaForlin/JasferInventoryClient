package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetProducts(addr string, apikey string) ([]Product, error) {
	var prods []Product
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	req.Header.Set("apikey", apikey)
	if err != nil {
		return nil, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
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
