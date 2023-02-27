package module

import (
	"github.com/Corner-W/sk/log"
	"sync"
)

type Module interface {
	OnInit()
	OnDestroy()
	MsgProc(closeSig chan bool, message interface{})
}

type module struct {
	mi       Module
	closeSig chan bool
	//wg       sync.WaitGroup

	/*消息队列*/
	BufLen uint32
	Que    chan interface{}
	/*消息队列锁, channel本身线程安全，暂时无用*/
	qlocker sync.RWMutex

	Id   uint32
	Name string
	stat int
}

var modules sync.Map

func QueryMbyID(mid uint32) (*module, bool) {

	m, ok := modules.Load(mid)

	return m.(*module), ok
}

func ModuleReg(name string, mid uint32, mi Module, buflen uint32) bool {
	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool, 1)

	m.BufLen = buflen
	m.Que = make(chan interface{}, buflen)
	/*暂时没有支持并发安全*/
	m.Id = mid
	//
	m.Name = name
	m.stat = MODULE_STATE_INIT

	//mods = append(mods, m)
	v, ok := modules.Load(mid)
	if !ok {
		modules.Store(mid, m)
	} else {

		log.Fatal("Module %s  register Err! module id has been used by %s[%d] ", m.Name, v.(*module).Name, mid)

		return false
	}

	log.Debug("Module %s register successfully!!", m.Name)
	return true
	//modules.LoadOrStore()
}
func Init() {

	/*集中启动所有的业务模块*/
	modules.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		m := value.(*module)
		log.Debug("%s[%d] module starting...", m.Name, id)

		m.mi.OnInit()
		//	m.wg.Add(1)
		log.Debug("start module %s task queue!", m.Name)
		go run(m)

		return true
	})

}

func run(m *module) {

	m.stat = MODULE_STATE_RUNNING
	for {

		select {
		/*check channel */
		case msg, ok := <-m.Que:
			if ok == true {
				/*读到数据立即执行*/
				m.mi.MsgProc(m.closeSig, msg)
			} else {
				/*如果channel 被关闭，则会导致死循环，重写nil可以ignore该分支*/
				m.Que = nil
				m.stat = MODULE_STATE_CRASH
				log.Error("module running err")
				break
			}
		default:
			// log.Error("default: err!")
		}
	}

	//m.wg.Done()
}
