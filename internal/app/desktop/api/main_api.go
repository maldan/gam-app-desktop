package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-rapi/rapi_core"
)

type MainApi struct{}

// Get list of process
func (u MainApi) GetBackground(ctx *rapi_core.Context, args ArgsId) string {
	ctx.IsServeFile = true
	return core.DataDir + "/background/" + args.Id + ".jpeg"
}

func (u MainApi) PostApplicationData(args ArgsAppData) {
	// Prepare app id
	tuple := strings.Split(args.AppId, "/")
	if len(tuple) != 2 {
		rapi_core.Fatal(rapi_core.Error{
			Field:       "appId",
			Description: "Invalid application id",
		})
	}
	appId := tuple[0] + "-gam-app-" + tuple[1]

	// Store data
	fileName := fmt.Sprintf("%v.json", time.Now().UnixNano())
	cmhp_file.Write(core.DataDir+"/external_data/"+appId+"/"+fileName, args.Data)

	// Try to send data
	go SendDataToProcess(appId)
}
