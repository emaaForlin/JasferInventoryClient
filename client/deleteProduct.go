package client

import (
	"bytes"
	"fmt"
	"net/http"
)

func DeleteProduct(addr string, id int) (int, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%d", addr, id), bytes.NewBuffer([]byte{}))
	if err != nil {
		return 0, nil
	}

	req.Header.Set("Content-Type", "Application/json; charset=utf-8")

	res, err := httpClient.Do(req)
	if err != nil {
		return res.StatusCode, err
	}
	return res.StatusCode, nil
}
