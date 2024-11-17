package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

var EnvConfigDefaults = map[string]any{
	"LOG_LEVEL":            "INFO",
	"ENCRYPT_ENABLED":      true,
	"ENCRYPT_PASSWORD":     "123456",
	"BACKUP_LOCAL_CRON":    "0 2 * * *",
	"BACKUP_LOCAL_DIR":     "/data/vault-sync-backup/",
	"BACKUP_REMOTE_CRON":   "",
	"BACKUP_REMOTE_TYPE":   "",
	"BACKUP_REMOTE_PARAMS": "",
}

type EnvConfig struct {
	LogLevel           string `mapstructure:"LOG_LEVEL"`
	EncryptEnabled     bool   `mapstructure:"ENCRYPT_ENABLED"`
	EncryptPassword    string `mapstructure:"ENCRYPT_PASSWORD"`
	BackupLocalCron    string `mapstructure:"BACKUP_LOCAL_CRON"`
	BackupLocalDir     string `mapstructure:"BACKUP_LOCAL_DIR"`
	BackupRemoteCron   string `mapstructure:"BACKUP_REMOTE_CRON"`
	BackupRemoteType   string `mapstructure:"BACKUP_REMOTE_TYPE"`
	BackupRemoteParams string `mapstructure:"BACKUP_REMOTE_PARAMS"`
}

func NewEnvConfig() (*EnvConfig, error) {
	// for dev
	if err := godotenv.Load(); err != nil {
		log.Println("Cannot loading .env file")
	}

	// read config from env
	setViperEnvDefaults()
	viper.AutomaticEnv()

	var config EnvConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func setViperEnvDefaults() {
	for k, v := range EnvConfigDefaults {
		viper.SetDefault(k, v)
	}
}
