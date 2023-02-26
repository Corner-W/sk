package module

/*消息分发模块，负责接收client端的消息*/

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/Corner-W/sk/log"
	"github.com/Corner-W/sk/netty"
	"io"
	"net"
	"unsafe"
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
	log.Debug("msgRecv Enter...")
	reader := bufio.NewReader(conn)
	// 阻塞读数据
	dec := json.NewDecoder(reader)
	for {
		var m Msg
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Error("decode msg Err!")
			break

		}

		fmt.Println(m)
	}
	/*client 端socket关闭时，退出循环*/
	log.Debug("Receive Message end!!!")
	/*
		packs := make([]byte, 1024)
		for {
			n, _ := reader.Read(packs)

			if n == 0 {

				log.Debug("read msg end!!")
				break
			}
			spkg := string(packs)
			log.Debug("receive msg = %s", spkg)

			dec := json.NewDecoder(strings.NewReader(spkg))
			for {
				var m Msg
				if err := dec.Decode(&m); err == io.EOF {
					break
				} else if err != nil {
					log.Error("err is not nil")
					break

				}
				fmt.Printf("%v", m.H.Op)
			}

		}

	*/

}

func tcpServerInit() {

	ser := netty.New("127.0.0.1:9900")

	ser.MsgHandler = MsgRecv

	go ser.ListenAndServe()

}

func DecodeMsg(packs []byte, pkgLen int) bool {

	pkg := packs

	for {
		headlen := int(unsafe.Sizeof(Mhead{}))

		h := pkg[0 : headlen-1]

		lengthbyte := h[2:]

		lengthbuf := bytes.NewBuffer(lengthbyte)
		var msgLen int
		binary.Read(lengthbuf, binary.LittleEndian, &msgLen)
		log.Debug("msgLen=%d, pkgLen=%d, headlen=%d", msgLen, pkgLen, headlen)
		SingMsgProc(pkg[:msgLen-1])
		//
		if pkgLen > msgLen {

			pkg = pkg[msgLen:]
			continue
		} else if pkgLen == msgLen {
			log.Debug("Decode msg end!!!!")
			break
		} else {

			log.Error("Decode msg err!!, msgLen = %d, pkgLen = %d", msgLen, pkgLen)
			return false
		}

	}

	return true
}

func SingMsgProc(pkg []byte) {

	msg := Msg{}
	buf := bytes.NewBuffer(pkg)

	binary.Read(buf, binary.LittleEndian, &msg)

	//log.Debug("msg:len=%d, op=%d ", msg.H.Len, msg.H.Op)

}
