package networking

import (
	"bytes"
	"log"
	"net"
	"net/url"

	"github.com/gongt/remote-shell/internal/constants"
)

type Server struct {
	sock     net.Listener
	callback func(target net.Conn, message []byte)
}

func CreateServer(cb func(target net.Conn, message []byte)) (*Server, error) {
	log.Printf("server = %s\n", constants.Server)
	u, err := url.ParseRequestURI("http://" + constants.Server)
	log.Printf("listening: %s\n", u.Port())
	sock, err := net.Listen("tcp", ":"+u.Port())
	if err != nil {
		log.Println("failed: ", err)
		return nil, err
	}
	log.Println("listen ok.")

	return &Server{
		sock:     sock,
		callback: cb,
	}, nil
}

func (s *Server) Start() {
	for {
		s.Handle()
	}
}

func (s *Server) Handle() {
	conn, err := s.sock.Accept()
	if err != nil {
		log.Printf("Can not accept socket: %s\n", err)
		return
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		log.Printf("Error reading: %s\n", err)
		return
	}

	b := bytes.Trim(buf, "\x00")
	log.Printf("receive: %s\n", string(b))
	s.callback(conn, b)
}
