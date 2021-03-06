package desktop

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/maldan/go-restserver"
	"github.com/zserge/lorca"
)

var DataDir = "."

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16001, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")
	var gui = flag.Bool("gui", false, "Use Gui")
	var initDev = flag.Bool("initDev", false, "Install dev")
	var width = flag.Int("width", 1100, "Window Width")
	var height = flag.Int("height", 900, "Window Height")
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()
	DataDir = *dataDir

	// Copy as dev app
	if *initDev {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		GamAppDir := strings.ReplaceAll(dirname, "\\", "/") + "/.gam-app"
		source, _ := os.Open(os.Args[0])
		os.MkdirAll(GamAppDir+"/dev-desktop-v0.0.0", 0755)
		destination, err := os.Create(GamAppDir + "/dev-desktop-v0.0.0/app.exe")
		if err != nil {
			panic(err)
		}
		io.Copy(destination, source)
	}

	if *gui {
		go (func() {
			ui, _ := lorca.New("", "", *width, *height)
			defer ui.Close()
			ui.Load(fmt.Sprintf("http://%s:%d/", *host, *port))
			<-ui.Done()
			os.Exit(0)
		})()
	}

	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"process":     new(ProcessApi),
			"application": new(ApplicationApi),
		},
	})
}
