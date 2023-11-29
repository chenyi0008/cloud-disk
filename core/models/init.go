package models

import (
	"cloud-disk/core/define"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()
var RDB = InitRedis()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     define.RedisAddr,
		Password: define.RedisPassword, // no password set
		DB:       3,                    // use default DB
	})
	return rdb
}

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}

	return engine
}
