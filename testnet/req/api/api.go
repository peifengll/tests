package api

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

var BaseUrl = `http://peifeng.site:8000/`

func ParseResponseByMqp(response *http.Response) (map[string]any, error) {
	var res map[string]any
	body, err := io.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &res)
	}
	return res, err
}

func ParseResponseGJson(response *http.Response) (gjson.Result, error) {
	var parse gjson.Result
	body, err := io.ReadAll(response.Body)
	if err == nil {
		parse = gjson.Parse(string(body))
	}
	return parse, err
}
