package telnet

import (
	"fmt"
	"github.com/Corner-W/sk/module"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

func ModuleManagerHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	//LongWriteString(stdout, "module manager command \r\n")
	fmt.Fprintln(stdout, "modules display ")
	module.Modules.Range(func(key, value interface{}) bool {
		//id := key.(uint32)
		m := value.(*module.Mod)

		fmt.Fprintf(stdout, "%d   %s       %s        %d   %d \n",
			m.Id,
			m.Name,
			module.ModStateMap[m.Stat],
			cap(m.Que),
			len(m.Que))

		return true
	})
	//fmt.Println(args)

	return nil
}

func ModuleManagerProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {

	//fmt.Println(args)

	//fmt.Print(stdout)
	return telsh.PromoteHandlerFunc(ModuleManagerHandler, args...)
}

func ModuleMsgSendProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {

	return telsh.PromoteHandlerFunc(
		func(stdin io.ReadCloser,
			stdout io.WriteCloser,
			stderr io.WriteCloser,
			args ...string) error {

			module.MsgSend(module.MOUDLE_ID_GAME, "this is one test message!")
			fmt.Fprintf(stdout, "this is msgsend command %s \n", args)
			return nil
		},
		args...)
}
