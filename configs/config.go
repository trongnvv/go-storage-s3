package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port        string      `yaml:"port"`
	Prefix      string      `json:"prefix"`
	ServiceName string      `yaml:"service_name"`
	Jaeger      *Jaeger     `yaml:"jaeger"`
	Mode        string      `yaml:"mode"`
	Postgresql  *Postgresql `yaml:"postgresql"`
}

type Jaeger struct {
	Endpoint         string   `yaml:"endpoint"`
	Active           bool     `yaml:"active"`
	PathIgnoreLogger []string `yaml:"path_ignore_logger"`
}

type Postgresql struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	DbName      string `yaml:"db_name"`
	SslMode     string `yaml:"ssl_mode"`
	Password    string `yaml:"password"`
	MaxLifeTime int64  `yaml:"max_life_time"`
	AutoMigrate bool   `yaml:"auto_migrate"`
}

var config *Config

func Get() *Config {
	return config
}

func LoadConfig(path string) {
	//configFile, err := os.Open(path)
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
	//byteValue, err := ioutil.ReadAll(configFile)
	//if err = yaml.Unmarshal(byteValue, &config); err != nil {
	//	panic(err)
	//}
}
