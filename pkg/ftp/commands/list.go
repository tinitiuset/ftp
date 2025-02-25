package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

func List(s Session) CommandFunc {
	return func(args []string) {
		if !s.IsAuthenticated() {
			s.Respond(status.NotLoggedIn)
			return
		}

		// Default to current directory if no path specified
		path := s.GetWorkDir()
		if len(args) > 0 {
			path = filepath.Join(s.GetWorkDir(), args[0])
		}

		// Ensure path is within root directory
		if !strings.HasPrefix(path, s.GetRootDir()) {
			s.Respond(status.FileUnavailable)
			return
		}

		// Check if directory exists and is readable
		if _, err := os.ReadDir(path); err != nil {
			s.Respond(status.FileUnavailable)
			return
		}

		// Check if we have an active data listener
		if s.GetDataListener() == nil {
			s.Respond(status.DataSessionNotOpen)
			return
		}

		s.Respond(status.FileOK)

		// Accept connection from client
		dataSession, err := s.AcceptDataSession()
		if err != nil {
			s.Respond(status.DataSessionNotOpen)
			return
		}
		defer dataSession.Close()

		// Read directory and send listing
		entries, err := os.ReadDir(path)
		if err != nil {
			s.Respond(status.FileUnavailable)
			return
		}

		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			// Only list regular files
			if !info.Mode().IsRegular() {
				continue
			}
			fmt.Fprintf(dataSession, "%s %d %s\r\n",
				info.Mode().String(),
				info.Size(),
				info.Name())
		}

		s.Respond(status.ClosingData)
	}
}
