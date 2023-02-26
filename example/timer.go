package main

import (
	"github.com/Corner-W/sk/log"
	"github.com/antlabs/timer"
	"time"
)

var jumpNum = 0

func main() {
	tm := timer.NewTimer()

	tm.ScheduleFunc(1*time.Second, func() {

		log.Debug("schedule, jumpNum: %v\n", jumpNum)

		jumpNum++
	})

	tm.Run()
}
