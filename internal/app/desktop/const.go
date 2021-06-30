package desktop

type Window struct {
	Pid         int     `json:"pid"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	IsMinimized bool    `json:"isMinimized"`
	IsMaximized bool    `json:"isMaximized"`
}

type Process struct {
	Pid    int               `json:"pid"`
	Name   string            `json:"name"`
	Cmd    string            `json:"cmd"`
	Args   map[string]string `json:"args"`
	Window Window            `json:"window"`
}

type Application struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
