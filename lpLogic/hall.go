/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import ()

const (
	task_hall_cmd_add          = 0x11001 //加入大厅
	task_hall_cmd_player_leave = 0x11002 //玩家离开
)

type Hall struct {
	BaseClass
	MapDesks   map[int]*DeskEx      //桌子列表
	MapPlayers map[string]*PlayerEx //玩家列表 key:uid
}

var GHall = &Hall{}

func (this *Hall) Start() {
	this.MapDesks = make(map[int]*DeskEx)
	this.BaseClass.Start()
}

func (this *Hall) AddHall(p *PlayerEx) {
	this.CreateTask(task_hall_cmd_add, "", p, nil)
}

// 任务处理函数
func (this *Hall) ProcessTask(c *CommMessage) {
	switch c.Cmd {
	case task_hall_cmd_add:
		p := c.ExInfo.(*PlayerEx)
		this.MapPlayers[p.uid] = p
		// do same thing
	}
}
