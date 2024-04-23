package cli

type Conf struct {
	Global			Global						`yaml:"global"`
	Servers			[]ServerConfigInput				`yaml:"servers"`
}

type Global struct {
	LogLevel		string						`yaml:"logLevel"`
	BufferSize		string						`yaml:"bufferSize"`
}

type ServerConfigInput struct {
	Port			uint16						`yaml:"port"`
	HostsRegs		[]string					`yaml:"hosts"`
	Locations		[]LocationConfigInput			`yaml:"locations"`
}

type LocationConfigInput struct {
	Path			string						`yaml:"path"`
	ReturnType		string						`yaml:"returnType"`
	Return			map[string]interface{}		`yaml:"return"`
}

type ReturnStaticConfig struct {
	Code			int							`yaml:"code"`
	Headers			map[string]string			`yaml:"headers"`
	Body			string						`yaml:"body"`
}
