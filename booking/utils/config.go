package utils

import "github.com/spf13/viper"

type ApiConfig struct {
	ServePort   string `mapstructure:"ServePort"`
	RunMode     string `mapstructure:"RunMode"`
	DBDriver    string `mapstructure:"DBDriver"`
	DBSource    string `mapstructure:"DBSource"`
	PageSize    int    `mapstructure:"PageSize"`
	MaxPageSize int    `mapstructure:"MaxPageSize"`
}

type WebConfig struct {
	ServePort   string `mapstructure:"ServePort"`
	RunMode     string `mapstructure:"RunMode"`
	DBDriver    string `mapstructure:"DBDriver"`
	DBSource    string `mapstructure:"DBSource"`
	PageSize    int    `mapstructure:"PageSize"`
	MaxPageSize int    `mapstructure:"MaxPageSize"`
}

func LoadConfig(path string, key string, ptr interface{}) error {
	vp := viper.New()
	vp.AddConfigPath(path)
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	return vp.UnmarshalKey(key, ptr)
}
