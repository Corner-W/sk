package shellcommand

import (
	"fmt"
	"io"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print sk version",
	Long:      `Version prints the sk version`,
}

func runVersion(cmd *Command, args []string, rw io.ReadWriter) bool {
	if len(args) != 0 {
		cmd.Usage()
	}

	//fmt.Printf("version %s %s %s\n", util.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(rw, "sk version: %s\r\n", "v1.0")
	return true
}
