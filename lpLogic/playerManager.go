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

type PlayerManager struct {
	MapPlayers map[string]*PlayerEx //uid->pinfp
	muxPlayers sync.Mutex
}

var GPlayerManager = &PlayerManager{}

func (this *PlayerManager) Start() {
	this.MapPlayers = make(map[string]*PlayerEx)
	go this.Routine_CleanPlayer()
}

// 太久的删掉
func (this *PlayerManager) Routine_CleanPlayer() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			this.muxPlayers.Lock()
			nowTime := time.Now().Unix()
			for k, p := range this.MapPlayers {
				if nowTime-p.nLastOnline > 1800 {
					p.CreateTask(task_cmd_exit, "", nil, nil)
					delete(this.MapPlayers, k)
				}
			}
			this.muxPlayers.Unlock()
		}
	}
}

//获取玩家信息
func (this *PlayerManager) GetPlayerInfo(uid string) *PlayerEx {
	this.muxPlayers.Lock()
	defer this.muxPlayers.Unlock()

	pInfo, ok := this.MapPlayers[uid]
	// 已经登陆
	if ok && pInfo.bOnLine {
		return nil
	}
	// 新登陆
	if !ok {
		pInfo := &PlayerEx{}
		pInfo.Create(uid)
		// 加载用户信息
		if err := this.LoadPlayerInfo(pInfo); err != nil {
			return nil
		}
		this.MapPlayers[uid] = pInfo
	}
	return pInfo
}

//加载信息
func (this *PlayerManager) LoadPlayerInfo(p *PlayerEx) error {
	return nil
}
