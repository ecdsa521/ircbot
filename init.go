package main

import (
	"fmt"

	"github.com/ecdsa521/ircbot/lib"
)

func main() {

	var irc = ircbot.IrcConn{
		Nick:  "test",
		Ident: "test",
		Host:  "irc.freenode.net",
	}

	ircbot.Connect(&irc)
	fmt.Printf("%v\n", irc)
}
