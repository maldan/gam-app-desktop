package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
)

type ProcessApi struct{}

var WindowInfo map[string]core.Window = make(map[string]core.Window)

// Start
func init() {
	go (func() {
		time.Sleep(time.Second)
		cmhp_file.ReadJSON(core.DataDir+"/window.json", &WindowInfo)

		for {
			// Save each second windows info
			cmhp_file.WriteJSON(core.DataDir+"/window.json", &WindowInfo)
			time.Sleep(time.Second)
		}
	})()
}

// Get list of process
func (u ProcessApi) GetList() interface{} {
	out := cmhp_process.Exec("gam", "pl")
	list := make([]core.Process, 0)

	// json.Unmarshal([]byte(out), &list)
	lines := strings.Split(out, "\n\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		l := strings.Split(line, "\n")
		m := make(map[string]string)
		for _, kv := range l {
			kvv := strings.Split(kv, ": ")
			m[kvv[0]] = kvv[1]
		}
		var proc core.Process
		core.To(m, &proc)
		proc.Pid = core.Atoi(m["pid"])
		proc.Args = m
		delete(proc.Args, "pid")
		delete(proc.Args, "cmd")
		list = append(list, proc)
	}

	for i := 0; i < len(list); i++ {
		list[i].Window = WindowInfo[fmt.Sprintf("%v", list[i].Pid)]
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

// Kill process
func (u ProcessApi) PostKill(args core.Process) {
	fmt.Println(cmhp_process.Exec("gam", "kill", fmt.Sprintf("%v", args.Pid)))
}

// Run process
func (u ProcessApi) PostRun(args ArgRun) {
	out := cmhp_process.Exec("gam",
		"run",
		fmt.Sprintf("%v", args.Url),
		fmt.Sprintf("--host=%v", args.Host))

	l := strings.Split(out, "\n")
	m := make(map[string]string)
	for _, kv := range l {
		if kv == "" {
			continue
		}
		kvv := strings.Split(kv, ": ")
		m[kvv[0]] = kvv[1]
	}

	WindowInfo[(m["pid"])] = core.Window{
		Pid:       core.Atoi(m["pid"]),
		X:         100,
		Y:         100,
		Width:     720,
		Height:    480,
		DesktopId: args.DesktopId,
	}
}

// Set window position
func (u ProcessApi) PostSetWindow(args core.Window) {
	WindowInfo[fmt.Sprintf("%v", args.Pid)] = args
}
