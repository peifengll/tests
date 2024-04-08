package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Flower struct {
	*http.Client
}

func (o *Flower) FlowerSortShow(data map[string]any) {
	url := "http://peifeng.site:8000/flower/sort/?id=0041"
	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
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
	fmt.Println("开始那啥")
	response, err := ParseResponseGJson(resp)
	fmt.Println("那啥结束")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("err", err)
	fmt.Println(response.Get("data"))
	fmt.Println("err", err)

	fmt.Printf(":::")
}
