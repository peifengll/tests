package api

import (
	"net/http"
	"testing"
)

func TestFlower_FlowerSortShow(t *testing.T) {
	client := &http.Client{}
	f := Flower{client}
	f.FlowerSortShow(nil)
}
