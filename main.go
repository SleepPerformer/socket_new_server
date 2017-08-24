package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	//	"time"
)

func main() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "10.0.0.1:6060")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	downloadFile, _ := readFile()
	for {
		message, err := reader.ReadString('#')
		if err != nil {
			return
		}

		fmt.Println(string(message))
		//		msg := time.Now().String() + "\n"
		b := []byte(downloadFile)
		conn.Write(b)
	}
}

// 文件的内容二进制的形式存在[]byte中，实际长度为 int
func readFile() ([]byte, int) {
	userFile := "./test.txt"
	file, err := os.Open(userFile)
	defer file.Close()
	buf := make([]byte, 51200)
	if err != nil {
		fmt.Println(userFile, err)
		return buf, 0
	}
	// buf := make([byte, 2048])
	for {
		n, _ := file.Read(buf)
		// os.Stdout.Write(buf[:n])
		return buf, n
	}

}
