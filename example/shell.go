package main

import (
	"errors"
	shell "github.com/Corner-W/sk/compt/shell"
	"github.com/Corner-W/sk/log"
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

	if line == "exit" {
		h.s.AddString(" session exit...\n")
		return errors.New("session Exit!")
	}
	//h.s.rw
	return nil
}

func (h *SkHandler) HandleEof() error {
	log.Debug("EOF from %s", h.s.InstanceName())
	return errors.New("Ssh session is end...")
}

func main() {
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
