package config

import "github.com/spf13/viper"

type Config struct {
	WebPort  int
	Channel  string
	Server   string
	Password string
	Nickname string
	DBPath   string
}

func GetConfig() *Config {
	viper.SetDefault("SERVER", "")
	viper.SetDefault("PASSWORD", "")
	viper.SetDefault("NICK", "")
	viper.SetDefault("WEB_PORT", 8000)
	viper.SetDefault("CHANNEL", "")
	viper.SetDefault("DB_PATH", "./data/db")
	viper.AutomaticEnv()
	return &Config{
		WebPort:  viper.GetInt("WEB_PORT"),
		Channel:  viper.GetString("CHANNEL"),
		Server:   viper.GetString("SERVER"),
		Password: viper.GetString("PASSWORD"),
		Nickname: viper.GetString("NICK"),
		DBPath:   viper.GetString("DB_PATH"),
	}
}
