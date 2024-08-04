package utils

import (
	"html/template"
	"log/slog"
	"os"
	"strings"
)

var LOG_LEVEL string = strings.ToUpper(GetEnvVarDefault("LOG_LEVEL", "INFO"))
var log_set bool = false

var TemplateAccountConfirmation string = "/srv/html/templates/account_confirmation.ptbr.html"

func InitSlogger() {

	if log_set {
		return
	}

	levelsMap := map[string]slog.Level{
		"DEBUG":   slog.LevelDebug,
		"INFO":    slog.LevelInfo,
		"WARN":    slog.LevelWarn,
		"WARNING": slog.LevelWarn,
		"ERROR":   slog.LevelError,
	}

	logger := slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     levelsMap[LOG_LEVEL],
		},
	))

	slog.SetDefault(logger)

	log_set = true
}

// Gets the `envVarName`, returns defaultVal if envvar is non-existant.
func GetEnvVarDefault(envVarName string, defaultVal string) string {
	envVar := os.Getenv(envVarName)

	if envVar == "" {
		return defaultVal
	}

	return envVar
}

// Removes all occurences of item in slice
func RemoveFrom[T comparable](slice []T, item T) []T {
	var newSlice []T
	for _, v := range slice {
		if v != item {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

// Load html template for email
func LoadHTMLTemplate(templateName string) *template.Template {

	t, err := template.ParseFiles(templateName)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	return t
}

func IsSubset(subset []string, superset []string) bool {
	checkMap := make(map[string]bool)
	for _, element := range superset {
		checkMap[element] = true
	}
	for _, value := range subset {
		if !checkMap[value] {
			return false // Return false if an element is not found in the superset
		}
	}
	return true // Return true if all elements are found in the superset
}
