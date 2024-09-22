package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/nithinkm8055/golang-tcp-echo-server/config"
)

func readConn(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respondConn(c net.Conn, in string) error {
	_, err := c.Write([]byte(in))
	return err
}

func ListenAndServe() error {

	log.Println("listening for connections on ", config.Host, config.Port)
	lsnr, err := net.Listen("tcp4", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		return fmt.Errorf("failed to listen on %s:%d : %w", config.Host, config.Port, err)
	}

	var conn_clients int = 0

	for {
		conn, err := lsnr.Accept()
		if err != nil {
			return fmt.Errorf("cannot accept connections %w", err)
		}

		conn_clients += 1
		log.Println("client connected with address:", conn.RemoteAddr(), "concurrent clients: ", conn_clients)

		for {
			read, err := readConn(conn)
			if err != nil {
				conn.Close()
				conn_clients -= 1
				log.Println("client disconnected: ", conn.RemoteAddr(), "concurrent clients: ", conn_clients)

				if err == io.EOF {
					break
				}
				log.Println("err read", err)
			}

			log.Print("command: ", read)

			if err := respondConn(conn, read); err != nil {
				log.Println("cannot write to connection ", err)
			}
		}
	}
}
