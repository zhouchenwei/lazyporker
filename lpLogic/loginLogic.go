/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import (
	"sync"
	"time"
)

const (
	define_loginkey_timeout_second = 60
)

type LoginInfo struct {
	LoginTime int64
	Uid       string
}

type LoginLogic struct {
	LoginInfoPool *sync.Pool
	MapLoginKey   map[string]*LoginInfo // loginkey
	MuxLoginKey   sync.Mutex
}

func (this *LoginLogic) Start() {
	this.LoginInfoPool = &sync.Pool{New: func() interface{} { return new(LoginInfo) }}
	go this.Routine_CheckLoginKey()
}

func (this *LoginLogic) Routine_CheckLoginKey() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			this.CleanLoginKeyTimeOut()
		}
	}
}

func (this *LoginLogic) CleanLoginKeyTimeOut() {
	this.MuxLoginKey.Lock()
	defer this.MuxLoginKey.Unlock()

	nowTime := time.Now().Unix()
	for k, v := range this.MapLoginKey {
		if nowTime-v.LoginTime > define_loginkey_timeout_second {
			delete(this.MapLoginKey, k)
			defer this.LoginInfoPool.Put(v)
		}
	}
}

func (this *LoginLogic) GetLoginKey(uid string) string {
	this.MuxLoginKey.Lock()
	defer this.MuxLoginKey.Unlock()

	loginkey := "" //产生loginkey

	linfo := this.LoginInfoPool.Get().(*LoginInfo)
	linfo.LoginTime = time.Now().Unix()
	linfo.Uid = uid
	this.MapLoginKey[loginkey] = linfo

	return loginkey
}

func (this *LoginLogic) IsLoginKeyVaild(loginKey string) string {
	this.MuxLoginKey.Lock()
	defer this.MuxLoginKey.Unlock()
	linfo, ok := this.MapLoginKey[loginKey]
	if ok {
		return linfo.Uid
	}
	return ""
}
