package api

import "github.com/maldan/go-restserver"

type ArgsPath struct {
	Context *restserver.RestServerContext
	Path    string `json:"path"`
}

type ArgsAppId struct {
	AppId string `json:"appId"`
}

type ArgsConfig struct {
	AppId  string                 `json:"appId"`
	Config map[string]interface{} `json:"config"`
}

type ArgRun struct {
	Url  string `json:"url"`
	Host string `json:"host"`
}
