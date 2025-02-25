package commands

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

func Cwd(s Session) CommandFunc {
	return func(args []string) {
		if !s.IsAuthenticated() {
			s.Respond(status.NotLoggedIn)
			return
		}

		if len(args) < 1 {
			s.Respond(status.BadArguments)
			return
		}

		newPath := filepath.Join(s.GetWorkDir(), args[0])
		if !filepath.IsAbs(newPath) {
			newPath = filepath.Join(s.GetRootDir(), newPath)
		}

		if !strings.HasPrefix(newPath, s.GetRootDir()) {
			s.Respond(status.FileUnavailable)
			return
		}

		if stat, err := os.Stat(newPath); err != nil || !stat.IsDir() {
			s.Respond(status.FileUnavailable)
			return
		}

		s.SetWorkDir(newPath)
		s.Respond(status.OK)
	}
}
