package telnet

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

func ModuleManagerHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	LongWriteString(stdout, "module manager command \r\n")

	//fmt.Println(args)

	return nil
}

func ModuleManagerProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {

	//fmt.Println(args)

	//fmt.Print(stdout)
	return telsh.PromoteHandlerFunc(ModuleManagerHandler, args...)
}
