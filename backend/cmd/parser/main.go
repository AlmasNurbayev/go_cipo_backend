package main

import "github.com/AlmasNurbayev/go_cipo_backend/internal/parser"

func main() {
	var Version = "v0.1.0"

	p := parser.New(Version)
	p.Init()
	p.Run()
}
