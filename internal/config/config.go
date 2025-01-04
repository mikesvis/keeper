package config

import (
	"errors"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

// Config конфиг приложения
type Config struct {
	// Environment - окружение для сервера
	Environment string

	// ServerAddress - адрес сервера приложения
	ServerAddress string

	// DatabaseDSN - адрес подключения к базе postgres, нужен при выборе движка хранения коротких ссылок в базе.
	DatabaseDSN string

	// ServerCertPath - сертификат
	ServerCertPath string
}

// NewConfig Инициализация настроек сервера
func NewConfig() (*Config, error) {
	var configFilePath string
	flag.StringVarP(&configFilePath, "config", "c", "", "path to config file in yaml format")
	flag.Parse()

	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	environment := viper.GetString("environment")
	if environment != "production" {
		environment = "development"
	}

	serverAddress := viper.GetString("server_address")
	if serverAddress == "" {
		return nil, errors.New("no server address provided")
	}

	databaseDSN := viper.GetString("database_dsn")
	if databaseDSN == "" {
		return nil, errors.New("no database DSN provided")
	}

	serverCertPath := viper.GetString("server_cert_path")
	if serverCertPath == "" {
		return nil, errors.New("no path for server certificates is provided")
	}

	_, err = os.Stat(serverCertPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		Environment:    environment,
		ServerAddress:  serverAddress,
		DatabaseDSN:    databaseDSN,
		ServerCertPath: serverCertPath,
	}, nil
}
