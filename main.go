/*
@Time : 2019/4/18 15:24 
@Author : Tian Kejie.
*/
package main

import (
	"Golang-RESTful/library/my_log"
)

func initLog(logPath, logName string) {
	config := make(map[string]string, 8)
	config["log_path"] = logPath
	config["log_name"] = logName
	config["log_chan_size"] = "50000" //chan size 可以不用
	config["log_split_type"] = "size"
	config["log_split_size"] = "104857600" // 100MB

	err := my_log.InitLog("file", config)
	if err != nil {
		return
	}
}

func Run() {
	for {
		my_log.Info("This is info.")
		my_log.Debug("This is info.")
		my_log.Warn("This is info.")
		my_log.Error("This is info.")
	}
}

func main() {
	initLog("./runtime/logs/", "server")
	Run()
}
