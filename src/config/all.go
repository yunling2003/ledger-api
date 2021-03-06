package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"

	"gopkg.in/yaml.v2"
)

var (
	// All contains all config
	All map[string]string
)

func init() {
	configFile := getConfigFile()
	bs, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf("could not find config file [%s]", configFile))
	}

	if err := yaml.Unmarshal(bs, &All); err != nil {
		panic(err.Error())
	}
	fmt.Println("*************************", "loaded config file at[", configFile, "]*************************")
}

func getConfigFile() string {
	// 1. check if env has configured the config file
	if configFile := os.Getenv("CONFIG_FILE"); configFile != "" && fileExists(configFile) {
		return configFile
	}

	// 2. get file by gin Mode
	mode := gin.Mode()
	configFile := fmt.Sprintf("./config/%s.yaml", mode)

	// will check three levels
	if ok := fileExists(configFile); !ok {
		configFile = upperLevel(configFile)
		if ok := fileExists(configFile); !ok {
			configFile = upperLevel(configFile)
			if ok := fileExists(configFile); !ok {
				configFile = upperLevel(configFile)
			}
		}
	}

	return configFile
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func upperLevel(fileName string) string {
	return "../" + fileName
}
