package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Bind на порт ОС
	listener, _ := net.Listen("tcp", ":5000")

	for {
		// ждём пока не придёт клиент
		conn, _ := listener.Accept()

		fmt.Println("Connected")

		// создаём Reader для чтения информации из сокета
		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")
		go func() {//реализуем асинхронность подключений
			defer conn.Close()

			for {
				// побайтово читаем
				rbyte, _ := bufReader.ReadByte()

				fmt.Print(string(rbyte))
			}
		}()
	}
}
