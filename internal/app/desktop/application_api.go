package desktop

import (
	"strings"

	"github.com/maldan/go-cmhp/cmhp_process"
)

type ApplicationApi int

func (u ApplicationApi) GetList() interface{} {
	out := cmhp_process.Exec("gam", "al")
	var list []Application

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
		var app Application
		To(m, &app)
		list = append(list, app)
	}

	return list
}
