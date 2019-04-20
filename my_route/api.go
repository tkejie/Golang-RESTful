/*
@Time : 2019/4/20 11:28 
@Author : Tian Kejie.
*/
package my_route

import (
	"Golang-RESTful/library/my_config"
	"Golang-RESTful/library/my_log"
	"Golang-RESTful/library/my_sql"
	"Golang-RESTful/models"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	my_config.InitConfig("./config.conf")
	config := my_config.Run()

	m := my_sql.NewConnector(config)

	res := m.Table("test").Where("name", "宁").Count("id", "total")

	lists := models.Lists{
		models.Index{Id: 1, Name: "Kejie", Data: res.ToStringMapList()},
	}

	if err := json.NewEncoder(w).Encode(lists); err != nil {
		my_log.Error("%v", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	my_config.InitConfig("./config.conf")
	config := my_config.Run()

	m := my_sql.NewConnector(config)

	res := m.Table("test").Where("name", "宁").Count("id", "total")

	json.NewEncoder(w).Encode(res.ToStringMapList())

}
