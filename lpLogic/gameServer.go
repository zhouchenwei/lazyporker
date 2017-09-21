/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import (
	"fmt"
	"log"
	"net"
)

type GameServer struct {
	listenfd net.Listener
}

var GGameServer = &GameServer{}

func (this *GameServer) Start(port int) error {
	var err error
	this.listenfd, err = net.Listen("tcp", fmt.Sprintf(`:%d`, port))
	if err != nil {
		return err
	}
	go this.Routine_Listen()
	return nil
}

func (this *GameServer) Routine_Listen() {
	defer this.listenfd.Close()

	for {
		if conn, err := this.listenfd.Accept(); err != nil {
			log.Println(err)
		} else {
			log.Println(conn.RemoteAddr().String())
			conn.(*net.TCPConn).SetLinger(0)
			go this.ProcessConn(conn)
		}
	}
}

func (this *GameServer) ProcessConn(conn net.Conn) {
	//get loginkey
	loginKey := ""

	//无效
	uid := GLoginServer.IsLoginKeyVaild(loginKey)
	if len(uid) == 0 {
		return
	}
	//获取玩家信息
	player := GPlayerManager.GetPlayerInfo(uid)
	if player == nil {
		return
	}
	player.Init(conn)
	GHall.AddHall(player)
}
