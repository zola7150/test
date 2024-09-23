package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

// decodeBase64 decodes a base64 encoded string and returns the decoded bytes.
func decodeBase64(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

// encodeBase64 encodes a byte slice into a base64 encoded string.
func encodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func main() {

	code := "abcdedfghijklmnopqrstuvwxyz"
	encoded := encodeBase64([]byte(code))
	fmt.Println("encoded string:", string(encoded))
	encoded2 := "YWJjZGVkZmdoaWprbG1ub3BxcnN0dXZ3eHl6"
	decoded, err := decodeBase64(encoded2)
	if err != nil {
		log.Fatal("Error decoding base64:", err)
	}
	fmt.Println("Decoded string:", string(decoded))
	//fmt.Println("value2:", totalCards[1])

	// 服務器的WebSocket地址
	// url := "ws://192.168.152.62:8080/ws/abcd"

	// // 創建一個websocket.Dialer實例
	// dialer := websocket.Dialer{
	// 	// 可以在這裡設置dialer選項
	// }

	// // 使用Dialer連接到服務器
	// conn, _, err := dialer.Dial(url, nil)
	// if err != nil {
	// 	log.Fatal("Error connecting to WebSocket server:", err)
	// }
	// defer conn.Close()

	// // 發送消息到服務器
	// err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from client"))
	// if err != nil {
	// 	log.Fatal("Error writing to WebSocket:", err)
	// }

	// // 循環接收消息
	// for {
	// 	_, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Println("Error reading from WebSocket:", err)
	// 		break
	// 	}
	// 	log.Printf("Received from server: %s", message)
	// }
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	// Connect to the server
// 	conn, err := net.Dial("tcp", "127.0.0.1:8080")
// 	// conn, err := net.Dial("tcp", "192.168.152.62:8080")
// 	if err != nil {
// 		fmt.Println("Error connecting:", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close()

// 	// Send a message to the server
// 	fmt.Fprintf(conn, "Hello from client\n")

// 	// Read the server's response
// 	response, err := bufio.NewReader(conn).ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading:", err)
// 		os.Exit(1)
// 	}

// 	// Print the server's response
// 	fmt.Print("Server says: " + response)
// }
