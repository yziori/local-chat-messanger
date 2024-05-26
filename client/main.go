package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// ソケットパスの設定
	socketPath := "/tmp/unix_socket"

	// サーバーへ接続
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// 標準入力からの読み取り
	reader := bufio.NewReader(os.Stdin)

	for {
		// ユーザーメッセージを読み取りサーバーに送信する
		fmt.Println("Enter Message: ")
		userMessage, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(userMessage))
		if err != nil {
			log.Println("Error writing to server:", err)
		}

		// サーバーからの応答を受信して表示する
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			log.Println("Error reading:", err)
		}
		fmt.Println("Recieved from Server:", string(response[:n]))
	}

}
