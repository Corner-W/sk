package main

import (
	"github.com/Corner-W/sk/log"
	"github.com/Corner-W/sk/module"
)

func main() {

	log.Debug("main start...")

	module.Modules_main()

	select {}

}
