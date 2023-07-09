package main

import (
	"fmt"
	"net"
)

func main() {
	s := NewServer("127.0.0.1", 8099)
	s.Run()
}

type server struct {
	ip   string
	port int
}

func NewServer(ip string, port int) *server {
	return &server{
		ip:   ip,
		port: port,
	}
}

func (s *server) Run() {
	addr := fmt.Sprintf("%s:%d", s.ip, s.port)
	Listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := Listener.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer func() {
				fmt.Println("conn close")
				// conn.Close()
			}()

			// for {
			b := make([]byte, 10)
			conn.Read(b)
			fmt.Println(string(b))
			// b, _ := io.ReadAll(conn)
			conn.Write(b)
			// }
		}()
	}
}
