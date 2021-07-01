package desktop

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type ProcessApi int

type PA_PostRunArgs struct {
	Url  string `json:"url"`
	Host string `json:"host"`
}

var WindowInfo map[int]Window = make(map[int]Window)

func (u ProcessApi) GetList() interface{} {
	c, b := exec.Command("gam", "process", "list", "--format=json"), new(strings.Builder)
	c.Stdout = b
	c.Run()
	var list []Process
	json.Unmarshal([]byte(b.String()), &list)
	for i := 0; i < len(list); i++ {
		list[i].Window = WindowInfo[list[i].Pid]
		list[i].Window.Pid = list[i].Pid

		if list[i].Window.X < 0 {
			list[i].Window.X = 0
		}
		if list[i].Window.Y < 0 {
			list[i].Window.Y = 0
		}
		if list[i].Window.Width < 100 {
			list[i].Window.Width = 100
		}
		if list[i].Window.Height < 100 {
			list[i].Window.Height = 100
		}
	}
	return list
}

func (u ProcessApi) PostKill(args Process) {
	c, b := exec.Command("gam", "process", "kill", fmt.Sprintf("%v", args.Pid)), new(strings.Builder)
	c.Stdout = b
	c.Run()
	fmt.Println(b.String())
}

func (u ProcessApi) PostRun(args PA_PostRunArgs) {
	c, b := exec.Command(
		"gam",
		"run",
		fmt.Sprintf("%v", args.Url),
		fmt.Sprintf("--host=%v", args.Host)), new(strings.Builder)
	c.Stdout = b
	c.Run()

	huilo := strings.Split(b.String(), ", ")
	pid, _ := strconv.Atoi(strings.Split(huilo[0], ":")[1])
	WindowInfo[pid] = Window{Pid: pid, X: 100, Y: 100, Width: 720, Height: 480}
}

func (u ProcessApi) PostSetWindow(args Window) {
	WindowInfo[args.Pid] = args
}
