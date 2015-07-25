package config

import (
    "errors"
    "github.com/pelletier/go-toml"
    "log"
    "os"
    "runtime"
    "strings"
)

const operatingSystem = runtime.GOOS

type Config struct {
    Path string
    
}

func checkConfig(fileName string) (string, error) {
    var configFileName = fileName
    var defaultPath string = strings.Join([]string{getDefaultPath(), configFileName}, "")
    var homePath string = strings.Join([]string{getHomePath(), configFileName}, "")
    var globalPath string = strings.Join([]string{getGlobalPath(), configFileName}, "")

    if defaultPathInfo, _ := os.Stat(defaultPath); defaultPathInfo != nil {
        return defaultPath, nil
    } else if homePathInfo, _ := os.Stat(homePath); homePathInfo != nil {
        return homePath, nil
    } else if globalPathInfo, _ := os.Stat(globalPath); globalPathInfo != nil {
        return globalPath, nil
    } else {
        return "", errors.New("[FATAL ERROR]: Unable to find configuration file.")
    }
}


func getHomePath() string {    
    var path string

    switch operatingSystem {
    case "linux":
        path = strings.Join([]string{os.Getenv("HOME"), "/"}, "")
    case "darwin":
        path = strings.Join([]string{os.Getenv("HOME"), "/"}, "")
    case "windows":
        path = strings.Join([]string{os.Getenv("HOMEPATH"), "\\"}, "")
    }

    return path
}

func getGlobalPath() string {
    var path string

    switch operatingSystem {
    case "linux":
        path = strings.Join([]string{os.Getenv("/etc"), "/"}, "")
    case "darwin":
        path = strings.Join([]string{os.Getenv("/etc"), "/"}, "")
    case "windows":
        path = strings.Join([]string{os.Getenv("SYSTEMROOT"), "\\"}, "")
    }

    return path
}

func getDefaultPath() string {
    var path string

    switch operatingSystem {
    case "linux":
        path = strings.Join([]string{"config", "/"}, "")
    case "darwin":
        path = strings.Join([]string{"config", "/"}, "")
    case "windows":
        path = strings.Join([]string{"config", "\\"}, "")
    }

    return path
}

func (conf *Config) Load(fileName string) *toml.TomlTree {
    configFile, err := checkConfig(fileName)
    if err != nil {
        log.Fatal(err)
    }

    config, err := toml.LoadFile(configFile)
    return config
}

/*func main() {
    var config Config
    tomlConfig := config.Load()
    fmt.Println(tomlConfig.Get("database.name").(string))
}*/