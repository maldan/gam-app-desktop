package desktop

import (
	"embed"
	"flag"
	"fmt"
	"time"

	"github.com/maldan/gam-app-desktop/internal/app/desktop/api"
	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-restserver"
)

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16001, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")
	var _ = flag.Bool("gui", false, "Use Gui")
	// var initDev = flag.Bool("initDev", false, "Install dev")
	var _ = flag.Int("width", 1100, "Window Width")
	var _ = flag.Int("height", 900, "Window Height")
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()
	core.DataDir = *dataDir

	// Copy as dev app
	/* if *initDev {
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
	} */

	// Load config
	cmhp_file.ReadJSON(core.DataDir+"/config.json", &core.AppConfig)

	// Backup schedule
	go func() {
		fmt.Println("START SCHEDULLER")

		for {
			// Backup each file
			for app, _ := range core.AppConfig.BackupQueue {
				fmt.Println("gam", "backup", app)
				fmt.Println(cmhp_process.Exec("gam", "backup", app))
			}
			fmt.Println("Backuped")
			time.Sleep(time.Hour * 6)
		}
	}()

	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"main":        api.MainApi{},
			"process":     api.ProcessApi{},
			"application": api.ApplicationApi{},
		},
	})
}
