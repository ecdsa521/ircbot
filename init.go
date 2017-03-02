package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/ecdsa521/ircbot/lib"
)

func main() {

	var ch = make(chan string)
	var irc = ircbot.IrcConn{
		Nick:  "moobot",
		Ident: "moobot",
		Host:  "127.0.0.1",
		Port:  "6667",
	}

	ircbot.Connect(&irc)
	ircbot.JoinChannel(&irc, "#")
	irc.WaitGroup.Add(2)

	go func() {
		go ircbot.Poll(&irc, ch)
		for line := range ch {
			if matches, err := regexp.MatchString("INVITE", line); err == nil && matches {
				ircbot.JoinChannel(&irc, "#")
			}
			fmt.Println(line)
		}
		defer irc.WaitGroup.Done()
	}()

	time.Sleep(1 * time.Hour)
	irc.WaitGroup.Wait()
}
