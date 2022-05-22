package desktop

import (
	"embed"
	"flag"
	"fmt"
	"github.com/maldan/gam-app-desktop/internal/app/desktop/api"
	"github.com/maldan/gam-app-desktop/internal/app/desktop/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-rapi"
	"github.com/maldan/go-rapi/rapi_core"
	"github.com/maldan/go-rapi/rapi_rest"
	"github.com/maldan/go-rapi/rapi_vfs"
	"log"
	"time"
)

func Start(frontFs embed.FS) {
	// Server
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16001, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")

	// Data
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()

	// Set
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

	// https://d1v-h.phncdn.com/hls/videos/202111/21/398393381/,1080P_4000K,720P_4000K,480P_2000K,240P_1000K,_398393381.mp4.urlset/index-f2-v1-a1.m3u8?ttl=1641250256&l=0&clientip=77.40.3.32&ipa=77.40.3.32&hash=ebb60e6baa41bc447514d07230b1df1f&

	// Load config
	cmhp_file.ReadJSON(core.DataDir+"/config.json", &core.AppConfig)

	// Backup schedule
	/*go func() {
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
	}()*/

	go func() {
		log.Println("START SCHEDULER")

		for {
			log.Println(cmhp_process.Exec(
				"rsync", "-ra", "--progress", "--delete", "--no-compress",
				"/root/.gam-data/", "/home/backup/gam-data",
			))
			time.Sleep(time.Hour * 6)
		}
	}()

	// Start server
	rapi.Start(rapi.Config{
		Host: fmt.Sprintf("%s:%d", *host, *port),
		Router: map[string]rapi_core.Handler{
			"/": rapi_vfs.VFSHandler{
				Root: "frontend/build",
				Fs:   frontFs,
			},
			"/api": rapi_rest.ApiHandler{
				Controller: map[string]interface{}{
					"main":        api.MainApi{},
					"process":     api.ProcessApi{},
					"application": api.ApplicationApi{},
				},
			},
		},
		DbPath: core.DataDir,
	})
}
