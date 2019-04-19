/*
@Time : 2019/4/19 15:04 
@Author : Tian Kejie.
*/
package main

import "Golang-RESTful/library/my_config"

func main() {
	my_config.InitConfig("./config.conf")
	my_config.Run()
}
