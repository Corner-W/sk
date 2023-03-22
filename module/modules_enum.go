package module

const (
	MOUDLE_ID_NULL = iota
	MOUDLE_ID_GAME
	MOUDLE_ID_USER
	MODULE_ID_DISPATCH
	MODULE_ID_AGENT
	MODULE_ID_MQTT
)

const (
	MODULE_STATE_INIT = iota
	MODULE_STATE_INIT_ERR
	MODULE_STATE_RUNNING
	MODULE_STATE_CRASH
	MODULE_STATE_CLOSE

	MODULE_STATE_END
)

/*module 状态映射值，只读不写*/
var ModStateMap = []string{
	"init",
	"init_error",
	"running",
	"crash",
	"closed",
	"end",
}
