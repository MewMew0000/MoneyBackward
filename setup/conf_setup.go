package setup

import (
	"MoneyBackward/conf"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2" // yaml parsing library
)

func InitConfFrom(fPath string) (confList *conf.ListOfConf, retErr error) {
	confStream, err := os.ReadFile(fPath)
	if err != nil {
		return confList, fmt.Errorf("reading configuration file failed: %v", err)
	}

	confObj := &conf.ListOfConf{}
	if err := yaml.Unmarshal(confStream, confObj); err != nil {
		return confList, fmt.Errorf("parsing configuration file failed: %v", err)
	}

	log.Printf("Reading configuration file from %v\n", fPath)
	return confObj, nil
}
