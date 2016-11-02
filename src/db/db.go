package db

import (
	"fmt"
	"github.com/gwanted/mgo.v2"
)

const URL = "127.0.0.1:27017"

func Exec(name string, f func(*mgo.Collection)) {
	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			fmt.Printf("MongoDB recover: %s\n", err)
		}
	}()
	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(4096)

	db := session.DB("privateSite").C(name)
	f(db)
}
