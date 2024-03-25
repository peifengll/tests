package main

import (
	"github.com/peifengll/tests/testnet/req/api"
	"net/http"
)

func main() {
	client := &http.Client{}
	o := api.Order{client}
	redata := map[string]any{
		"user_id":    85,
		"money":      78,
		"cart_id":    "[43]",
		"address_id": 3,
	}
	o.OrderAdd(redata)
}
