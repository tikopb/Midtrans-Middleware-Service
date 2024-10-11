package main_module

import (
	"github.com/spf13/viper"
)

type Repository interface {
	//CreatePayment
	GetEnvVariabel(envName string) string
}

type masterMain struct {
}

func GetRepository() masterMain {
	return masterMain{}
}

func (m *masterMain) GetEnvVariabel(envName string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic("env of " + envName + "not found")
	}

	value := viper.GetString(envName)
	return value
}
