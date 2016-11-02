package model

import (
	"../db"
	"github.com/gwanted/mgo.v2"
	"github.com/gwanted/mgo.v2/bson"
	"time"
)

type Article struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Picture   string        `bson:"picture" json:"picture"`
	Describe  string        `bson:"describe" json:"describe"`
	Author    User          `bson:"author" json:"author"`
	Contents  []Content     `bson:"content" json:"content"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

type Content struct {
	Picture  string `bson:"picture" json:"picture"`
	Describe string `bson:"describe" json:"describe"`
}

func (m *Article) CollectionName() string {
	return ArticleCollection()
}

func ArticleCollection() string {
	return "article"
}

func FindArticle(condition bson.M) (result *Article, err error) {
	db.Exec(ArticleCollection(), func(c *mgo.Collection) {
		err = c.Find(condition).One(&result)
	})
	return
}
