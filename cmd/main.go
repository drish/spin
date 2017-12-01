package main

import (
	"github.com/drish/spin"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug = kingpin.Flag("debug", "Enable debug mode.").Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	spin := spin.New()
	spin.Spin()
}
