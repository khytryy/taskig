package main

import (
	"github.com/khytryy/taskig/cmd"
	"github.com/khytryy/taskig/internal/conf"
)

func main() {
	conf.Init()
	cmd.Execute()
}
