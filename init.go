package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ecdsa521/ircbot/lib"
	"github.com/ecdsa521/ircbot/scripts"
	"gopkg.in/yaml.v2"
)

func main() {

	var ch = make(chan string)
	var irc = ircbot.IrcConn{}
	data, _ := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal([]byte(data), &irc)

	ircbot.Init(&irc)
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
	for e, r := range irc.Events {
		match := r.FindStringSubmatch(event)
		if len(match) > 0 {
			//first, try to match the event line to all regexpes defined in lib/events.go
			//if any is found, call all scripts connected to it in scripts/*.go
			//example: message is matched by PRIVMSG regexp in events.go,
			//then we call all functions defined in scripts.Scripts["PRIVMSG"]
			mr := mapRegexp(match, r.SubexpNames())
			fmt.Printf("X: %v\nY: %v\nZ: %v\nR: %v\n---\n", e, match[1:], r.SubexpNames(), mr)
			for _, f := range scripts.Scripts[e] {
				f(irc, mr)
			}
		}
	}
	defer irc.WaitGroup.Done()
}
func mapRegexp(matchData []string, matchNames []string) map[string]string {
	var data = make(map[string]string)
	if len(matchData) != len(matchNames) {
		return data
	}

	for idx, val := range matchNames {
		if idx > 0 && idx <= len(matchData) {
			data[val] = matchData[idx]
		}

	}
	return data
}
