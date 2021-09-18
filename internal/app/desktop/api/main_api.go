package api

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-restserver"
)

type MainApi struct{}

// Get list of process
func (u MainApi) GetBackground(args ArgsId) *os.File {
	file, _ := os.Open(core.DataDir + "/background/" + args.Id + ".jpeg")
	return file
}

func (u MainApi) PostApplicationData(args ArgsAppData) {
	// Prepare app id
	tuple := strings.Split(args.AppId, "/")
	if len(tuple) != 2 {
		restserver.Fatal(500, "unknown", "appId", "Invalid application id")
	}
	appId := tuple[0] + "-gam-app-" + tuple[1]

	// Store data
	fileName := fmt.Sprintf("%v.json", time.Now().UnixNano())
	cmhp_file.WriteText(core.DataDir+"/external_data/"+appId+"/"+fileName, args.Data)

	// Try to send data
	go SendDataToProcess(appId)
}
