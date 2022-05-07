package config

import (
	"github.com/proxima-one/streamdb-client-go/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
)

type Config struct {
	Host        string      `yaml:"host"`
	Port        int         `yaml:"port"`
	User        string      `yaml:"user"`
	Password    string      `yaml:"password"`
	Token       string      `yaml:"token"`
	StreamID    string      `yaml:"stream_id"`
	ChannelSize int         `yaml:"channel_size"`
	LogLevel    string      `yaml:"log_level"`
	State       model.State `yaml:"state"`
}

func (cfg *Config) GetHost() string {
	return cfg.Host
}

func (cfg *Config) GetPort() int {
	return cfg.Port
}

func (cfg *Config) GetUser() string {
	return cfg.User
}

func (cfg *Config) GetPassword() string {
	return cfg.Password
}

func (cfg *Config) GetToken() string {
	return cfg.Token
}

func (cfg *Config) GetStreamID() string {
	return cfg.StreamID
}

func (cfg *Config) GetChannelSize() int {
	return cfg.ChannelSize
}

func (cfg *Config) GetFullAddress() string {
	return cfg.Host + ":" + strconv.Itoa(cfg.Port)
}

func (cfg *Config) GetLogLevel() string {
	return cfg.LogLevel
}

func (cfg *Config) GetState() model.State {
	if cfg.State.Id == "" {
		return model.Genesis()
	}
	return cfg.State
}

func NewConfigFromYamlFile(filePath string) Config {
	config := &Config{}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	return *config
}

type OptionFunc func(*Config)

func NewConfigWithOptions(options ...OptionFunc) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func NewConfigFromFileOverwriteOptions(filepath string, options ...OptionFunc) Config {
	config := NewConfigFromYamlFile(filepath)
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithHost(host string) OptionFunc {
	return func(cfg *Config) {
		cfg.Host = host
	}
}

func WithPort(port int) OptionFunc {
	return func(cfg *Config) {
		cfg.Port = port
	}
}

func WithUser(user string) OptionFunc {
	return func(cfg *Config) {
		cfg.User = user
	}
}

func WithPassword(password string) OptionFunc {
	return func(cfg *Config) {
		cfg.Password = password
	}
}

func WithToken(token string) OptionFunc {
	return func(cfg *Config) {
		cfg.Token = token
	}
}

func WithStreamID(streamID string) OptionFunc {
	return func(cfg *Config) {
		cfg.StreamID = streamID
	}
}

func WithChannelSize(chanSize int) OptionFunc {
	return func(cfg *Config) {
		cfg.ChannelSize = chanSize
	}
}

func WithLogLevel(logLevel string) OptionFunc {
	return func(cfg *Config) {
		cfg.LogLevel = logLevel
	}
}

func WithState(state model.State) OptionFunc {
	return func(cfg *Config) {
		cfg.State = state
	}
}
