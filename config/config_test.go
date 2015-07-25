package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var sampleConfig string = `[database]
name="postgresql"
host="localhost"
port=5432
dbName="db_cheatsheet"
username="budi"
password="budibodonk"
connectionString="postgresql://budi:budibodonk@localhost/db_cheatsheet"

[log]
logPath="/path/to/logfile.log"
`

func TestFindConfigFile(t *testing.T) {
	config := new(Config)
	defaultFile := "testconfig.toml"
	homeFile := fmt.Sprintf("%s%s", pathMap[fmt.Sprintf("%s_%s", operatingSystem, "home")], defaultFile)

	ioutil.WriteFile(defaultFile, []byte(sampleConfig), os.ModePerm)
	ioutil.WriteFile(homeFile, []byte(sampleConfig), os.ModePerm)

	_, err := config.CheckConfig(defaultFile)

	if err != nil {
		t.Error(err)
	}

	os.Remove(defaultFile)
	os.Remove(homeFile)
}

func TestReadConfig(t *testing.T) {
	config := new(Config)
	defaultFile := "testconfig.toml"

	ioutil.WriteFile(defaultFile, []byte(sampleConfig), os.ModePerm)

	content, _ := config.ReadConfig(defaultFile)

	if string(content) != sampleConfig {
		t.Error("Read config file failed")
	}

	os.Remove(defaultFile)

}
