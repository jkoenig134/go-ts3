package go_ts3

import (
	"time"
)

type TeamspeakEvent string

//noinspection GoUnusedConst
const (
	NotifyTextMessage               TeamspeakEvent = "notifytextmessage"
	NotifyClientEnterView           TeamspeakEvent = "notifycliententerview"
	NotifyClientLeftView            TeamspeakEvent = "notifyclientleftview"
	NotifyServerEdited              TeamspeakEvent = "notifyserveredited"
	NotifyChannelEdited             TeamspeakEvent = "notifychanneledited"
	NotifyChannelDescriptionChanged TeamspeakEvent = "notifychanneldescriptionchanged"
	NotifyClientMoved               TeamspeakEvent = "notifyclientmoved"
	NotifyChannelCreated            TeamspeakEvent = "notifychannelcreated"
	NotifyChannelDeleted            TeamspeakEvent = "notifychanneldeleted"
	NotifyChannelMoved              TeamspeakEvent = "notifychannelmoved"
	NotifyChannelPasswordChanged    TeamspeakEvent = "notifychannelpasswordchanged"
	NotifyTokenUsed                 TeamspeakEvent = "notifytokenused"
)

func (c *TeamspeakHttpClient) StartEventClient(host, user, password string) error {
	client, err := newEventClient(host, user, password, c.eventBus)
	if err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	_ = client.SwitchServer(c.serverID)

	c.eventClient = client

	return nil
}

func (c *TeamspeakHttpClient) SubscribeEvent(event TeamspeakEvent, fn interface{}) error {
	return c.eventBus.Subscribe(string(event), fn)
}

func (c *TeamspeakHttpClient) UnsubscribeEvent(event TeamspeakEvent, handler interface{}) error {
	return c.eventBus.Unsubscribe(string(event), handler)
}
