package common

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReturnResult(w http.ResponseWriter, code int, msg string, data interface{}) {
	result := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	res, err := json.Marshal(result)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	w.Write(res)
}
