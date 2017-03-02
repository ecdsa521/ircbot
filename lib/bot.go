package ircbot

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

//IrcEvent ...
type IrcEvent struct {
	Nick  string
	Ident string
	Host  string
}

//IrcConn ...
type IrcConn struct {
	Host      string
	Port      string
	UseSsl    bool
	Ident     string
	Name      string
	Nick      string
	Socket    net.Conn
	Buffer    *bufio.ReadWriter
	Reader    *bufio.Reader
	Writer    *bufio.Writer
	WaitGroup sync.WaitGroup
}

func init() {

}

//Connect ...
func Connect(irc *IrcConn) {
	var err error
	fmt.Printf("Connecting to %s:%s\n", irc.Host, irc.Port)
	if irc.Socket, err = net.Dial("tcp", irc.Host+":"+irc.Port); err == nil {
		fmt.Printf("Connected to %s:%s\n", irc.Host, irc.Port)
	} else {
		fmt.Printf("Error connecting: %v\n", err)
	}

	irc.Reader = bufio.NewReader(irc.Socket)
	irc.Writer = bufio.NewWriter(irc.Socket)
	//irc.Writer = bufio.NewWriter(&irc.Socket)
	irc.Buffer = bufio.NewReadWriter(irc.Reader, irc.Writer)
	time.Sleep(1 * time.Second)
	Write(irc, fmt.Sprintf("USER %s %s %s :%s", irc.Ident, irc.Ident, irc.Ident, irc.Ident))
	Write(irc, fmt.Sprintf("NICK %s", irc.Nick))
}

//Write to irc Socket
func Write(irc *IrcConn, message string) {
	fmt.Fprintf(irc.Socket, "%s\n", message)
}

//JoinChannel .. joins a channel
func JoinChannel(irc *IrcConn, channel string) {
	Write(irc, "JOIN "+channel)
}

//Poll ..
func Poll(irc *IrcConn, ch chan<- string) {

	var err error
	var msg string
	for {
		fmt.Println("polling")
		msg, err = irc.Buffer.ReadString('\n')
		if err != nil {
			break
		}
		ch <- fmt.Sprintf("%s\n", msg)

	}
	defer close(ch)
	fmt.Println("Out of loop")
	defer irc.WaitGroup.Done()

}
