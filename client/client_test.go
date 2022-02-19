package client

import (
	"fmt"
	"testing"
)

var addr string = "http://localhost:9090/products"

func testGetOneProduct(t *testing.T) {
	prods, err := GetOneProduct(addr, 4)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prods)
	t.Fail()
}
