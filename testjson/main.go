package main

import (
	"encoding/json"
	"fmt"
)

type Corp struct {
	ID  int32  `json:"id" gorm:"Column:id;type:int;not null;primaryKey:KEY"`
	Cid string `json:"cid,omitempty" gorm:"Column:cid;type:varchar(32);not null;default:'';uniqueIndex:KEY"` // Cid 集团用户注册登录时的用户名

}

func main() {
	var p Corp
	marshal, err := json.Marshal(p)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
