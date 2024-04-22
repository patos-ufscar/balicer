package models

import "regexp"

type ServerConfig struct {
	Port			uint16
	HostsRegs		[]regexp.Regexp
	Locations		[]HandlerConfig
}

type HandlerConfig struct {
	Path			string
	ReturnType		string
	Return			ReturnConfig
}

type ReturnConfig struct {
	Code			int
	Headers			map[string]string
	Body			[]byte
}
