package api

type ArgsPath struct {
	// Context *restserver.RestServerContext
	Path string `json:"path"`
}

type ArgsId struct {
	Id string `json:"id"`
}

type ArgsAppId struct {
	AppId string `json:"appId"`
}

type ArgsConfig struct {
	AppId  string                 `json:"appId"`
	Config map[string]interface{} `json:"config"`
}

type ArgsRun struct {
	Url       string `json:"url"`
	Host      string `json:"host"`
	DesktopId int    `json:"desktopId"`
}

type ArgsAppData struct {
	AppId string
	Data  string
}
