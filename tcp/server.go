package tcp

import (
	"fmt"
	"net"
)

func TcpServer() {
	fmt.Println("server has been start===>")
	tcdAddr, _ := net.ResolveTCPAddr("tcp", ":8000")
	//服务器端一般不定位具体的客户端套接字
	tcpListener, _ := net.ListenTCP("tcp", tcdAddr)

	connMap := make(map[string]*net.TCPConn)
	for {
		tcpConn, _ := tcpListener.AcceptTCP()
		defer tcpConn.Close()

		connMap[tcpConn.RemoteAddr().String()] = tcpConn
		fmt.Println("连接的客户端信息：", tcpConn.RemoteAddr().String())
	}
}
