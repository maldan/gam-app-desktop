package desktop

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/maldan/go-restserver"
	"github.com/zserge/lorca"
)

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "host")
	var port = flag.Int("port", 16001, "port")
	var gui = flag.Bool("gui", false, "gui")
	flag.Parse()

	if *gui {
		go (func() {
			ui, _ := lorca.New("", "", 480, 320)
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
