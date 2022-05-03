package conf

import (
	"os"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
)

func GetConf(name string) (string, error) {
	confPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	err = configreader.SetConfigFile(confPath + "/conf.json")
	if err != nil {
		return "", err
	}

	return configreader.GetString(name), nil
}
