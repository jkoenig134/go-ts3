package main

import (
	"fmt"
	ts3 "github.com/jkoenig134/go-ts3"
	"github.com/jkoenig134/schema"
	"github.com/spf13/viper"
	"net/url"
)

func createViperConfig() {
	viper.SetDefault("apiKey", "ChangeMe")
	viper.SetDefault("baseUrl", "http://localhost:10080")
	viper.SetDefault("debug", false)
	viper.SetDefault("eventHost", "localhost:10011")
	viper.SetDefault("eventUser", "serveradmin")
	viper.SetDefault("eventPassword", "ChangeMe")
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
	channelGroupList()
}

func channelGroupList() {
	list, err := client.ServerGroupList()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, cg := range *list {
		fmt.Printf("%s %d\n", cg.Name, cg.Type)
	}
}

//noinspection GoUnusedFunction
func customSetTokens() {
	customSet := map[string]string{
		"forumUser":  "asd bla",
		"teamMember": "false",
	}

	channelToken := ts3.NewCustomSetChannelToken(1, 1, "Test", customSet)
	fmt.Printf("%+v", channelToken)

	groupToken := ts3.NewCustomSetGroupToken(1, "Test", customSet)
	fmt.Printf("%+v", groupToken)
}

//noinspection GoUnusedFunction
func channelPerm() {
	list, err := client.ChannelList()
	if err != nil {
		fmt.Println(err)
		return
	}

	channel := (*list)[0]

	printIfNoError(client.ChannelPermissionList(channel.ChannelId))
	printIfNoError(client.ChannelStringPermissionList(channel.ChannelId))
}

//noinspection GoUnusedFunction
func custom() {
	err := client.CustomSet(34, "IGName", "Fancy Ingame Name")
	if err != nil {
		fmt.Println(err)
		return
	}

	list, err := client.CustomInfo(34)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", list)

	err = client.CustomDelete(34, "IGName")
}

//noinspection GoUnusedFunction
func permissionOverview() {
	overview, err := client.PermissionOverview(1, 34)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", overview)
}

//noinspection GoUnusedFunction
func permissions() {
	list, err := client.PermissionList()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", list)
}

//noinspection GoUnusedFunction
func serverGroupsByClient() {
	groupList, err := client.ServerGroupsByClientId(34)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", groupList)
}

//noinspection GoUnusedFunction
func serverGroups() {
	groupList, err := client.ServerGroupList()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", groupList)
}

//noinspection GoUnusedFunction
func clientPerms() {
	info, err := client.ClientStringPermissionList(34)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", info)
}

//noinspection GoUnusedFunction
func clientInfo() {
	list, _ := client.ClientList()

	info, _ := client.ClientInfo((*list)[0].ClientId)
	fmt.Printf("%+v", info)
}

//noinspection GoUnusedFunction
func event() {
	err := client.SubscribeEvent(ts3.NotifyClientMoved, func(v *ts3.ClientMovedEvent) {
		fmt.Printf("%+v\n", v)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.SubscribeEvent(ts3.NotifyClientEnterView, func(v *ts3.ClientEnterViewEvent) {
		fmt.Printf("%+v\n", v)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.StartEventClient(
		viper.GetString("eventHost"),
		viper.GetString("eventUser"),
		viper.GetString("eventPassword"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = fmt.Scanf("Type to exit")
}

//noinspection GoUnusedFunction
func printIfNoError(v interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", v)
}

//noinspection GoUnusedFunction
func server() {
	serverList()

	err := client.ServerEdit(ts3.ServerEditRequest{
		VirtualserverName: "Test789",
	})
	if err != nil {
		return
	}

	serverList()

	err = client.ServerEdit(ts3.ServerEditRequest{
		VirtualserverName: "jkoenig.dev",
	})
	if err != nil {
		return
	}

	serverList()
}

func serverList() {
	serverList, err := client.ServerList()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, server := range *serverList {
		fmt.Println(server.VirtualserverName)
	}
}

//noinspection GoUnusedFunction
func serverSnapshot() {
	password := "asd"
	snap, err := client.ServerSnapshotCreate(ts3.ServerSnapshotCreateRequest{Password: password})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", snap)

	deployments, err := client.ServerSnapshotDeploy(snap.AsDeployRequest(&password, true, true))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(deployments)
}

//noinspection GoUnusedFunction
func channelfind() {
	channels, err := client.ChannelFind("ein")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", channels)
}

//noinspection GoUnusedFunction
func instanceinfo() {
	info, err := client.InstanceInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", info)
}

//noinspection GoUnusedFunction
func logview() {
	logs, err := client.LogViewInstance(1, true, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", logs)
}

//noinspection GoUnusedFunction
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
		Scope: ts3.ApiKeyScopeManage,
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
		Subsystem: ts3.SubsystemVoice,
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
