package model

import (
	"../db"
	"github.com/gwanted/mgo.v2"
	"github.com/gwanted/mgo.v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Password  string        `bson:"password" json:"-"`
	CreatedAt time.Time     `bson:"createdAt" json:"-"`
}

func (m *User) CollectionName() string {
	return UserCollectionName()
}

func UserCollectionName() string {
	return "user"
}

func (m *User) Insert() (err error) {
	db.Exec(m.CollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}

func FindUser(condition bson.M) (result *User, err error) {
	db.Exec(UserCollectionName(), func(c *mgo.Collection) {
		err = c.Find(condition).One(&result)
	})
	return
}
