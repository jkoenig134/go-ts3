package go_ts3_http

import (
	"github.com/jkoenig134/go-ts3-http/rawevent"
	"github.com/spf13/viper"
)

type TeamspeakEvent string

//noinspection GoUnusedConst
const (
	TextMessage               = "notifytextmessage"
	ClientEnterView           = "notifycliententerview"
	ClientLeftView            = "notifyclientleftview"
	ServerEdited              = "notifyserveredited"
	ChannelEdited             = "notifychanneledited"
	ChannelDescriptionChanged = "notifychanneldescriptionchanged"
	ClientMoved               = "notifyclientmoved"
	ChannelCreated            = "notifychannelcreated"
	ChannelDeleted            = "notifychanneldeleted"
	ChannelMoved              = "notifychannelmoved"
	ChannelPasswordChanged    = "notifychannelpasswordchanged"
	TokenUsed                 = "notifytokenused"
)

func (c *TeamspeakHttpClient) StartEventClient() error {
	client, err := rawevent.Start(
		viper.GetString("eventHost"),
		viper.GetString("eventUser"),
		viper.GetString("eventPassword"),
		c.eventBus,
	)
	if err != nil {
		return err
	}

	client.SwitchServer(c.serverID)

	c.eventClient = client

	return nil
}

func (c *TeamspeakHttpClient) SubscribeEvent(event TeamspeakEvent, fn interface{}) error {
	return c.eventBus.Subscribe(string(event), fn)
}

func (c *TeamspeakHttpClient) UnsubscribeEvent(event TeamspeakEvent, handler interface{}) error {
	return c.eventBus.Unsubscribe(string(event), handler)
}
