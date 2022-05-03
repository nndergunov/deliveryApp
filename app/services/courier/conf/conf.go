package conf

import (
	"fmt"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"os"
)

func GetConf(name string) (string, error) {
	confPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fmt.Println(confPath)
	err = configreader.SetConfigFile(confPath + "/conf.json")
	if err != nil {
		return "", err
	}

	return configreader.GetString(name), nil

}
