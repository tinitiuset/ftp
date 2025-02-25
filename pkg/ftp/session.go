package ftp

import (
	"fmt"
	"github.com/go-kit/log"

	"net"
	"path/filepath"
)

// Session represents a connection to the FTP server
type Session struct {
	log           *log.Logger
	conn          net.Conn
	dataType      dataType
	dataPort      *dataPort
	rootDir       string
	workDir       string
	username      string
	password      string
	authenticated bool
	dataListener  net.Listener
}

// NewSession returns a new FTP connection
func NewSession(log *log.Logger, conn net.Conn, rootDir string) *Session {
	absRoot, err := filepath.Abs(rootDir)
	if err != nil {
		absRoot = rootDir
	}

	return &Session{
		log:      log,
		conn:     conn,
		rootDir:  absRoot,
		workDir:  absRoot,
		dataType: ascii,
	}
}

func (s *Session) Respond(code string)            { s.respond(code) }
func (s *Session) GetWorkDir() string             { return s.workDir }
func (s *Session) GetRootDir() string             { return s.rootDir }
func (s *Session) SetWorkDir(path string)         { s.workDir = path }
func (s *Session) IsAuthenticated() bool          { return s.authenticated }
func (s *Session) GetDataListener() net.Listener  { return s.dataListener }
func (s *Session) SetDataListener(l net.Listener) { s.dataListener = l }
func (s *Session) GetUsername() string            { return s.username }
func (s *Session) SetUsername(username string)    { s.username = username }
func (s *Session) SetAuthenticated(auth bool)     { s.authenticated = auth }
func (s *Session) SetDataType(t string) {
	switch t {
	case "A":
		s.dataType = ascii
	case "I":
		s.dataType = binary
	}
}
func (s *Session) SetDataPort(host string, port int) {
	s.dataPort = &dataPort{
		host: host,
		port: port,
	}
}
func (s *Session) AcceptDataSession() (net.Conn, error) {
	if s.dataListener == nil {
		return nil, fmt.Errorf("no data listener")
	}
	conn, err := s.dataListener.Accept()
	if err != nil {
		return nil, err
	}
	defer s.dataListener.Close()
	s.dataListener = nil
	return conn, nil
}
