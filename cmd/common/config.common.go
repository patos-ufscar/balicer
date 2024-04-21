package common

import (
	"regexp"
	"strings"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

func ParseConfig(configPath string) []models.ServerConfig {

	regs := []regexp.Regexp{}
	// regs = append(regs, ParseHostConfig(".+"))

	locs := []models.LocationConfig{}

	confs := []models.ServerConfig{}
	confs = append(confs, models.ServerConfig{
		Port: 4221,
		HostsRegs: regs,
		Locations: locs,
	})

	return confs

	// return nil
}

func ParseHostConfig(hostStr string) regexp.Regexp {
	str := ExtractRegExpFromHostStr(hostStr)

	return *regexp.MustCompile(str)
}

func ExtractRegExpFromHostStr(hostStr string) string {

	if strings.HasPrefix(hostStr, "r`") && strings.HasSuffix(hostStr, "`") {
		return hostStr[1+1:len(hostStr)-1]
	}

	switch hostStr {
	case "*":
		return `.+`
	default:
		return strings.ReplaceAll(hostStr, ".", "\\.")
	}
}