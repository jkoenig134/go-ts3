package go_ts3_http

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

func (c *TeamspeakHttpClient) SubscribeEvent(event TeamspeakEvent, fn interface{}) error {
	return c.eventBus.Subscribe(string(event), fn)
}

func (c *TeamspeakHttpClient) UnsubscribeEvent(event TeamspeakEvent, handler interface{}) error {
	return c.eventBus.Unsubscribe(string(event), handler)
}
