package commands

import (
	"fmt"

	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

func Pass(s Session) CommandFunc {
	return func(args []string) {
		if s.GetUsername() == "" {
			s.Respond(status.BadSequence)
			return
		}

		if len(args) < 1 {
			s.Respond(status.BadArguments)
			return
		}

		// For now, accept any password
		s.SetAuthenticated(true)
		s.Respond(fmt.Sprintf(status.LoggedIn, s.GetUsername()))
	}
}
