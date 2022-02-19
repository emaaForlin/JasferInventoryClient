package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOneProduct(addr string, id int) ([]Product, error) {
	var prods []Product
	res, err := http.Get(fmt.Sprintf("%s/%d", addr, id))
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
