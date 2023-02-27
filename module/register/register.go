package register

import (
	"github.com/Corner-W/sk/module"
	"github.com/Corner-W/sk/module/agent"
	"github.com/Corner-W/sk/module/dispatch"
	"github.com/Corner-W/sk/module/game"
	"github.com/Corner-W/sk/module/user"
)

func Modules_main() {

	/*模块统一注册*/
	module.ModuleReg("game", module.MOUDLE_ID_GAME, game.New(), 1024)
	module.ModuleReg("user", module.MOUDLE_ID_USER, user.New(), 1024)
	module.ModuleReg("dispatch", module.MODULE_ID_DISPATCH, dispatch.NewDispatch(), 1024)
	module.ModuleReg("agent", module.MODULE_ID_AGENT, agent.NewAgent(), 2048)
	// ModuleReg("mqtt", MODULE_ID_MQTT, NewMqtt(), 2048)

	/*模块初始化，和任务启动*/
	module.Init()
}
