package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:6060")
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		//timeouSec :=10
		//conn.
		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {

	// 缓冲区，存储被截断的数据
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)

	defer conn.Close()

	if err != nil && err == io.EOF {
		fmt.Println(buffer)
	}

}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			Log(string(data))
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
