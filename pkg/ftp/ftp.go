package ftp

import (
	"bufio"
	"github.com/go-kit/log/level"
	"strings"

	"github.com/tinitiuset/ftp/pkg/ftp/commands"
	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

// Serve scans incoming requests for valid commands and routes them to handler functions.
func Serve(s *Session) {
	s.respond(status.Ready)

	handlers := map[string]commands.CommandFunc{
		"CWD":  commands.Cwd(s),
		"PWD":  commands.Pwd(s),
		"LIST": commands.List(s),
		"STOR": commands.Stor(s),
		"EPSV": commands.Epsv(s),
		"USER": commands.User(s),
		"PASS": commands.Pass(s),
		"QUIT": func(args []string) {
			s.respond(status.Closing)
		},
		"TYPE": commands.Type(s),
	}

	scanner := bufio.NewScanner(s.conn)
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		level.Info(*s.log).Log("request", command, "args", strings.Join(args, ","))

		if handler, ok := handlers[command]; ok {
			handler(args)
		} else {
			s.respond(status.NotImplemented)
		}

		if command == "QUIT" {
			return
		}
	}
	if scanner.Err() != nil {
		level.Error(*s.log).Log("error", scanner.Err())
	}
}
