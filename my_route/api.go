/*
@Time : 2019/4/20 11:28 
@Author : Tian Kejie.
*/
package my_route

import (
	"Golang-RESTful/library/my_log"
	"Golang-RESTful/models"
	"encoding/json"
	"net/http"
)

type MyStruct struct {
	Name string
	Age  int64
}

func BugList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	res := models.GetBugList(r.Form)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		my_log.Error("%v", err)
	}
}

func UpList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	res := models.GetUpdateList(r.Form)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		my_log.Error("%v", err)
	}

}

func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	res := models.BuildBugData(r.Form)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		my_log.Error("%v", err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	res := models.BuildUpdateData(r.Form)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		my_log.Error("%v", err)
	}
}
