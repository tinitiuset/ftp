package commands

import (
	"net"
)

// CommandFunc represents a command handler function
type CommandFunc func(args []string)

// Session interface defines methods needed by commands
type Session interface {
	Respond(code string)
	GetWorkDir() string
	GetRootDir() string
	SetWorkDir(path string)
	IsAuthenticated() bool
	GetDataListener() net.Listener
	SetDataListener(net.Listener)
	AcceptDataSession() (net.Conn, error)
	GetUsername() string
	SetUsername(username string)
	SetAuthenticated(bool)
	SetDataType(string)
	SetDataPort(host string, port int)
}
