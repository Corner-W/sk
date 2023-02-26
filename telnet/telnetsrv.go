package telnet

import (
	"github.com/Corner-W/sk/log"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

func fiveHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	LongWriteString(stdout, "The number FIVE looks like this: 5\r\n")

	//fmt.Println(args)

	return nil
}

func fiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {

	//fmt.Println(args)

	//fmt.Print(stdout)
	return telsh.PromoteHandlerFunc(fiveHandler, args...)
}

func danceHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {

	LongWriteString(stdout, "\r \r\n")

	return nil
}

func danceProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {

	return telsh.PromoteHandlerFunc(danceHandler)
}

//var shellHandler *telsh.ShellHandler

func init() {

}
func Run() {

	shellHandler := telsh.NewShellHandler()

	shellHandler.WelcomeMessage = `
 __          __ ______  _        _____   ____   __  __  ______ 
 \ \        / /|  ____|| |      / ____| / __ \ |  \/  ||  ____|
  \ \  /\  / / | |__   | |     | |     | |  | || \  / || |__   
   \ \/  \/ /  |  __|  | |     | |     | |  | || |\/| ||  __|  
    \  /\  /   | |____ | |____ | |____ | |__| || |  | || |____ 
     \/  \/    |______||______| \_____| \____/ |_|  |_||______|

`

	// Register the "five" command.
	commandName := "five"
	commandProducer := telsh.ProducerFunc(fiveProducer)

	shellHandler.Register(commandName, commandProducer)

	// Register the "dance" command.
	commandName = "dance"
	commandProducer = telsh.ProducerFunc(danceProducer)

	shellHandler.Register(commandName, commandProducer)

	shellHandler.Register("dance", telsh.ProducerFunc(danceProducer))

	shellHandler.Register("mod", telsh.ProducerFunc(ModuleManagerProducer))
	//fmt.Println("run success!...")
	addr := ":5001"
	log.Debug("telnet server is starting ...")
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
		panic(err)
	}

}
