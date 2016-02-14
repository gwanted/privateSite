package mydb
import (
	"gopkg.in/mgo.v2"
	"fmt"
)

var Session *mgo.Session

func InitDB(urls string) (error) {
	sessions, err := mgo.Dial(urls) //连接服务器
	if err != nil {
		panic(err)
	}
	Session = sessions
	fmt.Printf("mongodb init finish")
	return nil
}

func GetDbCollection(db, collection string) (*mgo.Collection) {
	return Session.DB(db).C(collection)
}