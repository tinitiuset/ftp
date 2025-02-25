package commands

import (
	"io"
	"os"
	"path/filepath"

	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

func Stor(s Session) CommandFunc {
	return func(args []string) {
		if !s.IsAuthenticated() {
			s.Respond(status.NotLoggedIn)
			return
		}

		if len(args) < 1 {
			s.Respond(status.BadArguments)
			return
		}

		// Only allow files in workDir
		filename := filepath.Base(args[0])
		path := filepath.Join(s.GetWorkDir(), filename)

		// Ensure directory exists
		if err := os.MkdirAll(s.GetWorkDir(), 0755); err != nil {
			s.Respond(status.FileUnavailable)
			return
		}

		if s.GetDataListener() == nil {
			s.Respond(status.DataSessionNotOpen)
			return
		}

		s.Respond(status.FileOK)

		dataSession, err := s.AcceptDataSession()
		if err != nil {
			s.Respond(status.DataSessionNotOpen)
			return
		}
		defer dataSession.Close()

		file, err := os.Create(path)
		if err != nil {
			s.Respond(status.FileUnavailable)
			return
		}
		defer file.Close()

		if _, err = io.Copy(file, dataSession); err != nil {
			s.Respond(status.FileUnavailable)
			return
		}

		s.Respond(status.ClosingData)
	}
}
