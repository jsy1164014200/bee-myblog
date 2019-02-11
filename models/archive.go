package models

import "gopkg.in/mgo.v2/bson"

type Archive struct {
	Id    bson.ObjectId   `bson:"_id"`
	Name  string          `bson:"name"`
	Blogs []bson.ObjectId `bson:"blogs"`
}
