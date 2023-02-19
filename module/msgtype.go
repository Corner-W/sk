package module

type GameMsg struct {
	Id  uint32
	Tid uint32
}

type Mhead struct {
	Op  uint16 `json:"operator"`
	Len int    `json:"len"`
}

type Mbody struct {
	Id  uint32 `json:"id"`
	Tid uint32 `json:"tid"`
	Bd  string `json:"body"`
}

type Msg struct {
	H Mhead

	B Mbody
}
