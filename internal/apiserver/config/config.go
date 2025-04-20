package config

import (
	"os"

	"github.com/ividernvi/iviuser/internal/pkg/options"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Options *options.Options
}

var (
	defaultConfig          *Config
	IVIUSER_MYSQL_HOSTNAME = os.Getenv("IVIUSER_MYSQL_HOSTNAME")
	IVIUSER_MYSQL_PORT     = os.Getenv("IVIUSER_MYSQL_PORT")
	IVIUSER_MYSQL_USERNAME = os.Getenv("IVIUSER_MYSQL_USERNAME")
	IVIUSER_MYSQL_PASSWORD = os.Getenv("IVIUSER_MYSQL_PASSWORD")
	IVIUSER_MYSQL_DATABASE = os.Getenv("IVIUSER_MYSQL_DATABASE")
)

var (
	IVIUSER_MINIO_ENDPOINT          = os.Getenv("IVIUSER_MINIO_ENDPOINT")
	IVIUSER_MINIO_ACCESS_KEY_ID     = os.Getenv("IVIUSER_MINIO_ACCESS_KEY_ID")
	IVIUSER_MINIO_SECRET_ACCESS_KEY = os.Getenv("IVIUSER_MINIO_SECRET_ACCESS_KEY")
	IVIUSER_MINIO_USE_SSL           = os.Getenv("IVIUSER_MINIO_USE_SSL")
	IVIUSER_MINIO_BUCKET_NAME       = os.Getenv("IVIUSER_MINIO_BUCKET_NAME")
)

func newConfig() *Config {
	return &Config{
		Options: options.NewOptions(),
	}
}

func DefaultConfig() *Config {
	if defaultConfig == nil {
		getDefault()
	}
	return defaultConfig
}

func getDefault() {
	checkDefaultEnv()
	defaultConfig = newConfig()
	defaultConfig.Options.MySQLOpts.HostName = IVIUSER_MYSQL_HOSTNAME
	defaultConfig.Options.MinioOpts.Endpoint = IVIUSER_MINIO_ENDPOINT
}

func checkDefaultEnv() {
	if IVIUSER_MYSQL_HOSTNAME == "" {
		logrus.Printf("database hostname is not set, please set IVIUSER_MYSQL_HOSTNAME")
		os.Exit(1)
	}

	if IVIUSER_MINIO_ENDPOINT == "" {
		logrus.Printf("minio endpoint is not set, please set IVIUSER_MINIO_ENDPOINT")
		os.Exit(1)
	}
}
