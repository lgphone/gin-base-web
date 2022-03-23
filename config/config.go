package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

const (
	SessionCookieName  = "session_id"
	SessionRedisPrefix = "session:"
	SessionContextName = "sessionInfo"
)

type config struct {
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`
	Server struct {
		Port    int `yaml:"port"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
	Mysql struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		DB          string `yaml:"db"`
		TablePrefix string `yaml:"table_prefix"`
	} `yaml:"mysql"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	SqlLogFilePath string `yaml:"sql_log_file_path"`
	SessionTimeout int    `yaml:"session_timeout"`
	Debug          bool   `yaml:"debug"`
}

var Config *config

func Setup(configPath string) {
	if configPath == "" {
		configPath = "./config.yaml"
	}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err = d.Decode(&Config); err != nil {
		panic(err)
	}
	// 默认值
	if Config.SqlLogFilePath == "" {
		Config.SqlLogFilePath = "/tmp/slow.log"
	}
	if Config.SessionTimeout == 0 {
		Config.SessionTimeout = 3600
	}
}
