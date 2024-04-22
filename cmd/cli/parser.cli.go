package cli

import (
	"log/slog"
	"os"
	"regexp"

	"github.com/patos-ufscar/http-web-server-example-go/common"
	"github.com/patos-ufscar/http-web-server-example-go/models"
	"gopkg.in/yaml.v2"
)

func ParseConfig(configPath string) ([]models.ServerConfig, error) {

	data, err := os.ReadFile(configPath)
    if err != nil {
		slog.Warn("could not read configPath, fallback to default")
		return nil, err
    }

	// fmt.Println(string(data))

	var conf Conf
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	// fmt.Println(conf)

	hostRegs := []regexp.Regexp{}
	for _, v := range conf.Servers[0].HostsRegs {
		r, err := regexp.Compile(common.ExtractRegExpFromHostStr(v))
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
		hostRegs = append(hostRegs, *r)
	}

	confs := []models.ServerConfig{}
	for _, v := range conf.Servers {
		locs := []models.LocationConfig{}
		for _, loc := range v.Locations {
			locs = append(locs, models.LocationConfig{
				Path: loc.Path,
				ReturnType: loc.ReturnType,
				Return: models.ReturnConfig{
					Code: loc.Return.Code,
					Headers: loc.Return.Headers,
					Body: []byte(loc.Return.Body),
				},
			})
		}
		confs = append(confs, models.ServerConfig{
			Port: v.Port,
			HostsRegs: hostRegs,
			Locations: locs,
		})
	}

	return confs, nil
}