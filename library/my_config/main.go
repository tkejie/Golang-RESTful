/*
@Time : 2019/4/18 15:24
@Author : Tian Kejie.
*/
package my_config

import (
	"fmt"
	"sync/atomic"
	"time"
)

type AppConfig struct {
	hostname string
	port     int
}

type AppConfigMgr struct {
	config atomic.Value
}

var appConfigMgr = &AppConfigMgr{}

func (a *AppConfigMgr) Callback(conf *Config) {
	appConfig := &AppConfig{}
	hostname, err := conf.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err: %v\n", err)
		return
	}
	appConfig.hostname = hostname
	port, err := conf.GetInt("port")
	if err != nil {
		fmt.Printf("get port err: %v\n", err)
		return
	}
	appConfig.port = port
	appConfigMgr.config.Store(appConfig)
}

func InitConfig(file string) {
	conf, err := NewConfig(file)
	if err != nil {
		fmt.Printf("read config file err: %v\n", err)
		return
	}

	conf.AddObserver(appConfigMgr)

	var appConfig AppConfig
	appConfig.hostname, err = conf.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err: %v\n", err)
		return
	}
	fmt.Println("Hostname:", appConfig.hostname)

	appConfig.port, err = conf.GetInt("port")
	if err != nil {
		fmt.Printf("get port err: %v\n", err)
		return
	}
	fmt.Println("Port:", appConfig.port)

	appConfigMgr.config.Store(&appConfig)
	fmt.Println("first load success.")
}

func Run() {
	for {
		appConfig := appConfigMgr.config.Load().(*AppConfig)

		fmt.Println("Hostname:", appConfig.hostname)
		fmt.Println("Port:", appConfig.port)
		fmt.Printf("%v\n", "------------")
		time.Sleep(5 * time.Second)

	}
}

//func main() {
//	confFile := "./config.conf"
//	initConfig(confFile)
//	run()
//}
