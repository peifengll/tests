package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	*http.Client
}

func (o *Order) OrderAdd(data map[string]any) {
	url := BaseUrl + "order/add"
	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := o.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()
	// 打印响应内容
	fmt.Println("Response:", resp.Status)
}
