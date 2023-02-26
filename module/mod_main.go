package module

import (
	"github.com/Corner-W/sk/module/game"
	"github.com/Corner-W/sk/module/user"
)

func Modules_main() {

	/*模块统一注册*/
	ModuleReg("game", MOUDLE_ID_GAME, game.New(), 1024)
	ModuleReg("user", MOUDLE_ID_USER, user.New(), 1024)
	ModuleReg("dispatch", MODULE_ID_DISPATCH, NewDispatch(), 1024)
	ModuleReg("agent", MODULE_ID_AGENT, NewAgent(), 2048)
	// ModuleReg("mqtt", MODULE_ID_MQTT, NewMqtt(), 2048)

	/*模块初始化，和任务启动*/
	Init()
}
