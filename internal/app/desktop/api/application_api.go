package api

import (
	"os"
	"strings"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
)

type ApplicationApi struct{}

func (u ApplicationApi) GetList() interface{} {
	out := cmhp_process.Exec("gam", "al")
	var list []core.Application

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
		var app core.Application
		core.To(m, &app)
		list = append(list, app)
	}

	return list
}

func (u ApplicationApi) GetIcon(args ArgsPath) *os.File {
	args.Context.ContentType = "image/svg+xml"

	if cmhp_file.Exists(args.Path + "/" + "icon.svg") {
		f, _ := os.Open(args.Path + "/" + "icon.svg")
		return f
	}
	f, _ := os.Open("app.svg")
	return f
}

// Save config
func (r ApplicationApi) PostConfig(args ArgsConfig) {
	cmhp_file.WriteJSON(core.DataDir+"/../"+args.AppId+"/config.json", &args.Config)
}

// Get config
func (r ApplicationApi) GetConfig(args ArgsAppId) map[string]interface{} {
	config := make(map[string]interface{})
	cmhp_file.ReadJSON(core.DataDir+"/../"+args.AppId+"/config.json", &config)
	return config
}
