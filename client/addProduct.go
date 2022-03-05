package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func AddProduct(addr string, prod Product, apikey string) (int, error) {
	postBody, _ := json.Marshal(prod)
	resBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, addr, resBody)
	req.Header.Set("apikey", apikey)
	if err != nil {
		return 0, nil
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}
