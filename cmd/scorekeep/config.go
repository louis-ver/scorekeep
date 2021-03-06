package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/louis-ver/scorekeep/pkg"
	"gopkg.in/yaml.v2"
)

const (
	NHL            = "nhl"
	NBA            = "nba"
	configDirname  = ".scorekeep"
	configFilename = "config"
)

type Config struct {
	ServerUrl string `yaml:"server-url"`
	Timezone  string
	Favorites Favorites
}

type Favorites struct {
	NHL []string
	NBA []string
}

func CreateConfigDirAndFile() {
	os.MkdirAll(getConfigDirectory(), 0755)
	file, err := os.OpenFile(getConfigFilePath(), os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
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

func (c *Config) WriteToFile() {
	d, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(getConfigFilePath(), d, 0644)
}

func (c *Config) GetFavorites(league string) []string {
	var favorites []string
	switch league {
	case NHL:
		favorites = GetConfig().Favorites.NHL
	case NBA:
		favorites = GetConfig().Favorites.NBA
	}
	return favorites
}

func (c *Config) AddFavorite(f string, league string) {
	switch league {
	case NHL:
		if !pkg.StringInSlice(f, c.Favorites.NHL) {
			c.Favorites.NHL = append(c.Favorites.NHL, f)
		}
	default:
		log.Print("League not supported")
		// TODO: Add return error when league not supported
	}
}

func getConfigFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFile := fmt.Sprintf("%s/%s/%s", usr.HomeDir, configDirname, configFilename)

	return configFile
}

func getConfigDirectory() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/%s", usr.HomeDir, configDirname)
}
