package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jaswdr/faker/v2"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	fake := faker.New()

	for {
		// クライアントから受け取ったデータをbufferに格納
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading:", err.Error())
			return
		}

		clientMessage := string(buffer[:n])
		fmt.Println("Received from client:", clientMessage)

		// fakeのレスポンス送信
		fakeResponse := fake.Lorem().Sentence(15)
		_, err = conn.Write([]byte(fakeResponse))
		if err != nil {
			log.Println("Error writing:", err.Error())
			return
		}

		fmt.Println("Sent to client:", fakeResponse)
	}
}

func main() {
	fmt.Println("Starting server...")

	socketPath := "/tmp/unix_socket"

	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatal("Error removing existing socket file:", err)
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
