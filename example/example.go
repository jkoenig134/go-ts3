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
	channelInfo()
}

func channelInfo() {
	info, err := client.ChannelInfo(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", info)
}

//noinspection GoUnusedFunction
func ban() {
	banlist, err := client.BanList(ts3.BanListRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", banlist)
}

//noinspection GoUnusedFunction
func complain() {
	complains, err := client.ComplainList()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", complains)
}

//noinspection GoUnusedFunction
func querylogin() {
	login, err := client.QueryloginAddGlobal("thethinggoesskrr")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", login)

	queryloginlist()

	err = client.QueryloginDeleteGlobal(6)
	if err != nil {
		fmt.Println(err)
		return
	}

	queryloginlist()
}

func queryloginlist() {
	logins, err := client.QueryloginListGlobal(ts3.QueryloginListRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", logins)
}

//noinspection GoUnusedFunction
func newToken() {
	token, err := client.TokenAdd(ts3.NewGroupToken(6, "added by go-ts3-http"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", token)

	token, err = client.TokenAdd(ts3.NewChannelToken(5, 1, "added by go-ts3-http"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", token)

}

//noinspection GoUnusedFunction
func token() {
	tokens, err := client.TokenList()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", tokens)
}

//noinspection GoUnusedFunction
func apikeyadd() {
	newKey, err := client.ApiKeyAdd(ts3.ApiKeyAddRequest{
		Scope: ts3.MANAGE,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", newKey)
}

//noinspection GoUnusedFunction
func apikeylist() {
	client.SetServerID(0)
	list, err := client.ApiKeyList(ts3.ApiKeyListRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", list)
}

//noinspection GoUnusedFunction
func messageList() {
	asd, err := client.MessageList()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(asd)
}

//noinspection GoUnusedFunction
func sendMessage() {
	err := client.SendChannelMessage("[b]asd asdanflk[/b] \n jnueoindaeo")
	if err != nil {
		fmt.Println(err)
	}

	err = client.SendClientMessage(16, "[b]asd asdanflk[/b] \n jnueoindaeo")
	if err != nil {
		fmt.Println(err)
	}
}

//noinspection GoUnusedFunction
func gm() {
	err := client.GlobalMessage("[b]asd asdanflk[/b] \n jnueoindaeo")
	if err != nil {
		fmt.Println(err)
		return
	}
}

//noinspection GoUnusedFunction
func binding() {
	bindings, err := client.BindingList(ts3.BindingListRequest{
		Subsystem: ts3.VOICE,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, binding := range *bindings {
		fmt.Printf("%+v\n", binding)
	}
}

//noinspection GoUnusedFunction
func hostInfo() {
	hostInfo, err := client.HostInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *hostInfo)
}

//noinspection GoUnusedFunction
func channelList() {
	channelList, err := client.ChannelList()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, channel := range *channelList {
		fmt.Printf("%+v\n", channel)
	}
}

//noinspection GoUnusedFunction
func clientList() {
	clientList, err := client.ClientList()
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

//noinspection GoUnusedFunction
func encoding() {
	filter := Filter{
		Limit:  1,
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
