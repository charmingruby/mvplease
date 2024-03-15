package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func New(logger *logrus.Logger) (*Config, error) {
	logger.Info("Loading configuration...")

	environment := environment{}
	err := env.Parse(&environment)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Database: &DatabaseConfig{
			User:     environment.DatabaseUser,
			Password: environment.DatabasePassword,
			Host:     environment.DatabaseHost,
			Name:     environment.DatabaseName,
			SSL:      environment.DatabaseSSL,
		},

		Server: &ServerConfig{
			Port: environment.ServerPort,
			Host: environment.ServerHost,
		},
		Cloudflare: &CloudflareConfig{
			AccountID:       environment.CloudflareAccountID,
			BucketName:      environment.AWSBucketName,
			AccessKeyID:     environment.AWSAccessKeyID,
			SecretAccessKey: environment.AWSSecretAccessKey,
		},
		Logger:       logger,
		JWTSecretKey: environment.JWTSecretKey,
	}
	logger.Info("Configuration done.")

	return cfg, nil
}

type environment struct {
	DatabaseUser        string `env:"DATABASE_USER,required"`
	DatabasePassword    string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost        string `env:"DATABASE_HOST,required"`
	DatabaseName        string `env:"DATABASE_NAME,required"`
	DatabaseSSL         string `env:"DATABASE_SSL,required"`
	ServerPort          string `env:"SERVER_PORT,required"`
	ServerHost          string `env:"SERVER_HOST,required"`
	JWTSecretKey        string `env:"JWT_SECRET_KEY,required"`
	CloudflareAccountID string `env:"CLOUDFLARE_ACCOUNT_ID,required"`
	AWSBucketName       string `env:"AWS_BUCKET_NAME,required"`
	AWSAccessKeyID      string `env:"AWS_ACCESS_KEY_ID,required"`
	AWSSecretAccessKey  string `env:"AWS_SECRET_ACCESS_KEY,required"`
}

type Config struct {
	Database     *DatabaseConfig
	Server       *ServerConfig
	Cloudflare   *CloudflareConfig
	Logger       *logrus.Logger
	JWTSecretKey string
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	SSL      string
	Conn     *sqlx.DB
}

type ServerConfig struct {
	Port string
	Host string
}

type CloudflareConfig struct {
	AccountID       string
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
}

func (c *Config) SetDatabase(db *sqlx.DB) {
	c.Database.Conn = db
}
