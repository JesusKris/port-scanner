package modules

import (
	"net"
	"time"
)

func UdpConn(input Input) []Connection {

	var result []Connection

	time := time.Now().Add(1 * time.Second)
	for _, port := range input.ports {
		remoteAddr, err := net.ResolveUDPAddr("udp", input.host+":"+port)
		if err != nil {
			result = append(result, Connection{port: port, isOpen: false})
			continue
		}

		conn, err := net.DialUDP("udp", nil, remoteAddr)
		if err != nil {
			result = append(result, Connection{port: port, isOpen: false})
			continue
		}

		conn.SetReadDeadline(time)

		//close the connection
		defer conn.Close()

		_, err = conn.Write([]byte("peeka boo"))
		if err != nil {
			result = append(result, Connection{port: port, isOpen: false})
			continue
		}

		// buffer to get data
		received := make([]byte, 1024)
		_, err = conn.Read(received)
		if err != nil {
			result = append(result, Connection{port: port, isOpen: false})
			continue
		}

		result = append(result, Connection{port: port, isOpen: true})
	}

	return result
}
