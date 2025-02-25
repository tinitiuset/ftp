package ftp

import (
	"fmt"
	"github.com/go-kit/log/level"
	"log"
)

// respond copies a string to the client and terminates it with the appropriate FTP line terminator
// for the datatype.
func (s *Session) respond(response string) {
	level.Info(*s.log).Log("respond", response)
	_, err := fmt.Fprint(s.conn, response, s.EOL())
	if err != nil {
		log.Print(err)
	}
}

// EOL returns the line terminator matching the FTP standard for the datatype.
func (s *Session) EOL() string {
	switch s.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}
