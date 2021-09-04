package api

import (
	"os"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
)

type MainApi struct{}

// Get list of process
func (u MainApi) GetBackground(args ArgsId) *os.File {
	file, _ := os.Open(core.DataDir + "/background/" + args.Id + ".jpeg")
	return file
}
