package user

import (
	"sk/log"
)

type UserTask struct {
	msg string
}

func New() *UserTask {

	return &UserTask{
		msg: "user module",
	}
}
func (g *UserTask) OnInit() {

	log.Debug("module user init...")

}

func (g *UserTask) OnDestroy() {

}

func (g *UserTask) MsgProc(closeSig chan bool, message interface{}) {

	log.Debug("module user is Enter...")

}
