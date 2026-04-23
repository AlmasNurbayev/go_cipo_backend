package main

import "github.com/AlmasNurbayev/go_cipo_backend/internal/parserJSON"

func main() {
	var Version = "v0.1.0"

	p := parserJSON.New(Version)
	p.Init()
	p.Run()
}
