package env

// use viper package for environment variable
import (
	"github.com/spf13/viper"
)

func LoadEnv() (*viper.Viper, error) {
	cfg := viper.New()

	//cfg.SetConfigFile("config.env")
	cfg.AddConfigPath(".")

	cfg.SetConfigName("config")
	cfg.SetConfigType("env")

	//cfg.AutomaticEnv()

	err := cfg.ReadInConfig()
	if err != nil {

		return nil, err
	}

	return cfg, nil
}
