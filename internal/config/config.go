package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config 应用配置结构体
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

// LoggingConfig 日志配置
type LoggingConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

var AppConfig *Config

// LoadConfig 加载配置
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("/etc/gin/")
	viper.AutomaticEnv()

	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.read_timeout", 5)
	viper.SetDefault("server.write_timeout", 5)
	viper.SetDefault("logging.level", "info")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("无法读取配置文件: %v, 将使用默认值", err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("无法解析配置: %v", err)
	}

	AppConfig = config
	return config
}
