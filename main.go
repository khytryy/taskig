package main

import (
	"taskig/cmd"
	"taskig/internal/conf"
)

func main() {
	conf.Init()
	cmd.Execute()
}
