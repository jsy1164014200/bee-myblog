package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Comment struct {
	Id        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"username"`
	Content   string        `bson:"content"`
	CreatedAt time.Time     `bson:"createdAt"`
}
