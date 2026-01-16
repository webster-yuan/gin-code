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
	Port         string `mapstructure:"port"` // 告诉 Viper：YAML 里的 port → Go 里的 Port
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver string `mapstructure:"driver"` // 数据库驱动: mysql, sqlite3
	DSN    string `mapstructure:"dsn"`    // 数据库连接字符串
}

// LoggingConfig 日志配置
type LoggingConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

// AppConfig 提供一个全局可访问的配置实例
var AppConfig *Config

// LoadConfig 加载配置
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// 开发环境：第一个配置查找路径
	viper.AddConfigPath("./internal/config") // ./ 是程序启动的工作目录（在哪个目录执行 go run / go build / 可执行文件）

	// 环境变量自动读取
	viper.AutomaticEnv()

	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.read_timeout", 5)
	viper.SetDefault("server.write_timeout", 5)
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.dsn", "./data/app.db")

	if err := viper.ReadInConfig(); err != nil { // 读取配置
		log.Printf("无法读取配置文件: %v, 将使用默认值", err)
	}

	config := &Config{}
	// 将配置反序列化到结构体
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("无法解析配置: %v", err)
	}

	AppConfig = config
	return config
}
