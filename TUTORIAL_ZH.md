SK 服务器框架简介
==================

SK 是一个由 Go 语言（golang）编写的开发效率和执行效率并重的开源游戏服务器框架。SK 适用于各类高性能服务器的开发，包括tcp,MQ,game服务器。

SK 的关注点：

* 良好的使用体验。SK 总是尽可能的提供简洁和易用的module接口，尽可能的提升开发的效率
* 稳定性。SK 总是尽可能的恢复运行过程中的错误，避免崩溃
* 多核支持。SK 通过模块机制和 尽可能的利用多核资源，同时又尽量避免各种副作用
* 模块机制。

SK 的模块机制
---------------

一个服务器由多个模块组成，模块有以下特点：

* 每个模块运行在一个单独的 goroutine 中
* 模块间通过一套轻量的 channel 机制通讯

Leaf 不建议在游戏服务器中设计过多的模块。

游戏服务器在启动时进行模块的注册，例如：

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



源码概览
---------------

* sk/module 业务模块管理
* leaf/db 数据库相关，目前支持 [MongoDB](https://www.mongodb.org/)
* sk/gate 网关模块，负责客户端的接入
* sk/netty 网络层
* leaf/log 日志相关
* leaf/recordfile 用于管理游戏数据
* leaf/timer 定时器相关
* leaf/util 辅助库

使用 SK 开发服务器
---------------

