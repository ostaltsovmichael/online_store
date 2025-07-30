package configs

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dependens    string `yaml:"dependens"`
	Host         string `yaml:"host"`
	Port         int16  `yaml:"port"`
	Postgres_url string `yaml:"postgres_url"`
	Sqlite_path  string `yaml:"sqlite_path"`
	Langl_zone   string `yaml:"langl_zone"`
	Secret_jwt_key string `yaml:"secret_jwt_key"`
}

func (Config) getConfig() Config {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		slog.Error("config file not readable ")
	}

	var config Config

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		slog.Error("bad parse config.yaml file")
	}
	return config
}

var Cfg = Config.getConfig(Config{})
