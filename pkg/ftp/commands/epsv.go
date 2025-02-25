package commands

import (
	"fmt"
	"net"

	"github.com/tinitiuset/ftp/pkg/ftp/status"
)

func Epsv(s Session) CommandFunc {
	return func(args []string) {
		if !s.IsAuthenticated() {
			s.Respond(status.NotLoggedIn)
			return
		}

		// Create listener on random port
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			s.Respond(status.DataSessionNotOpen)
			return
		}

		// Get the port number
		_, port, err := net.SplitHostPort(listener.Addr().String())
		if err != nil {
			listener.Close()
			s.Respond(status.DataSessionNotOpen)
			return
		}

		s.SetDataListener(listener)
		s.Respond(fmt.Sprintf(status.ExtendedPassive, port))
	}
}
