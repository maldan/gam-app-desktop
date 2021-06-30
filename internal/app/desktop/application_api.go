package desktop

import (
	"encoding/json"
	"os/exec"
	"strings"
)

type ApplicationApi int

func (u ApplicationApi) GetList() interface{} {
	c, b := exec.Command("gam", "app", "list", "--format=json"), new(strings.Builder)
	c.Stdout = b
	c.Run()
	var list []Application
	json.Unmarshal([]byte(b.String()), &list)
	return list
}
