package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Employee struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var employeeDB map[string]*Employee

func init() {
	employeeDB = map[string]*Employee{}
	employeeDB["Mike"] = &Employee{"e-1","Mike",35}
	employeeDB["Rose"] = &Employee{"e-2","Rose",24}
}


func main() {
	router := httprouter.New()
	router.GET("/",Index)
	router.GET("/employees/:name",GetEmployeeByName)
	http.ListenAndServe(":8080",router)
}

func Index(w http.ResponseWriter,r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w,"welcome ! \n")
}

func GetEmployeeByName(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	qName := ps.ByName("name")
	var (
		ok bool
		info *Employee
		infoJson []byte
		err error
	)
	if info,ok = employeeDB[qName]; !ok{
		w.Write([]byte("error not found"))
		return
	}
	if infoJson, err = json.Marshal(info);err != nil {
		w.Write([]byte(fmt.Sprintf("%s",err)))
		return
	}
	w.Write(infoJson)
}
