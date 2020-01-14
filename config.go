package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Favorites Favorites
}

type Favorites struct {
	NHL []string
}

func GetConfig() Config {
	file, err := ioutil.ReadFile(getConfigFilePath())
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func (c *Config) Update() {
	d, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(getConfigFilePath(), d, 0644)
}

func getConfigFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFile := fmt.Sprintf("%s/.scorekeep/config.yaml", usr.HomeDir)

	return configFile
}

func touchConfigFile() {
	file, err := os.OpenFile(getConfigFilePath(), os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
