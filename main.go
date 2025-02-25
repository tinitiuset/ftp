package main

import (
	"os"

	"net"
	"path/filepath"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	ftp "github.com/tinitiuset/ftp/pkg/ftp"
)

func setUpLogger() *log.Logger {
	var logger log.Logger

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger,
		"ts", log.DefaultTimestampUTC,
	)
	logger = level.NewFilter(logger, level.AllowDebug())

	return &logger
}

func createRootDir(logger log.Logger) {
	absPath, err := filepath.Abs(config.RootDir)
	if err != nil {
		level.Error(logger).Log("err", err, "msg", "Failed to get absolute path")
		os.Exit(1)
	}
	level.Debug(logger).Log("path", absPath, "msg", "Server root directory")
	if err := os.MkdirAll(absPath, 0755); err != nil {
		level.Error(logger).Log("err", err, "path", absPath, "msg", "Failed to create root directory")
		os.Exit(1)
	}
}

var config *ftp.Config

func main() {
	logger := *setUpLogger()

	config = ftp.NewConfig()

	createRootDir(log.With(logger, "component", "storage"))

	listener, err := net.Listen("tcp", config.Interface)
	if err != nil {
		level.Error(logger).Log("err", err, "address", config.Interface, "msg", "Failed to start server")
		os.Exit(1)
	}
	level.Info(logger).Log("interface", config.Interface, "msg", "Server started")

	sessionIndex := 0

	for {
		conn, err := listener.Accept()
		if err != nil {
			level.Error(logger).Log("err", err, "msg", "Failed to accept connection")
			continue
		}
		level.Info(logger).Log("remote", conn.RemoteAddr().String(), "msg", "New connection")

		go handleSession(log.With(logger, "session", sessionIndex), conn)
		sessionIndex++
	}
}

func handleSession(logger log.Logger, c net.Conn) {

	defer c.Close()
	absPath, err := filepath.Abs(config.RootDir)
	if err != nil {
		level.Error(logger).Log("err", err, "msg", "Failed to get absolute path")
		os.Exit(1)
	}

	ftp.Serve(ftp.NewSession(&logger, c, absPath))
}
