package main

import (
	"fmt"
	ts3 "github.com/jkoenig134/go-ts3-http"
	"github.com/jkoenig134/schema"
	"github.com/spf13/viper"
	"net/url"
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
	client = ts3.NewClient(ts3.NewConfig(
		viper.GetString("baseUrl"),
		viper.GetString("apiKey"),
	))
}

func main() {
	hostInfo()
}

func hostInfo() {
	hostInfo, err := client.HostInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *hostInfo)
}

func help() {
	help, err := client.Help("clientlist")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *help)
}

func channelList() {
	channelList, err := client.ChannelList(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, channel := range *channelList {
		fmt.Printf("%+v\n", channel)
	}
}

func clientList() {
	clientList, err := client.ClientList(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, client := range *clientList {
		fmt.Printf("%+v\n", client)
		fmt.Println(client.IsBot())
	}
}

type Filter struct {
	Offset int64  `schema:"offset,omitempty"`
	Limit  int64  `schema:"limit,required"`
	SortBy string `schema:"sortby,omitempty"`
	Asc    bool   `schema:"asc,omitempty"`

	//User specific filters
	Username  string `schema:"username,omitempty"`
	FirstName string `schema:"first_name,omitempty"`
	LastName  string `schema:"last_name,omitempty"`
	Status    string `schema:"status,omitempty"`
}

func encoding() {
	filter := Filter{
		Offset: 123,
		SortBy: "asd",
		Asc:    true,
	}

	form := url.Values{}
	err := schema.NewEncoder().Encode(filter, form)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s\n", form.Encode())
}
