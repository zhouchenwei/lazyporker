/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import (
	"net"
	"time"
)

const (
	task_player_cmd_sendmsg = 0x3001 //发送信息
	task_player_cmd_kick    = 0x3002 //踢掉
)

type PlayerBase struct {
	conn net.Conn //连接
	uid  string   //唯一ID
}

type PlayerEx struct {
	PlayerBase
	BaseClass

	bOnLine     bool    //是否在线
	nLastOnline int64   //最后在线时间
	desk        *DeskEx //哪个桌子
	deskId      int     //桌号
	siteId      int     //坐位号
	bLook       bool    //是否旁观
}

// 创建
func (this *PlayerEx) Create(uid string) {
	this.uid = uid
	this.BaseClass.Start()
	go this.Routine_Recv()
}

// 初始化/重新初始化
func (this *PlayerEx) Init(conn net.Conn) {
	this.nLastOnline = time.Now().Unix()
	this.conn = conn
	this.bOnLine = true
	this.desk = nil
	this.deskId = -1
	this.siteId = -1
	this.bLook = false
}

func (this *PlayerEx) ProcessTask(c *CommMessage) {
	switch c.Cmd {
	case task_player_cmd_sendmsg:
		// do same thing
	}
}

func (this *PlayerEx) Routine_Recv() {
	for {
		content, err := RecvCommond(this.conn)
		if err != nil {
			break
		}
		// 解析请求
		cmd, info, err := this.paerseQuery(content)
		if err != nil {
			break
		}
		// 处理请求
		if err = GProcess.ProcessCmd(cmd, info, this); err != nil {
			break
		}
	}
	// 处理错误
	GProcess.ProcessCmd(cmd_player_net_error, "", this)
}

func (this *PlayerEx) paerseQuery(text []byte) (int, interface{}, error) {
	return 0, nil, nil
}
