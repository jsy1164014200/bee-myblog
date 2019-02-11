package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Blog struct {
	Id           bson.ObjectId   `bson:"_id"`
	CreatedAt    time.Time       `bson:"createdAt"`
	UpdatedAt    time.Time       `bson:"updatedAt"`
	Title        string          `bson:"title"`
	Summary      string          `bson:"summary"`
	CommentCount int             `bson:"commentCount"`
	Comments     []bson.ObjectId `bson:"comments"`
	ReadCount    int             `bson:"readCount"`
	Tags         []string        `bson:"tags"`
}
