package main

import "github.com/Corner-W/sk/cmd"

func main() {
<<<<<<< HEAD

	log.Debug("main start...")
	go telnet.Run()
	register.Modules_main()

	//select {}

=======
	cmd.Execute()
>>>>>>> dev
}
