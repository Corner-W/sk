package agent

/*消息分发模块，负责接收client端的消息*/

import (
	"github.com/Corner-W/sk/log"
)

type AgentTask struct {
	msg string
}

func NewAgent() *AgentTask {

	return &AgentTask{
		msg: "user module",
	}
}

var (
	tcpclient *TcpClient
)

func (a *AgentTask) OnInit() {

	log.Debug("module agent init...")

	//tcpclient = NewClient()

}

func (a *AgentTask) OnDestroy() {

}

func (a *AgentTask) MsgProc(closeSig chan bool, message interface{}) {

	log.Debug("module agent  Enter...")

}
