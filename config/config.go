package config

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	AppName        string `mapstructure:"app_name" validate:"required"`
	AppVersion     string `mapstructure:"app_version" validate:"required"`
	AppEnvironment string `mapstructure:"app_environment" validate:"oneof=development staging production"`

	ServerHost       string        `mapstructure:"server_host" validate:"required"`
	ServerPort       string        `mapstructure:"server_port" validate:"required"`
	ServerTimeout    time.Duration `mapstructure:"server_timeout" validate:"required"`
	ServerProtocol   string        `mapstructure:"server_protocol" validate:"required,oneof=http https"`
	ServerApiPrefix  string        `mapstructure:"server_api_prefix" validate:"required"`
	ServerApiVersion string        `mapstructure:"server_api_version" validate:"required"`

	DatabaseHost     string `mapstructure:"database_host" validate:"required"`
	DatabasePort     int    `mapstructure:"database_port" validate:"required,min=1,max=65535"`
	DatabaseUsername string `mapstructure:"database_username" validate:"required"`
	DatabasePassword string `mapstructure:"database_password" validate:"required"`
	DatabaseName     string `mapstructure:"database_name" validate:"required"`
	DatabaseSSLMode  string `mapstructure:"database_ssl_mode" validate:"required,oneof=disable"`
	DatabaseTimezone string `mapstructure:"database_timezone" validate:"required"`

	RedisHost     string `mapstructure:"redis_host" validate:"required"`
	RedisPort     string `mapstructure:"redis_port" validate:"required,min=1,max=65535"`
	RedisPassword string `mapstructure:"redis_password" validate:"required"`

	LoggingEnabled bool   `mapstructure:"logging_enabled" validate:"required"`
	LoggingLevel   string `mapstructure:"logging_level" validate:"oneof=debug info warn error"`
	LoggingFormat  string `mapstructure:"logging_format" validate:"oneof=json text"`

	IdGeneratorNodeID  int64         `mapstructure:"id_generator_node_id" validate:"required"`
	AccessTokenSecret  string        `mapstructure:"access_token_secret" validate:"required"`
	RefreshTokenSecret string        `mapstructure:"refresh_token_secret" validate:"required"`
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry" validate:"required"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry" validate:"required"`
	JWTIssuer          string        `mapstructure:"jwt_issuer" validate:"required"`
	JWTAudience        string        `mapstructure:"jwt_audience" validate:"required"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, errors.Wrap(err, "config validation failed")
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return &cfg, nil
}
