package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-desktop/internal/app/desktop"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
