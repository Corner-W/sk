package main

import (
	"encoding/json"
	"sk/log"
	"sk/module"
)

// TCP 客户端
func main() {

	tc := module.NewClient()

	msg := module.Msg{
		H: module.Mhead{
			Op:  1,
			Len: 14,
		},

		B: module.Mbody{
			Id:  2,
			Tid: 2,
			Bd:  "test meessage",
		},
	}

	ss, err := json.Marshal(msg)
	if err != nil {
		log.Error("json Marshal err!!")
	}

	tc.Conn.Write(ss)

}
