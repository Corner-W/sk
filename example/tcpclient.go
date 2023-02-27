package main

import (
	"encoding/json"
	"github.com/Corner-W/sk/log"
	"github.com/Corner-W/sk/module"
	"github.com/Corner-W/sk/module/agent"
	"time"
	"unsafe"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

// TCP 客户端
func main() {

	tc := agent.NewClient()

	msg := &module.Msg{
		H: module.Mhead{
			Op:  0,
			Len: 14,
		},

		B: module.Mbody{
			Id:  2,
			Tid: 2,
			Bd:  "test msg ,hello!",
		},
	}
	msg.H.Len = int(unsafe.Sizeof(*msg))

	for i := 0; i < 100; i++ {
		msg.H.Op = uint16(i)
		ss, err := json.Marshal(msg)
		if err != nil {
			log.Error("json Marshal err!!")
		}

		n, err := tc.Conn.Write(ss)
		if err != nil {

			log.Error("send msg Err!")
		}
		//time.Sleep(5 * time.Second)
		log.Debug("n = %d, length=%d", n, msg.H.Len)
	}
	time.Sleep(10 * time.Second)
	tc.Conn.Close()

	//json.Unmarshal()

}
func Encode(obj *module.Msg) ([]byte, error) {
	//buf := make([]byte, 512)

	//copy(buf, obj)

	b := &SliceMock{
		addr: uintptr(unsafe.Pointer(obj)),
		cap:  int(unsafe.Sizeof(*obj)),
		len:  int(unsafe.Sizeof(*obj)),
	}

	buf := *(*[]byte)(unsafe.Pointer(b))
	return buf, nil
}
