/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import (
	"sync"
)

const (
	max_players = 4 //打牌人数

	task_desk_cmd_leave = 0x2001
)

// 桌子基本信息
type DeskBase struct {
	arrPlayers []*PlayerEx         //玩家列表
	MapLooker  map[int][]*PlayerEx //旁观列表

	deskId int //桌号
}

// 游戏相关
type DeskEx struct {
	DeskBase
	BaseClass

	bStarting bool //是否开始
	muxDesk   sync.Mutex
}

func (this *DeskEx) Create(deskId int) {
	this.deskId = deskId
	this.arrPlayers = make([]*PlayerEx, max_players)
	this.MapLooker = make(map[int][]*PlayerEx)
	this.BaseClass.Start()
}

func (this *DeskEx) ProcessTask(c *CommMessage) {
	switch c.Cmd {
	case task_desk_cmd_leave:
		// do same things
	}
}

// 加入桌子
func (this *DeskEx) AddDesk(p *PlayerEx) int {
	this.muxDesk.Lock()
	defer this.muxDesk.Unlock()

	for i, v := range this.arrPlayers {
		if v == nil {
			this.arrPlayers[i] = p
			return i
		}
	}
	return -1
}

func (this *DeskEx) LeaveDesk(p *PlayerEx) {
	this.CreateTask(task_desk_cmd_leave, "", p, nil)
}
