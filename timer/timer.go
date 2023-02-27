package timer

import "time"

type Next interface {
	Next(time.Time) time.Time
}

// 定时器接口
type Timer interface {
	// 一次性定时器
	AfterFunc(expire time.Duration, callback func()) TimeNoder

	// 周期性定时器
	ScheduleFunc(expire time.Duration, callback func()) TimeNoder

	// 自定义下次的时间
	CustomFunc(n Next, callback func()) TimeNoder

	// 运行
	Run()

	// 停止所有定时器
	Stop()
}

// 停止单个定时器
type TimeNoder interface {
	Stop()
}

// 定时器构造函数
func NewTimer(opt ...Option) Timer {
	var o option
	for _, cb := range opt {
		cb(&o)
	}

	if o.timeWheel {
		return newTimeWheel()
	}

	if o.minHeap {
		return newMinHeap()
	}

	return newTimeWheel()
}

// 定时器接口 version 2
type Timerv2 interface {
	// 一次性定时器
	//TimerStart(expire time.Duration, m interface{}, callback func(t TimeNoder, v interface{})) TimeNoder
	// 周期性定时器
	TimerLoopInit(expire time.Duration, m interface{}, callback func(t TimeNoder, v interface{})) TimeNoder

	// 运行
	TimerStart()

	// 停止所有定时器
	//Stop()
}
