package configs

import "github.com/spf13/viper"

type Config struct {
	AWSSecretKey string `mapstructure:"AWS_SECRET_KEY"`
	AWSPublicId  string `mapstructure:"AWS_PUBLIC_ID"`
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
