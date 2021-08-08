package desktop

import (
	"encoding/json"
	"strconv"
)

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
	Name    string `json:"name"`
	Author  string `json:"author"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

func To(m map[string]string, v interface{}) {
	out, _ := json.Marshal(m)
	json.Unmarshal(out, v)
}

func Atoi(x string) int {
	xx, _ := strconv.Atoi(x)
	return xx
}
