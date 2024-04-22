package models

import "regexp"

type ServerConfig struct {
	Port			uint16						`yaml:"port"`
	HostsRegs		[]regexp.Regexp				`yaml:"hosts"`
	Locations		[]LocationConfig			`yaml:"locations"`
}

type LocationConfig struct {
	Path			string						`yaml:"path"`
	ReturnType		string						`yaml:"returnType"`
	Return			ReturnConfig				`yaml:"return"`
}

type ReturnConfig struct {
	Code			int							`yaml:"code"`
	Headers			map[string]string			`yaml:"headers"`
	Body			[]byte						`yaml:"body"`
}
