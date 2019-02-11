package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Collection struct {
	Id     bson.ObjectId `bson:"_id"`
	Title  string        `bson:"title"`
	Author string        `bson:"author"`
	Url    string        `bson:"url"`
	Time   time.Time     `bson:"time"`
}
