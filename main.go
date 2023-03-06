package main

import (
	"github.com/Corner-W/sk/log"
	"github.com/Corner-W/sk/module/register"
	"github.com/Corner-W/sk/telnet"
)

func main() {

	log.Debug("main start...")
	go telnet.Run()
	register.Modules_main()

	//select {}

}
