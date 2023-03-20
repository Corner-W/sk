package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	liner := "sk create -p option"

	flags := strings.Split(liner, " ")

	fmt.Println(flags[0])

	flag.Parse()

	cmdLine := flag.NewFlagSet(flags[0], flag.ExitOnError)
	op := cmdLine.String("create", "", "create one demo")

	cmdLine.Parse(flags[1:])

	//cmdLine.Usage()

	fmt.Println(*op)

}
