package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const _envPrefix = "quick"

type Config struct {
	Server Server `mapstructure:"server" validate:"required"`
	HTTP   HTTP   `mapstructure:"http" validate:"required"`
}

type Server struct {
	Address      string        `mapstructure:"address" validate:"required,hostname_port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout" validate:"required"`
	WriteTimeout time.Duration `mapstructure:"write-timeout" validate:"required"`
	IdleTimeout  time.Duration `mapstructure:"idle-timeout" validate:"required"`
}

type HTTP struct {
	BodyLimitSize string  `mapstructure:"body-limit-size"`
	CORS          CORS    `mapstructure:"cors"`
	Recover       Recover `mapstructure:"recover"`
}

type CORS struct {
	AllowedOrigins   []string `mapstructure:"allowed-origins"`
	AllowedHeaders   []string `mapstructure:"allowed-headers"`
	AllowedMethods   []string `mapstructure:"allowed-methods"`
	AllowCredentials bool     `mapstructure:"allow-credentials"`
	ExposedHeaders   []string `mapstructure:"exposed-headers"`
	MaxAge           int      `mapstructure:"max-age"`
}

type Recover struct {
	StackSize         int  `mapstructure:"stack-size"`
	DisableStackAll   bool `mapstructure:"disable-stack-all"`
	DisablePrintStack bool `mapstructure:"disable-print-stack"`
}

func InitViper(configPath string) (*Config, error) {
	var c Config

	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(_builtinConfig))); err != nil {
		return nil, fmt.Errorf("error loading default configs: %w", err)
	}

	v.SetConfigFile(configPath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.SetEnvPrefix(_envPrefix)
	v.AutomaticEnv()

	switch err := v.MergeInConfig(); err.(type) {
	case nil:
	case *os.PathError:
		logrus.Infof("config file (%s) not found, Using defaults and environment variables", configPath)
	default:
		logrus.Warnf("failed to load config file: %s", err)
	}

	if err := v.UnmarshalExact(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config into struct: %w", err)
	}

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &c, nil
}

func (c *Config) Validate() error {
	v := validator.New()
	return v.Struct(c)
}
