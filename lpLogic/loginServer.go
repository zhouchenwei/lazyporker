/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import (
	"fmt"
	"log"
	"net/http"
)

type LoginServer struct {
	LoginLogic
}

var GLoginServer = &LoginServer{}

func (this *LoginServer) StartServer(port int) error {
	http.HandleFunc("/regist", this.OnRegist)
	http.HandleFunc("/login", this.OnLogin)
	this.Start()
	return http.ListenAndServe(fmt.Sprintf(`:%d`, port), nil)
}

func (this *LoginServer) OnRegist(rw http.ResponseWriter, req *http.Request) {
	log.Println(`regist`)
	// regist
	fmt.Sprintf(`ok`)
}

func (this *LoginServer) OnLogin(rw http.ResponseWriter, req *http.Request) {
	//login check
	uid := "" //数据库获取uid
	log.Println(`login check ok`)

	loginkey := this.GetLoginKey(uid)
	fmt.Sprintf(`loginkey=%s`, loginkey)
}
