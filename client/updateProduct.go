package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateProduct(addr string, id int, p Product, apikey string) (int, error) {
	p.ID = id
	jsonBlob, err := json.Marshal(p)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", addr, id), bytes.NewBuffer(jsonBlob))
	if err != nil {
		return 0, nil
	}

	req.Header.Set("Content-Type", "Application/json; charset=utf-8")
	req.Header.Set("apikey", apikey)
	res, err := httpClient.Do(req)
	if err != nil {
		return res.StatusCode, err
	}
	return res.StatusCode, nil

}
