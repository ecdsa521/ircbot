package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/ecdsa521/ircbot/lib"
	"gopkg.in/yaml.v2"
)

var rx = make(map[string]*regexp.Regexp)

func main() {

	var ch = make(chan string)
	var irc = ircbot.IrcConn{}
	data, _ := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal([]byte(data), &irc)

	ircbot.Connect(&irc)
	ircbot.JoinChannel(&irc, "#")
	irc.WaitGroup.Add(2)

	go func() {
		go ircbot.Poll(&irc, ch)
		for line := range ch {
			irc.WaitGroup.Add(1)
			go parser(&irc, strings.TrimSpace(line))
		}
		defer irc.WaitGroup.Done()
	}()

	irc.WaitGroup.Wait()
}

func parser(irc *ircbot.IrcConn, event string) {
	fmt.Println(event)
	for e, r := range rx {
		match := r.FindStringSubmatch(event)
		if len(match) > 0 {
			fmt.Printf("X: %v\nY: %v\nZ: %v\n---\n", e, match, r.SubexpNames())
		}
	}
	defer irc.WaitGroup.Done()
}
