/*
@Time : 2019/4/20 11:30 
@Author : Tian Kejie.
*/
package models

import (
	"Golang-RESTful/library/my_config"
	"Golang-RESTful/library/my_log"
	"Golang-RESTful/library/my_sql"
	"strconv"
	"strings"
	"time"
)

type Index struct {
	Code  int                 `json:"code"`
	Msg   string              `json:"msg"`
	Total int                 `json:"total"`
	Data  []map[string]string `json:"data"`
}

type Bug struct {
	Id           int    `json:"id"`
	BugInfo      string `json:"bug_info"`
	BugLevel     string `json:"bug_level"`
	BugFindTime  string `json:"bug_find_time"`
	BugEntrustor string `json:"bug_entrustor"`
	BugRemarks   string `json:"bug_remarks"`
}

type Return struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Lists []Index
type Bugs []Bug

/**
  插入Bug表
 */
func BuildBugData(data map[string][]string) Return {
	if _, ok := data["bug_info"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_info",
		}
	}
	if _, ok := data["bug_level"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_level",
		}
	}
	if _, ok := data["bug_find_time"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_find_time",
		}
	}
	if _, ok := data["bug_entrustor"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_entrustor",
		}
	}
	if _, ok := data["bug_remarks"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_remarks",
		}
	}

	my_config.InitConfig("./config.conf")
	config := my_config.Run()

	m := my_sql.NewConnector(config)

	a := make(map[string]interface{})

	a["bug_info"] = strings.Join(data["bug_info"], "")
	a["bug_level"] = strings.Join(data["bug_level"], "")
	a["bug_find_time"] = strings.Join(data["bug_find_time"], "")
	a["bug_entrustor"] = strings.Join(data["bug_entrustor"], "")
	a["bug_remarks"] = strings.Join(data["bug_remarks"], "")
	a["created_at"] = time.Now().Format("2006-01-02 15:04:05")

	res := m.Table("l_bug").InsertMap(a)

	if len(res.GetErrors()) != 0 {
		my_log.Error("插入bug表出现问题", res.GetErrors())
		return Return{
			Code: 500,
			Msg:  "服务器开小差",
		}
	}

	return Return{
		Code: 200,
		Msg:  "ok",
	}
}

/**
  插入Bug表
 */
func BuildUpdateData(data map[string][]string) Return {
	if _, ok := data["edition_info"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:edition_info",
		}
	}
	if _, ok := data["edition_desc"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:edition_desc",
		}
	}
	if _, ok := data["bug_uptime"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_uptime",
		}
	}
	if _, ok := data["bug_entrustor"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_entrustor",
		}
	}
	if _, ok := data["bug_remarks"]; !ok {
		return Return{
			Code: 500,
			Msg:  "缺少必传参数:bug_remarks",
		}
	}

	my_config.InitConfig("./config.conf")
	config := my_config.Run()

	m := my_sql.NewConnector(config)

	a := make(map[string]interface{})

	a["edition_info"] = strings.Join(data["edition_info"], "")
	a["edition_desc"] = strings.Join(data["edition_desc"], "")
	a["bug_uptime"] = strings.Join(data["bug_uptime"], "")
	a["bug_entrustor"] = strings.Join(data["bug_entrustor"], "")
	a["bug_remarks"] = strings.Join(data["bug_remarks"], "")
	a["created_at"] = time.Now().Format("2006-01-02 15:04:05")

	res := m.Table("l_bug").InsertMap(a)

	if len(res.GetErrors()) != 0 {
		my_log.Error("插入bug表出现问题", res.GetErrors())
		return Return{
			Code: 500,
			Msg:  "服务器开小差",
		}
	}

	return Return{
		Code: 200,
		Msg:  "ok",
	}
}

func GetBugList(data map[string][]string) Index {
	if _, ok := data["p"]; !ok {
		return Index{
			Code:  500,
			Msg:   "缺少必传参数:p",
			Total: 0,
		}
	}

	if _, ok := data["l"]; !ok {
		return Index{
			Code:  500,
			Msg:   "缺少必传参数:l",
			Total: 0,
		}
	}

	my_config.InitConfig("./config.conf")
	config := my_config.Run()
	m := my_sql.NewConnector(config)

	total := m.Table("l_bug").Count("id", "total").ToStringMapList()

	num, _ := strconv.Atoi(total[0]["total"])

	limit, _ := strconv.Atoi(strings.Join(data["l"], ""))

	offset, _ := strconv.Atoi(strings.Join(data["p"], ""))

	offset = (offset - 1) * limit

	res := m.Table("l_bug").OrderBy("id", "desc").Limit(limit).Offset(offset).Get().ToStringMapList()

	return Index{Code: 200, Msg: "ok", Total: num, Data: res}
}

func GetUpdateList(data map[string][]string) Index {
	if _, ok := data["p"]; !ok {
		return Index{
			Code:  500,
			Msg:   "缺少必传参数:p",
			Total: 0,
		}
	}

	if _, ok := data["l"]; !ok {
		return Index{
			Code:  500,
			Msg:   "缺少必传参数:l",
			Total: 0,
		}
	}

	my_config.InitConfig("./config.conf")
	config := my_config.Run()
	m := my_sql.NewConnector(config)

	total := m.Table("l_update").Count("id", "total").ToStringMapList()

	num, _ := strconv.Atoi(total[0]["total"])

	limit, _ := strconv.Atoi(strings.Join(data["l"], ""))

	offset, _ := strconv.Atoi(strings.Join(data["p"], ""))

	offset = (offset - 1) * limit

	res := m.Table("l_update").OrderBy("id", "desc").Limit(limit).Offset(offset).Get().ToStringMapList()

	return Index{Code: 200, Msg: "ok", Total: num, Data: res}
}