package db

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/peifengll/tests/testgasbin/self_demo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var once sync.Once

func GetOnceDb() *gorm.DB {
	once.Do(func() {
		c, err := config.GetConfig()
		if err != nil {
			log.Errorf("配置信息获取错误:%s", err)
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s",
			c.Mysql.User,
			c.Mysql.Password,
			c.Mysql.Host,
			c.Mysql.Port,
			c.Mysql.DbName,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})

	return db

}
