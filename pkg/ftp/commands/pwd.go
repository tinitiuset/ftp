package commands

import (
	"fmt"
	"strings"

	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

func Pwd(s Session) CommandFunc {
	return func(args []string) {
		if !s.IsAuthenticated() {
			s.Respond(status.NotLoggedIn)
			return
		}

		// Convert absolute path to relative path from root
		relativePath := strings.TrimPrefix(s.GetWorkDir(), s.GetRootDir())
		if relativePath == "" {
			relativePath = "/"
		}

		s.Respond(fmt.Sprintf(status.PathCreated, relativePath))
	}
}
