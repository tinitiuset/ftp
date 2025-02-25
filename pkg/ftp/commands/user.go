package commands

import "github.com/tinitiuset/ftp/pkg/ftp/status"

func User(s Session) CommandFunc {
	return func(args []string) {
		if len(args) < 1 {
			s.Respond(status.BadArguments)
			return
		}

		// For now, accept any username
		s.SetUsername(args[0])
		s.Respond(status.UserOK)
	}
}
