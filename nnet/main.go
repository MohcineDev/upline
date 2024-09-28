package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const welcomeMsg = "Welcome to TCP-Chat!\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n"

type client struct {
	conn     *net.Conn
	userName string
}

var (
	clients    []client
	historyMsg []string
)
var port = "8989"

func main() {
	myArgs := os.Args[1:]

	// handle Args Error
	if len(myArgs) >= 2 {
		log.Fatalln("[USAGE]: ./TCPChat $port")
	}
	///assign the first arg to the port to use
	if len(myArgs) == 1 {
		port = myArgs[0]
	}
	listener, err := net.Listen("tcp", ":"+port)
	///if the port is incorrect or something happens while trying to connect on that port
	if err != nil {
		log.Fatalln("error connecting on port", port)
	}
	// if there is no errors print the listening to terminal
	fmt.Println("listening on ", port, "...")
	// connNum := 0

	msgs := make(chan Message)
	go server(msgs)
	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(errors.New("Error: could not accept connection!"))
		}

		msgs <- Message{
			Type: ClientConnected,
			Conn: conn,
		}

		go handleConnection(conn, msgs)
		// connNum++
		// if connNum <= 2 {
		// 	fmt.Println("num of conn exceeded")
		// }
	}
}

func handleConnection(myConn net.Conn, messages chan Message) {
	////make a buffer to hold the req
	temp := make([]byte, 1024)
	packet := make([]byte, 1024)
	name := make([]byte, 1024)

	defer myConn.Close()

	///write the welcome msg to users
	_, err := myConn.Write([]byte(welcomeMsg))
	if err != nil {
		log.Printf("could not write the welcome massage %s", err)
		return
	}
	myConn.Write([]byte("[ENTER YOUR NAME]:"))

	myConn.Read(name)

	fmt.Println("name : ", string(name))

	for {
		_, err := myConn.Read(temp)

		if err != nil {
			fmt.Println("connection err")
			myConn.Close()
			return
		}

		messages <- Message{
			Type: NewMessage,
			Text: string(temp),
		}
		packet = append(packet, temp...)
	}
	c1 := &client{
		conn: &myConn,
	}
	clients = append(clients, *c1)

	///take the recieved data and write it to a res
	time := time.Now().Format(time.ANSIC)
	resString := fmt.Sprintf("Your Msg: %v , [%v]", string(packet), time)
	myConn.Write([]byte(resString))
}

type MessageType int

type Message struct {
	Type MessageType
	Conn net.Conn
	Text string
}

const (
	ClientConnected MessageType = iota + 1
	NewMessage
)

func server(messages chan Message) {
	users := []net.Conn{}
	for {
		msg := <-messages
		switch msg.Type {
		case ClientConnected:
			users = append(users, msg.Conn)
		case NewMessage:
			for _, user := range users {
				if user.RemoteAddr().String() != msg.Conn.RemoteAddr().String() {

					_, err := user.Write([]byte(msg.Text))
					if err != nil {
						fmt.Println("couldn't send data to user")
					}
				}
			}
		}

	}

}
