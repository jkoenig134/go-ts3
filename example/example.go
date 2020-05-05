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

var client ts3.Client

func init() {
	createViperConfig()
	config := ts3.NewConfig(
		viper.GetString("baseUrl"),
		viper.GetString("apiKey"),
		viper.GetBool("debug"),
	)

	client = ts3.NewClient(config)
}

func main() {
	whoami, _ := client.Whoami()
	fmt.Printf("%+v\n", whoami)

	clientList, _ := client.ChannelList(1)
	for _, user := range clientList.Channels {
		fmt.Printf("%+v\n", user)
	}
}
