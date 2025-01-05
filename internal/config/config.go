package config

import (
	"errors"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Environment     string
	ServerAddress   string
	DatabaseDSN     string
	ServerCertPath  string
	MinioBucketName string
	MinioEndpoint   string
	MinioAccessKey  string
	MinioSecretKey  string
	MinioUseSSL     bool
}

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

	minioBucketName := viper.GetString("minio_bucket_name")
	if minioBucketName == "" {
		return nil, errors.New("no minio bucket name is provided")
	}

	minioEndpoint := viper.GetString("minio_endpoint")
	if minioEndpoint == "" {
		return nil, errors.New("no minio endpoint is provided")
	}

	minioAccessKey := viper.GetString("minio_access_key")
	if minioAccessKey == "" {
		return nil, errors.New("no minio access key is provided")
	}

	minioSecretKey := viper.GetString("minio_secret_key")
	if minioSecretKey == "" {
		return nil, errors.New("no minio cert key is provided")
	}

	minioUseSSL := viper.GetBool("minio_use_ssl")

	return &Config{
		Environment:     environment,
		ServerAddress:   serverAddress,
		DatabaseDSN:     databaseDSN,
		ServerCertPath:  serverCertPath,
		MinioBucketName: minioBucketName,
		MinioEndpoint:   minioEndpoint,
		MinioAccessKey:  minioAccessKey,
		MinioSecretKey:  minioSecretKey,
		MinioUseSSL:     minioUseSSL,
	}, nil
}
