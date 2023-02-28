package module

import "github.com/Corner-W/sk/log"

func MsgSend(dest uint32, msg interface{}) bool {

	m, ok := QueryMbyID(dest)
	if ok != true {

		log.Error("msgsend err! dest = %d", dest)

		return ok
	}

	/*一旦channel写满，将会发生阻塞*/

	// m.Que <- msg

	select {
	/* 1. 如果channel发生写满的场景，将会走default分支
	2 . channel本身支持并发安全，不需要添加锁*/
	case m.Que <- msg:

	default:
		log.Error("module %s msg channel is full", m.Name)
		return false
	}

	return true
}
