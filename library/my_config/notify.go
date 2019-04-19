/*
@Time : 2019/4/18 17:58 
@Author : Tian Kejie.
*/
package my_config

type Notifyer interface {
	Callback(*Config)
}
