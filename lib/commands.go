package ircbot

import "fmt"

//JoinChannel .. joins a channel
func JoinChannel(irc *IrcConn, channel string) {
	Write(irc, "JOIN "+channel)
}

//PartChannel .. parts channel
func PartChannel(irc *IrcConn, channel string) {
	Write(irc, "PART "+channel)
}

//Privmsg .. send message to either person or channel
func Privmsg(irc *IrcConn, target string, message string) {
	Write(irc, fmt.Sprintf("PRIVMSG %s :%s", target, message))
}
