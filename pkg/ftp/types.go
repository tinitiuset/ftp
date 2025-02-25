package ftp

type dataType int

const (
	ascii dataType = iota
	binary
)

type dataPort struct {
	host string
	port int
}
