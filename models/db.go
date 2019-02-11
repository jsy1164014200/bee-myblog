package models

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/mgo.v2"
	"time"
)

var mongoDialInfo = &mgo.DialInfo{
	Addrs: []string{
		beego.AppConfig.String("mongodbHost"),
	},
	Database: "admin",
	Username: beego.AppConfig.String("mongodbUser"),
	Password: beego.AppConfig.String("mongodbPassword"),
	Timeout:  60 * time.Second,
}

// 记得用完关掉
func GetDB() *mgo.Session {
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	return session
}

var redisHost = beego.AppConfig.String("redisHost") + ":6379"

// 记得用完关掉
func GetRedis() redis.Conn {
	conn, err := redis.Dial("tcp", redisHost)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	return conn
}
