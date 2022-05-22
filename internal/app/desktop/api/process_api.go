package api

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_net"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-rapi/rapi_core"
)

type ProcessApi struct{}

// var WindowInfo map[string]core.Window = make(map[string]core.Window)

var WindowInfo sync.Map = sync.Map{}

// Start
func init() {
	go (func() {
		time.Sleep(time.Second)
		m := make(map[string]core.Window)
		cmhp_file.ReadJSON(core.DataDir+"/window.json", &m)
		for k, v := range m {
			WindowInfo.Store(k, v)
		}

		for {
			// Save each second windows info
			m = make(map[string]core.Window)
			WindowInfo.Range(func(key, value interface{}) bool {
				m[key.(string)] = value.(core.Window)
				return true
			})
			cmhp_file.Write(core.DataDir+"/window.json", &m)
			time.Sleep(time.Second)
		}
	})()
}

// Get list of process
func (u ProcessApi) GetList() interface{} {
	out, err := cmhp_process.Exec("gam", "pl")
	rapi_core.FatalIfError(err)

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
		x, ok := WindowInfo.Load(fmt.Sprintf("%v", list[i].Pid))
		if ok {
			list[i].Window = x.(core.Window)
		} else {
			list[i].Window = core.Window{}
		}
		// list[i].Window = WindowInfo[fmt.Sprintf("%v", list[i].Pid)]
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
func (u ProcessApi) PostRun(args ArgsRun) {
	out, err := cmhp_process.Exec("gam",
		"run",
		fmt.Sprintf("%v", args.Url),
		fmt.Sprintf("--host=%v", args.Host))
	rapi_core.FatalIfError(err)

	l := strings.Split(out, "\n")
	m := make(map[string]string)
	for _, kv := range l {
		if kv == "" {
			continue
		}
		kvv := strings.Split(kv, ": ")
		m[kvv[0]] = kvv[1]
	}

	// WindowInfo[(m["pid"])] = core.Window{
	WindowInfo.Store((m["pid"]), core.Window{
		Pid:       core.Atoi(m["pid"]),
		X:         100 + float64(rand.Intn(100)),
		Y:         100 + float64(rand.Intn(100)),
		Width:     900,
		Height:    600,
		DesktopId: args.DesktopId,
		Dock:      "",
	})

	go SendDataToProcess(m["appId"])
}

// Set window position
func (u ProcessApi) PostSetWindow(args core.Window) {
	WindowInfo.Store(fmt.Sprintf("%v", args.Pid), args)
}

func SendDataToProcess(appId string) {
	// Get process list
	u := ProcessApi{}
	l := u.GetList().([]core.Process)

	// Search process
	for _, v := range l {
		if v.Args["appId"] == appId {
			// Get files
			files, err := cmhp_file.List(core.DataDir + "/external_data/" + appId)
			if err != nil {
				return
			}

			// Send each file
			for _, file := range files {
				cmhp_net.Request(cmhp_net.HttpArgs{
					Method: "POST",
					Url:    fmt.Sprintf("http://%v:%v/system/signal/applicationData", v.Args["host"], v.Args["port"]),
					InputJSON: &map[string]string{
						"path": core.DataDir + "/external_data/" + appId + "/" + file.Name(),
					},
				})
			}

			return
		}
	}
}
