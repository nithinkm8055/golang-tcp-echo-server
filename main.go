package main

import (
	"flag"

	"github.com/nithinkm8055/golang-tcp-echo-server/config"
	"github.com/nithinkm8055/golang-tcp-echo-server/pkg/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "server-host", config.Host, "host address, ex: 127.0.0.1")
	flag.IntVar(&config.Port, "server-port", config.Port, "host port")
	flag.Parse()
}

func main() {
	setupFlags()
	panic(server.ListenAndServe())
}
