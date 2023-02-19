package main

import (
	"sk/game"
	//"sk/console"
	"sk/log"
	"sk/module"
)

func main() {

	log.Debug("main start...")
	//console.Init()
	module.Modules_main()

	message := game.Msg{
		Msgid: 12,
		Name:  "to game module test",
	}
	for i := 0; i < 10; i++ {
		module.MsgSend(module.MOUDLE_ID_GAME, message)

	}

	//tc := module.NewClient()
	//tc.Conn.Write([]byte("hello, this is client"))

	select {}

}
