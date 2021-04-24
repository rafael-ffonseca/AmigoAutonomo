package config

import (
	"AmigoAutonomo/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const globalConfigPathENV = "GLOBAL_CONFIG_PATH"
const globalConfigFileName = "globalconfig.json"
const serviceSection = "Service"

func GetConfigServiceSection() (*models.Config, error) {
	return GetConfigSection(serviceSection)
}

func getConfigPath() (string, error) {
	var path = os.Getenv(globalConfigPathENV)
	fi, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		path = filepath.Join(path, globalConfigFileName)
	}
	return path, err
}

func readConfig() (map[string]interface{}, error) {
	file, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, err
}

func GetConfigSection(section string) (*models.Config, error) {
	var sectionTyped *models.Config
	config, err := readConfig()
	if err != nil {
		return nil, err
	}
	sectionUntyped, err := json.Marshal(config[section])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(sectionUntyped, &sectionTyped)
	if err != nil {
		return nil, err
	}
	return sectionTyped, nil
}

func GetPropFromServiceSection() *models.Config {
	configSection, errConf := GetConfigServiceSection()
	if errConf != nil {
		log.Fatalf("Error while reading config: %s", errConf.Error())
	}

	return configSection
}