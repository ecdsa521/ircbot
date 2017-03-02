package ircbot

import (
	"fmt"
)

//IrcEvent ...
type IrcEvent struct {
	Nick  string
	Ident string
	Host  string
}

//IrcConn ...
type IrcConn struct {
	Host   string
	Port   int
	UseSsl bool
	Ident  string
	Name   string
	Nick   string
}

func init() {
	fmt.Println("init ircbot")
}

//Connect ...
func Connect(irc *IrcConn) {
	fmt.Printf("connecting %v\n", irc)

}
