package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"sync"
)

const CONFIG_NAME = "config.yml"

var once sync.Once
var instance *Config

type App struct {
	Environment string `yaml:"environment"`
}

type Metrics struct {
	Port   string `yaml:"port"`
	Handle string `yaml:"handle"`
}

type REST struct {
	Port         string `yaml:"port"`
	EndpointPort string `yaml:"endpoint-port"`
	Host         string `yaml:"host"`
}

type Jaeger struct {
	Port string `yaml:"port"`
	Name string `yaml:"name"`
	Host string `yaml:"host"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db-name"`
	SSLMode  string `yaml:"ssl-mode"`
}

type Config struct {
	App      App      `yaml:"app"`
	Database Database `yaml:"database"`
	REST     REST     `yaml:"rest"`
	Metrics  Metrics  `yaml:"metrics"`
	Jaeger   Jaeger   `yaml:"jaeger"`
}

func createConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Err(err).Msg("error open config file")
		}

		file = []byte(os.ExpandEnv(string(file)))
		if err := yaml.Unmarshal(file, instance); err != nil {
			log.Err(err).Msg("error while decode config file")
		}
	})

	return instance
}

func GetConfig() *Config {
	return createConfig(CONFIG_NAME)
}

func InitConfig(path string) *Config {
	return createConfig(path)
}
