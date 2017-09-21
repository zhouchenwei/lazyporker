/*=================================
User name: lazypos
Time: 2017-09-21
Explain:
=================================*/
package lpLogic

import (
	"fmt"
	"net"
)

//基本发送函数
func sendMesssage(conn net.Conn, msg []byte) error {
	sendlen := 0
	for sendlen < len(msg) {
		n, err := conn.Write(msg[sendlen:])
		if err != nil {
			return err
		} else if n <= 0 {
			return fmt.Errorf(`xxxx`)
		}
		sendlen += n
	}
	return nil
}

//基本接收函数
func recvMessage(conn net.Conn, totallen int) ([]byte, error) {
	recvlen := 0
	buf := make([]byte, totallen)
	for recvlen < totallen {
		n, err := conn.Read(buf[recvlen:])
		if err != nil {
			return []byte{}, err
		} else if n <= 0 {
			return []byte{}, fmt.Errorf(`xxxx`)
		}
		recvlen += n
	}
	return buf, nil
}

//接受一条指令
func RecvCommond(conn net.Conn) ([]byte, error) {

	return []byte{}, nil
}

//发送一条指令
func SendCommond(conn net.Conn, content []byte) error {

	return nil
}
