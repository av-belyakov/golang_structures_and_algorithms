package unixsocket_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Path_Unix_Socket string = "/tmp/unix_socket_test"
	Server_Response  string = "Message received successfully\n"
)

func handlerConn(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)

		return
	}

	fmt.Println("Server - received:", string(buffer[:n]))

	response := []byte(Server_Response)
	_, err = conn.Write(response)
	if err != nil {
		log.Printf("Error writing response to connection: %v", err)

		return
	}
}

func UnixSocketServer(ctx context.Context, pathSocket string) error {
	if err := os.Remove(pathSocket); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	listener, err := net.Listen("unix", pathSocket)
	if err != nil {
		return err
	}

	go func() {
		defer listener.Close()

		for {
			select {
			case <-ctx.Done():
				return

			default:
				conn, err := listener.Accept()
				if err != nil {
					log.Println(err)

					continue
				}

				handlerConn(conn)
			}
		}
	}()

	return err
}

func UnixSocketClient(pathSocket string) (string, error) {
	conn, err := net.Dial("unix", pathSocket)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("My some message!")); err != nil {
		return "", err
	}

	buffer := make([]byte, 1024)
	num, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	msg := string(buffer[:num])

	return msg, nil
}

func TestUnixSocket(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	err := UnixSocketServer(ctx, Path_Unix_Socket)
	assert.NoError(t, err)

	response, err := UnixSocketClient(Path_Unix_Socket)
	assert.NoError(t, err)
	assert.Equal(t, response, Server_Response)

	cancel()
}
