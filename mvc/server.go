package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan <- string

var (
	//用于通信的三个通道
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

//监听客户端的到来和离开
func boardcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <- messages: //当接收到消息的情况
			for cli := range clients {
				cli <- msg
			}
		case cli := <- entering://当收到客户端到达的通知
		     clients[cli] = true
		case cli := <- leaving:
			 delete(clients,cli)
			 close(cli)
		}
	}
}

func clientwriter(conn net.Conn,ch <- chan string) {
	for msg := range ch {
		fmt.Fprintln(conn,msg)
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) //创建一个STRING类型的管道
	go clientwriter(conn, ch)
	who := conn.RemoteAddr().String() //本次

	ch <- "you are" + who //将你的到达信息传入在通道中
	messages <- who + "has arrived" //放入在消息队列中
	entering <- ch //将信息放入在到达队列中
	input := bufio.NewScanner(conn)//从键盘中的输入
	for input.Scan() {//有输入信息就把信息放在输入队列中
		messages <- who + ":" + input.Text()
	}

	leaving <- ch//处理结束之后将信息放在离开队列中
	messages <- who + "has left" //将对应的信息放在消息队列中
	conn.Close()//关闭链接
}

func main() {
	//开启监听8000端口
	listener, err := net.Listen("tcp","localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	//执行监听程序
	//监听客户端的到达和离开
	go boardcast()
	//监听与回复的循环
	for {
		conn, err := listener.Accept() //得到到达的信息
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}


}