package modules

import (
	"net"
	"time"
)

type Connection struct {
	port   string
	isOpen bool
}

func TcpConn(input Input) []Connection {

	var result []Connection

	for _, port := range input.ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(input.host, port), timeout)
		if err != nil {
			result = append(result, Connection{port: port, isOpen: false})
		}

		if conn != nil {
			defer conn.Close()
			result = append(result, Connection{port: port, isOpen: true})
		}
	}

	return result
}
