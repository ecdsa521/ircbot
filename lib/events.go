package ircbot

import "regexp"

//Init regexp functions
func Init(irc *IrcConn) {
	irc.Events = make(map[string]*regexp.Regexp)
	irc.Events["PRIVMSG"] = regexp.MustCompile("^:(?P<Nick>.+?)!(?P<Ident>.+?)@(?P<Host>.+?) PRIVMSG (?P<Target>.+?) :(?P<Message>.+)")
	irc.Events["INVITE"] = regexp.MustCompile("^:(?P<Nick>.+?)!(?P<Ident>.+?)@(?P<Host>.+?) INVITE (?P<Target>.+?) :(?P<Channel>.+?)")
	irc.Events["PING"] = regexp.MustCompile("PING (?P<Ping>.*?)")
}
