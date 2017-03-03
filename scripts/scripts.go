package scripts

import (
	"fmt"

	ircbot "github.com/ecdsa521/ircbot/lib"
)

//Scripts hold map of functions to call
var Scripts = make(map[string][]func(irc *ircbot.IrcConn, event map[string]string))

func init() {
	fmt.Println("init from scripts.go")
}
