package api

import "github.com/maldan/go-restserver"

type ArgsIcon struct {
	Context *restserver.RestServerContext
	Path    string `json:"path"`
}

type ArgRun struct {
	Url  string `json:"url"`
	Host string `json:"host"`
}
