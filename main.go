package main

import (
	"github.com/Corner-W/sk/log"
	"github.com/Corner-W/sk/module"
	"github.com/Corner-W/sk/telnet"
)

func main() {

	log.Debug("main start...")

	module.Modules_main()

	go telnet.Run()

	select {}

}
