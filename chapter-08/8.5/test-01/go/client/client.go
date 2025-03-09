package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const MAX_CONNECTION_NUM = 50000

func buildConnect(lIp, sIp string, sPort int) (net.Conn, error) {
	localAddr, err := net.ResolveTCPAddr("tcp", lIp+":0")
	if err != nil {
		fmt.Println("\n Error: Could not resolve local address")
		return nil, err
	}

	remoteAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", sIp, sPort))
	if err != nil {
		fmt.Println("\n Error: Could not resolve server address")
		return nil, err
	}

	conn, err := net.DialTCP("tcp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println("\n Error: Connect Failed")
		return nil, err
	}

	return conn, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("\n Usage: %s <local ip> <server ip> <server port>\n", os.Args[0])
		return
	}

	lIp := os.Args[1]
	sIp := os.Args[2]
	sPort, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("\n Error: Invalid server port")
		return
	}

	sockets := make([]net.Conn, MAX_CONNECTION_NUM)
	for i := 1; i <= MAX_CONNECTION_NUM; i++ {
		if i%1000 == 0 {
			fmt.Printf("%s 连接 %s:%d成功了 %d 条！\n", lIp, sIp, sPort, i)
			time.Sleep(1 * time.Second)
		}

		conn, err := buildConnect(lIp, sIp, sPort)
		if err == nil {
			sockets[i-1] = conn
		} else {
			return
		}
	}
	time.Sleep(300 * time.Second)

	fmt.Println("关闭所有的连接...")
	for _, conn := range sockets {
		if conn != nil {
			conn.Close()
		}
	}
}
