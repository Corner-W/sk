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

var wg sync.WaitGroup

type Mod struct {
	mi Module

	//队列退出信号 if true; close this channel
	closeSig chan bool
	//wg       sync.WaitGroup

	/*消息队列*/
	BufLen uint32
	Que    chan interface{}
	/*消息队列锁, channel本身线程安全，暂时无用*/
	qlocker sync.RWMutex

	Id   uint32
	Name string
	Stat int
}

var Modules sync.Map

func QueryMbyID(mid uint32) (*Mod, bool) {

	m, ok := Modules.Load(mid)

	return m.(*Mod), ok
}

func Register(name string, mid uint32, mi Module, buflen uint32) bool {
	m := new(Mod)
	m.mi = mi
	m.closeSig = make(chan bool, 1)

	m.BufLen = buflen
	m.Que = make(chan interface{}, buflen)
	/*暂时没有支持并发安全*/
	m.Id = mid
	//
	m.Name = name
	m.Stat = MODULE_STATE_INIT

	//mods = append(mods, m)
	v, ok := Modules.Load(mid)
	if !ok {
		Modules.Store(mid, m)
	} else {

		log.Fatal("Module %s  register Err! module id has been used by %s[%d] ", m.Name, v.(*Mod).Name, mid)

		return false
	}

	log.Debug("Module %s register successfully!!", m.Name)
	return true
	//modules.LoadOrStore()
}
func Startup() {

	/*集中启动所有的业务模块*/
	Modules.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		m := value.(*Mod)
		log.Debug("%s[%d] module starting...", m.Name, id)

		m.mi.OnInit()
		//	m.wg.Add(1)
		log.Debug("start module %s task queue!", m.Name)

		wg.Add(1)
		//go run(m)
		go func(m *Mod) {

			defer wg.Done()
			m.Stat = MODULE_STATE_RUNNING
			for {

				select {
				/*check channel 由于没有default分支，如果channel是空的，那么for循环会阻塞到该分支*/
				case msg, ok := <-m.Que:
					if ok == true {
						/*读到数据立即执行*/
						m.mi.MsgProc(m.closeSig, msg)
					} else {
						/*如果channel 被关闭，则会导致死循环，重写nil可以ignore该分支*/
						//m.Que = nil
						m.Stat = MODULE_STATE_CRASH
						log.Error("module running err")
						break
					}
				case sig := <-m.closeSig:
					if sig == true {
						m.Stat = MODULE_STATE_CLOSE
						//close(m.Que)

						/*停止读写该channel*/
						break
					}
					//default:
					// 此处需要释放cpu，否则cpu占用会非常高
					//log.Error("msg channel is empty %s", m.Name)
				}
			}

		}(m)

		return true
	})
	wg.Wait()
}

/*
对channel操作的集中场景分析：
--------------------------------------------------------------
｜操作	      		｜ nil的channel	 正常channel	  已关闭的channel
--------------------------------------------------------------
｜读 <-ch	   		｜ 阻塞	         成功或阻塞	  读到零值
｜写 ch<-	  		｜ 阻塞	         成功或阻塞	  panic
｜关闭 close(ch)		｜ panic	     成功	      panic
*/
func run(m *Mod) {

	defer wg.Done()

	m.Stat = MODULE_STATE_RUNNING
	for {

		select {
		/*check channel 由于没有default分支，如果channel是空的，那么for循环会阻塞到该分支*/
		case msg, ok := <-m.Que:
			if ok == true {
				/*读到数据立即执行*/
				m.mi.MsgProc(m.closeSig, msg)
			} else {
				/*如果channel 被关闭，则会导致死循环，重写nil可以ignore该分支*/
				//m.Que = nil
				m.Stat = MODULE_STATE_CRASH
				log.Error("module running err")
				break
			}
			//default:
			// 此处需要释放cpu，否则cpu占用会非常高
			//log.Error("msg channel is empty %s", m.Name)
		}
	}

	//m.wg.Done()
}
