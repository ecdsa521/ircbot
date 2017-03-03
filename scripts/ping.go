package scripts

import (
	"fmt"

	"github.com/ecdsa521/ircbot/lib"
)

func init() {
	Scripts["PRIVMSG"] = append(Scripts["PRIVMSG"], pingCommand)
}

func pingCommand(irc *ircbot.IrcConn, event map[string]string) {
	fmt.Printf("PRIVMSG -> pingCommand: %v\n", event)
}
