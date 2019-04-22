/*
@Time : 2019/4/20 11:05
@Author : Tian Kejie.
*/
package my_route

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		BugList,
	},
	Route{
		"Index",
		"GET",
		"/_buglist",
		BugList,
	},
	Route{
		"Hello",
		"GET",
		"/_uplist",
		UpList,
	},
	Route{
		"CreateBug",
		"POST",
		"/bug/create",
		Create,
	},
	Route{
		"CreateUpdate",
		"POST",
		"/update/create",
		Update,
	},
}
