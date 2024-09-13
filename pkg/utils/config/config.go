package config

import (
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

const (
	AppConfigDefaultName = "psapi"

	KeyAppLogLevel     = "logLevel"
	DefaultAppLogLevel = "INFO"

	KeyApiCors     = "api.cors"
	DefaultApiCors = "*"

	KeyApiPort     = "api.port"
	DefaultApiPort = 8000

	KeyImportDataFolderPath     = "import.dataFolderPath"
	DefaultImportDataFolderPath = "data"
)

type AppConfig struct {
	LogLevel string
	Api      ApiConfig
	Import   ImportConfig
}

type ApiConfig struct {
	Cors string
	Port int
}

type ImportConfig struct {
	DataFolderPath string
}

func InitConfig() {
	currentDir, _ := os.Getwd()
	viper.AddConfigPath(currentDir)

	viper.SetConfigName(AppConfigDefaultName)
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	viper.SetDefault(KeyAppLogLevel, DefaultAppLogLevel)

	viper.SetDefault(KeyApiCors, DefaultApiCors)
	viper.SetDefault(KeyApiPort, DefaultApiPort)

	viper.SetDefault(KeyImportDataFolderPath, DefaultImportDataFolderPath)

	if err := viper.ReadInConfig(); err == nil {
		slog.Info("Using config file: ", viper.ConfigFileUsed())
	}

	level := ParseLogLevel(viper.GetString(KeyAppLogLevel))
	slog.SetLogLoggerLevel(level)
}

func GetConfig() *AppConfig {
	return &AppConfig{
		LogLevel: viper.GetString(KeyAppLogLevel),
		Api: ApiConfig{
			Cors: viper.GetString(KeyApiCors),
			Port: viper.GetInt(KeyApiPort),
		},
		Import: ImportConfig{
			DataFolderPath: viper.GetString(KeyImportDataFolderPath),
		},
	}
}

func ParseLogLevel(levelStr string) slog.Level {
	var level slog.Level
	switch levelStr {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	return level
}
