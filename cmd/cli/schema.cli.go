package cli

type Conf struct {
	Global			Global						`yaml:"global"`
	Servers			[]ServerConfig				`yaml:"servers"`
}

type Global struct {
	LogLevel		string						`yaml:"logLevel"`
	BufferSize		string						`yaml:"bufferSize"`
}

type ServerConfig struct {
	Port			uint16						`yaml:"port"`
	HostsRegs		[]string					`yaml:"hosts"`
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
	Body			string						`yaml:"body"`
}