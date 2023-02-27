SK Server 
====
A pragmatic  server framework in Go (golang).

Features
---------

* Extremely easy to use
* Reliable
* Multicore support
* Modularity

Community
---------
SK 服务器框架简介
==================

SK 是一个由 Go 语言（golang）编写的开发效率和执行效率并重的开源游戏服务器框架。SK 适用于各类高性能服务器的开发，包括tcp,MQ,game服务器。

SK 的关注点：

* 良好的使用体验。SK 总是尽可能的提供简洁和易用的module接口，尽可能的提升开发的效率
* 稳定性。SK 总是尽可能的恢复运行过程中的错误，避免崩溃
* 多核支持。SK 通过模块机制和 尽可能的利用多核资源，同时又尽量避免各种副作用
* 模块机制
* 增加telnet后端服务器，命令行配置功能


源码概览
---------------

* sk/module 业务模块管理
* sk/db 数据库相关，目前支持 [MongoDB](https://www.mongodb.org/)
* sk/gate 网关模块，负责客户端的接入
* sk/netty 网络层
* sk/log 日志相关
* sk/telnet telnet shell服务器模块
* sk/timer 定时器


SK 的模块机制
---------------

一个服务器由多个模块组成，模块有以下特点：

* 每个模块运行在一个单独的 goroutine 中
* 模块间通过一套轻量的 channel 机制通讯

服务器在启动时进行模块的注册，例如：

```go
func Modules_main() {

    /*模块统一注册*/
    ModuleReg("game", MOUDLE_ID_GAME, game.New(), 1024)
    ModuleReg("user", MOUDLE_ID_USER, user.New(), 1024)
    ModuleReg("dispatch", MODULE_ID_DISPATCH, NewDispatch(), 1024)
    ModuleReg("agent", MODULE_ID_AGENT, NewAgent(), 2048)

    /*模块初始化，和任务启动*/
    Init()
}
```

这里按顺序注册了 四个个模块。每个模块都需要实现接口：

```go
type Module interface {
	OnInit()
	OnDestroy()
	MsgProc(closeSig chan bool, message interface{})
}
```

增加telnet后端调试命令
-----------
```go
    shellHandler := telsh.NewShellHandler()
    shellHandler.Register("mod", telsh.ProducerFunc(ModuleManagerProducer))

    addr := ":5001"
    log.Debug("telnet server is starting ...")
    if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
        panic(err)
    }
```

Timer 定时器
----------
```go
//循环定时器
func main() {
        tm := timer.NewTimer()

        tm.ScheduleFunc(1*time.Second, func() {
                log.Printf("schedule\n")
        })

        tm.Run()
}

//一次定时器

func main() {
tm := timer.NewTimer()

tm.AfterFunc(1*time.Second, func() {
log.Printf("after\n")
})

tm.AfterFunc(10*time.Second, func() {
log.Printf("after\n")
})
tm.Run()
}

```


使用 SK 开发服务器
---------------




Documentation
---------


Licensing
---------

SK is licensed under the Apache License, Version 2.0. 
