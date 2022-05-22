package api

import (
	"strings"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-rapi/rapi_core"
)

type ApplicationApi struct{}

func (u ApplicationApi) GetList() interface{} {
	out, err := cmhp_process.Exec("gam", "al")
	rapi_core.FatalIfError(err)

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

func (u ApplicationApi) GetIcon(ctx *rapi_core.Context, args ArgsPath) string {
	ctx.IsServeFile = true

	if cmhp_file.Exists(args.Path + "/" + "icon.svg") {
		return args.Path + "/" + "icon.svg"
	}

	return "app.svg"
}

// Save config
func (r ApplicationApi) PostConfig(args ArgsConfig) {
	cmhp_file.Write(core.DataDir+"/../"+args.AppId+"/config.json", &args.Config)
}

// Get config
func (r ApplicationApi) GetConfig(args ArgsAppId) map[string]interface{} {
	config := make(map[string]interface{})
	cmhp_file.ReadJSON(core.DataDir+"/../"+args.AppId+"/config.json", &config)
	return config
}
