package configs

import "github.com/spf13/viper"

type Config struct {
	//  API URl
	IpURI string `mapstructure:"IP_URI"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
