package main

import (
	"sk/log"
	"sk/module"
)

func main() {

	log.Debug("main start...")

	module.Modules_main()

	select {}

}
