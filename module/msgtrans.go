package module

import "sk/log"

func MsgSend(dest uint32, msg interface{}) bool {

	m, ok := QueryMbyID(dest)
	if ok != true {

		log.Error("msgsend err! dest = %d", dest)

		return ok
	}

	//m.qlocker.Lock()
	m.Que <- msg
	//defer m.qlocker.Unlock()

	return true
}
