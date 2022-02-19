package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func AddProduct(addr string, prod Product) (int, error) {
	postBody, _ := json.Marshal(prod)

	resBody := bytes.NewBuffer(postBody)
	res, err := http.Post(addr, "application/json", resBody)
	if err != nil {
		return res.StatusCode, err
	}
	defer res.Body.Close()
	return res.StatusCode, nil
}
