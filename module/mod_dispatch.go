package module

/*消息分发模块，负责接收client端的消息*/

import (
	"bufio"
	"net"
	"sk/log"
	"sk/netty"
)

type DisTask struct {
	msg string
}

func NewDispatch() *DisTask {

	return &DisTask{
		msg: "user module",
	}
}
func (d *DisTask) OnInit() {

	log.Debug("module dispatch init...")

	tcpServerInit()

}

func (d *DisTask) OnDestroy() {

}

func (d *DisTask) MsgProc(closeSig chan bool, message interface{}) {

	log.Debug("module dispatch  Enter...")

}

func MsgRecv(conn net.Conn) {

	//var buffer []byte

	reader := bufio.NewReader(conn)

	packs := make([]byte, 512)

	n, err := reader.Read(packs)

	if err != nil {

		log.Error("receiv msg err!")
	}

	log.Debug("msg: %s, length=%d", string(packs), n)

}

func tcpServerInit() {

	ser := netty.New("127.0.0.1:9900")

	ser.MsgHandler = MsgRecv

	go ser.ListenAndServe()

}
