package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

const dsn = "pabb:pabb@(localhost:3306)/qjcg?charset=utf8mb4&parseTime=True&loc=Local"

/*
*
测试 在 defer里边进行commit 可行否

可以
*/
func TestTxCommit(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
		err = tx.Commit().Error
		log.Println(err)
	}()
	tx.Model(&Txtest{}).Create(&Txtest{
		Name: "666",
	})
	tx.Model(&Txtest{}).Create(&Txtest{
		Name: "666",
	})
	tx.Model(&Txtest{}).Create(&Txtest{
		Name: "666",
	})
	tx.Model(&Txtest{}).Create(&Txtest{
		Name: "666",
	})
	tx.Model(&Txtest{}).Create(&Txtest{
		Name: "666",
	})

}

type Txtest struct {
	Id   int    `json:"id"`
	Name string `json:"Txtest"`
}

func (c *Txtest) TableName() string {
	return "txtest"
}

func TestCommitT(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tx := db.WithContext(context.Background()).Begin()
	defer func() {
		err = tx.Commit().Error
		if err != nil {
			log.Fatalln(err)
		}
	}()

	time.Sleep(8 * time.Second)

}

func TestJIU(t *testing.T) {
	var tt any
	tt = 6
	defer func() {
		fmt.Println(tt)
	}()
	tt = 999
}
