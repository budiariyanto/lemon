package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

const operatingSystem = runtime.GOOS

type Config struct {
}

var pathMap map[string]string = map[string]string{
	"linux_default":   fmt.Sprintf("%s%s", "config", "/"),
	"darwin_default":  fmt.Sprintf("%s%s", "config", "/"),
	"windows_default": fmt.Sprintf("%s%s", "config", "\\"),
	"linux_home":      fmt.Sprintf("%s%s", os.Getenv("HOME"), "/"),
	"darwin_home":     fmt.Sprintf("%s%s", os.Getenv("HOME"), "/"),
	"windows_home":    fmt.Sprintf("%s%s", os.Getenv("HOMEPATH"), "\\"),
	"linux_global":    fmt.Sprintf("%s%s", "/etc", "/"),
	"darwin_global":   fmt.Sprintf("%s%s", "/etc", "/"),
	"windows_global":  fmt.Sprintf("%s%s", os.Getenv("SYSTEMROOT"), "\\"),
}

func (config *Config) CheckConfig(configFileName string) (string, error) {
	var defaultPath string = fmt.Sprintf("%s%s", pathMap[fmt.Sprintf("%s_%s", operatingSystem, "default")], configFileName)
	var homePath string = fmt.Sprintf("%s%s", pathMap[fmt.Sprintf("%s_%s", operatingSystem, "home")], configFileName)
	var globalPath string = fmt.Sprintf("%s%s", pathMap[fmt.Sprintf("%s_%s", operatingSystem, "global")], configFileName)

	if isConfigFound(defaultPath) {
		return defaultPath, nil
	} else if isConfigFound(homePath) {
		return homePath, nil
	} else if isConfigFound(globalPath) {
		return globalPath, nil
	} else {
		return "", errors.New("[FATAL ERROR]: Unable to find configuration file.")
	}
}

func isConfigFound(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (config *Config) ReadConfig(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return content, nil
}
