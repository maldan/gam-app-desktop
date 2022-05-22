module github.com/maldan/gam-app-desktop

go 1.18

replace github.com/maldan/go-rapi => ../../go_lib/rapi

replace github.com/maldan/go-cmhp => ../../go_lib/cmhp

require (
	github.com/maldan/go-cmhp v0.0.23
	github.com/maldan/go-rapi v0.0.6
)
