package main

import (
	"bufio" //缓存i/O接口
	"fmt"   //输入与输出
	"net"   //通讯包
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")//创建一个套接字对象, 指定其IP以及端口

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)        //开始监听套接字指定的端口.

	defer tcpListener.Close()        //执行完后关闭TCP

	for {
		/*
		如有新的客户端连接请求, 则建立一个goroutine,
		在goroutine中, 读取客户端消息, 并转发回去, 直到客户端断开连接
		*/
		tcpConn, err := tcpListener.AcceptTCP()        //接收一个请求

		if err != nil {
			continue //如果接收出错、继续监听下个请求

			fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
			地址        //输出客户端的IP
			go tcpPipe(tcpConn)        //调用tcpPipe函数处理客户端发送的数据
		}

	}

	func tcpPipe(conn * net.TCPConn)
	{
		ipStr := conn.RemoteAddr().String() //获取客户端ip地址
		defer func() {
			fmt.Println("disconnected :" + ipStr)
			conn.Close()
		}()
		reader := bufio.NewReader(conn)        //以缓冲的方式读取客户端数据

		for {
			message, err := reader.ReadString('\n')        //按行从缓冲区读取数据，
			if err != nil {
				return
			}

			fmt.Println(string(message))        //输出数据
			msg := time.Now().String() + "\n"        //获取时间
			b := []byte(msg)
			conn.Write(b)        //将数据返回客户端
		}
	}
