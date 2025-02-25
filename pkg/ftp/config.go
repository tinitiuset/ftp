package ftp

import (
	"flag"
)

type Config struct {
	Interface string
	RootDir   string
}

func NewConfig() *Config {

	var config Config

	flag.StringVar(&config.Interface, "interface", "0.0.0.0:2121", "Interface and port to listen on")
	flag.StringVar(&config.RootDir, "rootDir", "./public", "Root directory for the server")

	return &config
}
