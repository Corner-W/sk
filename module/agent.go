package module

import (
	"net"
	"sk/log"
)

type TcpClient struct {
	Stat uint32 //链接状态

	Conn *net.TCPConn

	Raddr string
}

func NewClient() *TcpClient {

	var stat uint32 = 0

	addr := "127.0.0.1:9900"

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {

		log.Error("ResolveTcpAddr err!!")
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err == nil {

		stat = 1
		log.Debug("agent Connect remote %s successfully!", addr)

	} else {

		log.Error("agent Connect remote %s failed!", addr)
	}

	//defer conn.Close()
	return &TcpClient{
		Stat:  stat,
		Conn:  conn,
		Raddr: addr,
	}

}