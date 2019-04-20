/*
@Time : 2019/4/20 11:38 
@Author : Tian Kejie.
*/
package main

import (
	"Golang-RESTful/library/my_log"
	"Golang-RESTful/my_route"
	"net/http"
)

func main() {
	// 初始化日志
	wLog()
	router := my_route.NewRouter()
	panic(http.ListenAndServe(":8080", router))
}

func wLog() {
	c := make(map[string]string, 8)
	c["log_path"] = "./runtime/logs/"
	c["log_name"] = "access"
	c["log_chan_size"] = "50000" //chan size
	c["log_split_type"] = "size"
	c["log_split_size"] = "104857600" // 100MB
	err := my_log.InitLog("file", c)

	if err != nil {
		return
	}
}
