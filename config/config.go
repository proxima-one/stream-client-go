package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
)

type Config struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Token      string `yaml:"token"`
	StreamID   string `yaml:"stream_id"`
	BufferSize int    `yaml:"buffer_size"`
}

func NewConfigFromYamlFile(filePath string) *Config {
	var config Config
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	return &config
}

func (c *Config) GetFullAddress() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
