package main

import (
	"fmt"
	"github.com/spf13/viper"
	ts3 "go-ts3-http"
)

func createViperConfig() {
	viper.SetDefault("apiKey", "ChangeMe")
	viper.SetDefault("baseUrl", "http://localhost:10080")
	viper.SetDefault("debug", false)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	_ = viper.SafeWriteConfigAs("config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

var client ts3.TeamspeakHttpClient

func init() {
	createViperConfig()
	config := ts3.NewConfig(
		viper.GetString("baseUrl"),
		viper.GetString("apiKey"),
	)

	client = ts3.NewClient(config)
}

func main() {
	clientList, err := client.ClientList(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range *clientList {
		fmt.Printf("%+v\n", user)
		fmt.Println(user.IsBot())
	}
}
