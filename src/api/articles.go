package api

import (
	"net/http"
	//"io"
	"../common"
	"../model"
	"fmt"
	"github.com/gwanted/mgo.v2/bson"
)

func GetArticle(w http.ResponseWriter, req *http.Request) {
	article, err := model.FindArticle(bson.M{"_id": bson.ObjectIdHex("581870463949ee77418636ad")})
	if err != nil {
		fmt.Println(err.Error())
	}
	common.ReturnResult(w, 200, "success", article)
	return
}
