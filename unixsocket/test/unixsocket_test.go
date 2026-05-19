package unixsocket_test

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

const Path_Unix_Socket string = "/tmp/unix_socket_test"

func TestUnixSocket(t *testing.T) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := UnixSocketServer(ctx, Path_Unix_Socket); err != nil {
		log.Fatal(err)
	}

	conn, err := UnixSocketClient(ctx, Path_Unix_Socket)
	if err != nil {
		log.Fatal(err)
	}

	go sendClientMessage(stop, conn)

	<-ctx.Done()

	log.Println("завершение работы программы")
}

// UnixSocketServer unix сервер
func UnixSocketServer(ctx context.Context, pathSocket string) error {
	// удаляем файл сокета если он существует
	if err := os.RemoveAll(pathSocket); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	// создаем unix сервер
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

				go func(conn net.Conn) {
					defer conn.Close()

					fmt.Println("подключен новый клиент", conn.LocalAddr())

					for {
						select {
						case <-ctx.Done():
							return

						default:
							if err = conn.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
								log.Println(err)

								return
							}

							scanner := bufio.NewScanner(conn)
							for scanner.Scan() {
								msg := scanner.Text()

								println("сообщение от клиента:", msg)

								if _, err = fmt.Fprintf(conn, "сообщение '%s' успешно получено", msg); err != nil {
									log.Println(err)

									return
								}
							}

							if err := scanner.Err(); err != nil {
								log.Printf("ошибка ввода: %s", err)
							}
						}
					}
				}(conn)
			}
		}
	}()

	return nil
}

// UnixSocketClient unix клиент
func UnixSocketClient(ctx context.Context, pathSocket string) (net.Conn, error) {
	conn, err := net.Dial("unix", pathSocket)
	if err != nil {
		return nil, err
	}

	// чтение ответов от сервера
	go func(conn net.Conn) {
		defer conn.Close()

		for {
			select {
			case <-ctx.Done():
				return

			default:
				buffer := make([]byte, 1024)
				num, err := conn.Read(buffer)
				if err != nil {
					log.Printf("ошибка: %s", err)
				}

				fmt.Println("сообщение от сервера:", string(buffer[:num]))
			}
		}
	}(conn)

	return conn, nil
}

// sendClientMessage прием сообщений из консоли и отправка на сервер
func sendClientMessage(cancel context.CancelFunc, conn net.Conn) {
	stdinScanner := bufio.NewScanner(os.Stdin)

	for stdinScanner.Scan() {
		msg := stdinScanner.Text()

		if msg == "exit" {
			cancel()

			return
		}

		if _, err := fmt.Fprintf(conn, "%s\n", msg); err != nil {
			log.Println(err)
		}
	}

	if err := stdinScanner.Err(); err != nil {
		log.Printf("ошибка ввода: %s", err)
	}
}
