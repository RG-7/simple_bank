package util

import "github.com/spf13/viper"

// Configuration stores all configurations of the application.
// the values are read by viper from a config file or environment varibales.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`
}

// LoadCOnfig reads configuratios from file or environment varibales.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
