package fake_server

import (
	"fmt"
	"net"
	"time"
)

//**
// Usage example:
// srv := fake_server.New("localhost", 12323)
// srv.Start()
// defer srv.Stop()
//


type Server struct {
	host     string
	port     int
	listener net.Listener
}

func New(host string, port int) Server {
	return Server{
		host: host,
		port: port,
	}
}

func (s Server) Start() (err error) {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer s.Stop()
	fmt.Println("Listening on " + addr)
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go s.handleRequest(conn)
	}
}

func (s Server) StartWithDelay(delaySeconds int) {
	go  func () {
		time.Sleep(time.Duration(delaySeconds) * time.Second)
		s.Start()
	}()
}

func (s Server) Stop() {
	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}
}

func (s Server) handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Write([]byte("Message received."))
	conn.Close()
}
