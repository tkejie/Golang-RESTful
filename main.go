/*
@Time : 2019/4/19 15:57 
@Author : Tian Kejie.
*/
package main

import (
	"Golang-RESTful/library/my_config"
	"Golang-RESTful/library/my_log"
	"Golang-RESTful/library/my_sql"
	"fmt"
)

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	c := make(map[string]string, 8)
	c["log_path"] = "./runtime/logs/"
	c["log_name"] = "server"
	c["log_chan_size"] = "50000" //chan size
	c["log_split_type"] = "size"
	c["log_split_size"] = "104857600" // 100MB
	err := my_log.InitLog("file", c)

	if err != nil {
		return
	}

	my_config.InitConfig("./config.conf")
	config := my_config.Run()

	m := my_sql.NewConnector(config)
	//a := c.Table("test")

	res := m.Table("test").Where("id", 1).Get()

	fmt.Println(res.ToStringMapList())
}
