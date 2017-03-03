package main

import "regexp"

func init() {
	rx["PRIVMSG"] = regexp.MustCompile("^:(?P<Nick>.+?)!(?P<Ident>.+?)@(?P<Host>.+?) PRIVMSG (?P<Target>.+?) :(?P<Message>.+)")
	rx["INVITE"] = regexp.MustCompile("^:(.*)!(.*)@(.*) INVITE (.*) :(.*)")
}
