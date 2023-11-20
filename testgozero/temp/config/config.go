package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type TomConfig struct {
	AppName    string
	Mysql      MySQLConfig
	Redis      RedisConfig
	CacheRedis CacheRedisConfig
}

type MySQLConfig struct {
	Host        string
	DbName      string
	Password    string
	Port        int
	TablePrefix string
	User        string
}
type RedisConfig struct {
	Host string
	Type string
	Pass string
	Key  string
}

type CacheRedisConfig struct {
	Host string
	Type string
	Pass string
}

var c *TomConfig

func init() {
	fmt.Println("Init Starting")
	InitConfig()
	fmt.Println("Init ending")

}

func InitConfig() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("D:\\code\\go\\trys\\tests\\testgozero\\temp\\config")
	err := viper.ReadInConfig()
	if nil != err {
		return err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	return nil
}
func GetConfig() (*TomConfig, error) {
	if c == nil {
		fmt.Println("config在GetConfig中初始化")
		err := InitConfig()
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}
