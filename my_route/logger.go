/*
@Time : 2019/4/20 11:18 
@Author : Tian Kejie.
*/
package my_route

import (
	"Golang-RESTful/library/my_log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		inner.ServeHTTP(w, r)

		my_log.Info("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(startTime))
	})
}


