package server

type Server interface {
	Start() error
	Shutdown() error
}
