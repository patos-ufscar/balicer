package cli

import (
	"errors"
	"fmt"
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

	fmt.Println(conf)

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
		locs := []models.HandlerConfig{}
		for _, loc := range v.Locations {
			locs = append(locs, models.HandlerConfig{
				Path: loc.Path,
				ReturnType: loc.ReturnType,
				Return: loc.Return,
				// models.ReturnConfig{
				// 	Code: loc.Return.Code,
				// 	Headers: loc.Return.Headers,
				// 	Body: []byte(loc.Return.Body),
				// },
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

func ParseStaticReturn(ret map[string]interface{}) (*models.ReturnStatic, error) {

	code, ok := ret["code"].(int)
	if !ok {
		return nil, errors.New("could not convert re.code to int")
	}

	headers, err := ConvertMap(ret["headers"])
	if err != nil {
		return nil, errors.New("could not convert re.headers to map[string]string")
	}

	body, ok := ret["body"].(string)
	if !ok {
		return nil, errors.New("could not convert re.body to string")
	}

	staticRet := models.ReturnStatic{
		Code: code,
		Headers: headers,
		Body: []byte(body),
	}

	
	return &staticRet, nil
}

func ConvertMap(input interface{}) (map[string]string, error) {
    result := make(map[string]string)

	inputMap, ok := input.(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("could not convert input to map[interface{}]interface{}")
	}

    for key, value := range inputMap {
        // Perform type assertions
        strKey, okKey := key.(string)
        strValue, okValue := value.(string)

        // Check if both key and value are strings
        if okKey && okValue {
            result[strKey] = strValue
        } else {
            // If either key or value is not a string, return an error
            return nil, fmt.Errorf("key or value is not a string")
        }
    }

    return result, nil
}
