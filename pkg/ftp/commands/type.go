package commands

import "github.com/tinitiuset/ftp/pkg/ftp/status"

func Type(s Session) CommandFunc {
	return func(args []string) {
		if !s.IsAuthenticated() {
			s.Respond(status.NotLoggedIn)
			return
		}

		if len(args) == 0 {
			s.Respond(status.BadArguments)
			return
		}

		switch args[0] {
		case "A", "I":
			s.SetDataType(args[0])
			s.Respond(status.OK)
		default:
			s.Respond(status.BadArguments)
		}
	}
}
