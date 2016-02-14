package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"mydb"
)

type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

type OnlineUser struct {
	User      []*Person
	LoginTime string
}

type U struct {
	Account string `json:"account"`
	Pwd     string    `json:"pwd"`
}

type R struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type RE struct {
	Code int64 `json:"code"`
	Msg  string`json:"msg"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dux := Person{
		Name: "zoro",
		Age: 27,
		Emails: []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role: "SE"}

	ch := Person{Name: "chxd", Age: 27, Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&dux, &ch}}

	//t := template.New("Person template")
	//t, err := t.Parse(templ)
	t, err := template.ParseFiles("tmpl.html")
	checkError(err)

	err = t.Execute(w, onlineUser)
	checkError(err)
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pwd := r.FormValue("pwd")
	if user == "" {
		ReturnEFormat(w, 1, "user is null")
		return
	}
	if pwd == "" {
		ReturnEFormat(w, 1, "pwd is null")
		return
	}
	var dddd U
	db := mydb.GetDbCollection("local", "user")
	err := db.Find(bson.M{"account":user, "pwd":pwd}).One(&dddd)
	if err != nil {
		ReturnEFormat(w, 1, err.Error())
		return
	}
	ReturnFormat(w, 0, dddd, "SUCCESS")
}

func Register(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pwd := r.FormValue("pwd")
	if user == "" {
		ReturnEFormat(w, 1, "user is null")
		return
	}
	if pwd == "" {
		ReturnEFormat(w, 1, "pwd is null")
		return
	}
	db := mydb.GetDbCollection("local", "user")
	var dddd []U
	err := db.Find(bson.M{"account":user, "pwd":pwd}).All(&dddd)
	if err != nil {
		ReturnEFormat(w, 1, err.Error())
		return
	}
	if len(dddd) != 0 {
		ReturnEFormat(w, 1, "user is already exist")
		return
	}
	err = db.Insert(bson.M{"account":user, "pwd":pwd, "status":"N"})
	if err != nil {
		ReturnEFormat(w, 1, err.Error())
		return
	}
	ReturnFormat(w, 0, nil, "SUCCESS")
}

const URLs = "127.0.0.1:27017/local"

func main() {
	mydb.InitDB(URLs)
	http.HandleFunc("/", Handler)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.ListenAndServe(":8888", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func ReturnFormat(w http.ResponseWriter, code int64, data interface{}, msg string) {
	res := R{Code:code, Data:data, Msg:msg}
	omg, _ := json.Marshal(res)
	w.Write(omg)
}

func ReturnEFormat(w http.ResponseWriter, code int64, msg string) {
	res := RE{Code:code, Msg:msg}
	omg, _ := json.Marshal(res)
	w.Write(omg)
}
