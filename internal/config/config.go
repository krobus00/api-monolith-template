package config

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	ServiceName    = ""
	ServiceVersion = ""
)

var (
	Env *EnvConfig
)

type EnvConfig struct {
	Env                     string        `mapstructure:"env"`
	LogLevel                string        `mapstructure:"log_level"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
	Token                   Token         `mapstructure:"token"`
	Server                  Server        `mapstructure:"server"`
	Database                Database      `mapstructure:"database"`
}

type Token struct {
	PasswordSalt string `mapstructure:"password_salt"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Database struct {
	DSN             string        `mapstructure:"dsn"`
	PingInterval    time.Duration `mapstructure:"ping_interval"`
	ReconnectFactor float64       `mapstructure:"reconnect_factor"`
	MinJitter       time.Duration `mapstructure:"min_jitter"`
	MaxJitter       time.Duration `mapstructure:"max_jitter"`
	MaxRetry        int           `mapstructure:"max_retry"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("failed to read config file: ", err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		logrus.Fatal("failed to unmarshal config file: ", err)
		return err
	}

	return nil
}
