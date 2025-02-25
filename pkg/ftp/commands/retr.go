package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

func Retr(s Session) CommandFunc {
	return func(args []string) {
		if !s.IsAuthenticated() {
			s.Respond(status.NotLoggedIn)
			return
		}

		if len(args) < 1 {
			s.Respond(status.BadArguments)
			return
		}

		filename := filepath.Base(args[0])
		path := filepath.Join(s.GetWorkDir(), filename)

		info, err := os.Stat(path)
		if err != nil || !info.Mode().IsRegular() {
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

		file, err := os.Open(path)
		if err != nil {
			s.Respond(status.FileUnavailable)
			return
		}
		defer file.Close()

		// Print to stdout and send to client
		reader := io.TeeReader(file, os.Stdout)
		fmt.Printf("--- Contents of %s ---\n", filename)
		_, err = io.Copy(dataSession, reader)
		fmt.Println("\n--- End of file ---")

		if err != nil {
			s.Respond(status.FileUnavailable)
			return
		}

		s.Respond(status.ClosingData)
	}
}
