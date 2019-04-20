/*
@Time : 2019/4/20 11:30 
@Author : Tian Kejie.
*/
package models

type Index struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Data []map[string]string `json:"data"`
}

type Lists []Index
