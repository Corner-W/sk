package shellcommand

import (
	"errors"
	"fmt"
	"github.com/Corner-W/sk/compt/shell"
	"github.com/Corner-W/sk/log"
	"strings"
)

type SkHandler struct {
	s *shell.Shell
}

func SkHandlerFactory(s *shell.Shell) shell.Handler {
	return &SkHandler{
		s: s,
	}
}

func (h *SkHandler) HandleLine(line string) error {
	log.Debug("LINE from %s: %s", h.s.InstanceName(), line)
	//h.s.AddString(" this is ssh command line...\n")
	flags := strings.Split(line, " ")

	cmd := flags[0]
	args := flags[1:]

	if cmd == "help" || cmd == "?" {

	} else if cmd == "exit" || cmd == "quit" {
		return errors.New("exit")
	} else {
		foundCommand := false
		for _, c := range Commands {
			if c.Name() == cmd || c.Name() == "fs."+cmd {
				if res := c.Run(c, args, h.s.RW()); res != true {
					fmt.Fprintf(h.s.RW(), "error: %v\n", res)
				}
				foundCommand = true
			}
		}
		if !foundCommand {
			fmt.Fprintf(h.s.RW(), "unknown command: %v\r\n", cmd)
		}
	}

	return nil
}

func (h *SkHandler) HandleEof() error {
	log.Debug("EOF from %s", h.s.InstanceName())
	return errors.New("ssh session is end...")
}

func Run() {

	sshServer := &shell.SshServer{
		Config: &shell.Config{
			Bind: ":2222",
			Users: map[string]shell.User{
				"user": {Password: "user"},
			},
		},
		HandlerFactory: SkHandlerFactory,
	}

	sshServer.ListenAndServe()
}
