package main

import "github.com/AlmasNurbayev/go_cipo_backend/internal/parserML"

func main() {
	var Version = "v0.1.0"

	p := parserML.New(Version)
	p.Init()
	p.Run()
}
