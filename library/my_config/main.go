/*
@Time : 2019/4/18 15:24
@Author : Tian Kejie.
*/
package my_config

import (
	"Golang-RESTful/library/my_log"
	"sync/atomic"
)

type AppConfig struct {
	dbUsername string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	dbCharset  string
}

type AppConfigMgr struct {
	config atomic.Value
}

var appConfigMgr = &AppConfigMgr{}

func (a *AppConfigMgr) Callback(conf *Config) {
	appConfig := &AppConfig{}
	username, err := conf.GetString("db_username")
	if err != nil {
		my_log.Error("get db_username err: %v\n", err)
		return
	}
	appConfig.dbUsername = username

	dbPass, err := conf.GetString("db_password")
	if err != nil {
		my_log.Error("get db_password err: %v\n", err)
		return
	}
	appConfig.dbPassword = dbPass

	dbHost, err := conf.GetString("db_host")
	if err != nil {
		my_log.Error("get db_host err: %v\n", err)
		return
	}
	appConfig.dbHost = dbHost

	dbPort, err := conf.GetString("db_port")
	if err != nil {
		my_log.Error("get db_port err: %v\n", err)
		return
	}
	appConfig.dbUsername = dbPort

	dbName, err := conf.GetString("db_name")
	if err != nil {
		my_log.Error("get db_name err: %v\n", err)
		return
	}
	appConfig.dbName = dbName

	dbCharset, err := conf.GetString("db_charset")
	if err != nil {
		my_log.Error("get db_charset err: %v\n", err)
		return
	}
	appConfig.dbCharset = dbCharset
}

func InitConfig(file string) {
	conf, err := NewConfig(file)
	if err != nil {
		my_log.Error("read config file err: %v\n", err)
		//my_log.Error("read config file err: %v\n", err)
		return
	}

	conf.AddObserver(appConfigMgr)

	var appConfig AppConfig
	appConfig.dbUsername, err = conf.GetString("db_username")
	if err != nil {
		my_log.Error("get db_username err: %v\n", err)
		return
	}

	appConfig.dbPassword, err = conf.GetString("db_password")
	if err != nil {
		my_log.Error("get db_password err: %v\n", err)
		return
	}

	appConfig.dbHost, err = conf.GetString("db_host")
	if err != nil {
		my_log.Error("get db_host err: %v\n", err)
		return
	}

	appConfig.dbPort, err = conf.GetString("db_port")
	if err != nil {
		my_log.Error("get db_port err: %v\n", err)
		return
	}

	appConfig.dbName, err = conf.GetString("db_name")
	if err != nil {
		my_log.Error("get db_name err: %v\n", err)
		return
	}

	appConfig.dbCharset, err = conf.GetString("db_charset")
	if err != nil {
		my_log.Error("get db_charset err: %v\n", err)
		return
	}

	appConfigMgr.config.Store(&appConfig)
	my_log.Info("first load success.")
}

//func Run()  {
	//for {
	//	appConfig := appConfigMgr.config.Load().(*AppConfig)
	//
	//	fmt.Println("Hostname:", appConfig.hostname)
	//	fmt.Println("Port:", appConfig.port)
	//	my_log.Error("%v\n", "------------")
	//	time.Sleep(5 * time.Second)
	//
	//}

//}

func Run() (c map[string]string)  {
	appConfig := appConfigMgr.config.Load().(*AppConfig)
	config := make(map[string]string)
	config["db_username"] = appConfig.dbUsername
	config["db_password"] = appConfig.dbPassword
	config["db_host"] = appConfig.dbHost
	config["db_port"] = appConfig.dbPort
	config["db_name"] = appConfig.dbName
	config["db_charset"] = appConfig.dbCharset

	return config
}

//func main() {
//	confFile := "./config.conf"
//	initConfig(confFile)
//	run()
//}
