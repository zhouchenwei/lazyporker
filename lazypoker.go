/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/

package main

import (
	"./lpLogic"
	"log"
)

func main() {
	log.Println("Start...")

	lpLogic.GHall.Start()
	lpLogic.GPlayerManager.Start()
	lpLogic.GGameServer.Start(10086)
	lpLogic.GLoginServer.StartServer(10087)
}
