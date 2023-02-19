package game

import (
	"fmt"
	"sk/log"
)

type GameTask struct {
	msg string
}

func New() *GameTask {

	return &GameTask{
		msg: "game module",
	}
}
func (g *GameTask) OnInit() {

	log.Debug("module game init...")

}

func (g *GameTask) OnDestroy() {

}

func (g *GameTask) MsgProc(closeSig chan bool, message interface{}) {

	log.Debug("module game Enter...")

	msg := message.(Msg)

	fmt.Println(msg)

}
