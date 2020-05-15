package go_ts3

type EventReasonId int

//noinspection GoUnusedConst
const (
	EventReasonUserMovedIndependent EventReasonId = 0
	EventReasonUserWasMoved         EventReasonId = 1
	EventReasonTimeout              EventReasonId = 3
	EventReasonChannelKick          EventReasonId = 4
	EventReasonServerKick           EventReasonId = 5
	EventReasonServerBan            EventReasonId = 6
	EventReasonServerLeft           EventReasonId = 8
	EventReasonServerChannelEdited  EventReasonId = 10
	EventReasonServerShutdown       EventReasonId = 11
)

// notifycliententerview
type ClientEnterViewEvent struct {
	CfId                                 int           `schema:"cfid"`
	ChannelTargetId                      int           `schema:"ctid"`
	ReasonId                             EventReasonId `schema:"reasonid"`
	ClientId                             int           `schema:"clid"`
	ClientUniqueIdentifier               string        `schema:"client_unique_identifier"`
	ClientNickname                       string        `schema:"client_nickname"`
	ClientInputMuted                     string        `schema:"client_input_muted"`
	ClientOutputMuted                    string        `schema:"client_output_muted"`
	ClientOutputonlyMuted                string        `schema:"client_outputonly_muted"`
	ClientInputHardware                  string        `schema:"client_input_hardware"`
	ClientOutputHardware                 string        `schema:"client_output_hardware"`
	ClientMetaData                       string        `schema:"client_meta_data"`
	ClientIsRecording                    string        `schema:"client_is_recording"`
	ClientDatabaseId                     int           `schema:"client_database_id"`
	ClientChannelGroupId                 int           `schema:"client_channel_group_id"`
	ClientServergroups                   []int         `schema:"client_servergroups"`
	ClientAway                           string        `schema:"client_away"`
	ClientAwayMessage                    string        `schema:"client_away_message"`
	ClientType                           int           `schema:"client_type"`
	ClientFlagAvatar                     string        `schema:"client_flag_avatar"`
	ClientTalkPower                      int           `schema:"client_talk_power"`
	ClientTalkRequest                    string        `schema:"client_talk_request"`
	ClientTalkRequestMsg                 string        `schema:"client_talk_request_msg"`
	ClientDescription                    string        `schema:"client_description"`
	ClientIsTalker                       string        `schema:"client_is_talker"`
	ClientIsPrioritySpeaker              string        `schema:"client_is_priority_speaker"`
	ClientUnreadMessages                 []string      `schema:"client_unread_messages"`
	ClientNicknamePhonetic               string        `schema:"client_nickname_phonetic"`
	ClientNeededServerqueryViewPower     int           `schema:"client_needed_serverquery_view_power"`
	ClientIconId                         string        `schema:"client_icon_id"`
	ClientIsChannelCommander             string        `schema:"client_is_channel_commander"`
	ClientCountry                        string        `schema:"client_country"`
	ClientChannelGroupInheritedChannelId string        `schema:"client_channel_group_inherited_channel_id"`
	ClientBadges                         []string      `schema:"client_badges"`
	ClientIntegrations                   []string      `schema:"client_integrations"`
	ClientMyteamspeakId                  string        `schema:"client_myteamspeak_id"`
}

func (e *ClientEnterViewEvent) IsBot() bool {
	return e.ClientType == 1
}

// notifyclientleftview
type ClientLeftViewEvent struct {
	CfId            int           `schema:"cfid"`
	ChannelTargetId int           `schema:"ctid"`
	ReasonId        EventReasonId `schema:"reasonid"`
	InvokerId       int           `schema:"invokerid"`
	InvokerName     string        `schema:"invokername"`
	InvokerUid      string        `schema:"invokeruid"`
	ReasonMessage   string        `schema:"reasonmsg"`
	BanTime         int           `schema:"bantime"`
	ClientId        int           `schema:"clid"`
}

// notifyserveredited
type ServerEditedEvent struct {
	ReasonId                                    EventReasonId `schema:"reasonid"`
	InvokerId                                   int           `schema:"invokerid"`
	InvokerName                                 string        `schema:"invokername"`
	InvokerUid                                  string        `schema:"invokeruid"`
	VirtualserverName                           string        `schema:"virtualserver_name"`
	VirtualserverCodecEncryptionMode            string        `schema:"virtualserver_codec_encryption_mode"`
	VirtualserverDefaultServerGroup             int           `schema:"virtualserver_default_server_group"`
	VirtualserverDefaultChannelGroup            int           `schema:"virtualserver_default_channel_group"`
	VirtualserverHostbannerUrl                  string        `schema:"virtualserver_hostbanner_url"`
	VirtualserverHostbannerGfxUrl               string        `schema:"virtualserver_hostbanner_gfx_url"`
	VirtualserverHostbannerGfxInterval          int           `schema:"virtualserver_hostbanner_gfx_interval"`
	VirtualserverPrioritySpeakerDimmModificator string        `schema:"virtualserver_priority_speaker_dimm_modificator"`
	VirtualserverHostbuttonTooltip              string        `schema:"virtualserver_hostbutton_tooltip"`
	VirtualserverHostbuttonUrl                  string        `schema:"virtualserver_hostbutton_url"`
	VirtualserverHostbuttonGfxUrl               string        `schema:"virtualserver_hostbutton_gfx_url"`
	VirtualserverNamePhonetic                   string        `schema:"virtualserver_name_phonetic"`
	VirtualserverIconId                         string        `schema:"virtualserver_icon_id"`
	VirtualserverHostbannerMode                 string        `schema:"virtualserver_hostbanner_mode"`
	VirtualserverChannelTempDeleteDelayDefault  int           `schema:"virtualserver_channel_temp_delete_delay_default"`
}

// notifychanneldescriptionchanged
type ChannelDescriptionChangedEvent struct {
	ChannelId int `schema:"cid"`
}

// notifychannelpasswordchanged
type ChannelPasswordChangedEvent struct {
	ChannelId int `schema:"cid"`
}

// notifychannelmoved
type ChannelMovedEvent struct {
	ChannelId       int           `schema:"cid"`
	ChannelParentId int           `schema:"cpid"`
	Order           int           `schema:"order"`
	ReasonId        EventReasonId `schema:"reasonid"`
	InvokerId       int           `schema:"invokerid"`
	InvokerName     string        `schema:"invokername"`
	InvokerUid      string        `schema:"invokeruid"`
}

// notifychanneledited
type ChannelEditedEvent struct {
	ChannelId                            int           `schema:"cid"`
	ReasonId                             EventReasonId `schema:"reasonid"`
	InvokerId                            int           `schema:"invokerid"`
	InvokerName                          string        `schema:"invokername"`
	InvokerUid                           string        `schema:"invokeruid"`
	ChannelName                          string        `schema:"channel_name"`
	ChannelTopic                         string        `schema:"channel_topic"`
	ChannelCodec                         string        `schema:"channel_codec"`
	ChannelCodecQuality                  string        `schema:"channel_codec_quality"`
	ChannelMaxclients                    int           `schema:"channel_maxclients"`
	ChannelMaxFamilyClients              int           `schema:"channel_maxfamilyclients"`
	ChannelOrder                         int           `schema:"channel_order"`
	ChannelFlagPermanent                 string        `schema:"channel_flag_permanent"`
	ChannelFlagSemiPermanent             string        `schema:"channel_flag_semi_permanent"`
	ChannelFlagDefault                   string        `schema:"channel_flag_default"`
	ChannelFlagPassword                  string        `schema:"channel_flag_password"`
	ChannelCodecLatencyFactor            int           `schema:"channel_codec_latency_factor"`
	ChannelCodecIsUnencrypted            string        `schema:"channel_codec_is_unencrypted"`
	ChannelDeleteDelay                   int           `schema:"channel_delete_delay"`
	ChannelFlagMaxclientsUnlimited       string        `schema:"channel_flag_maxclients_unlimited"`
	ChannelFlagMaxfamilyclientsUnlimited string        `schema:"channel_flag_maxfamilyclients_unlimited"`
	ChannelFlagMaxfamilyclientsInherited string        `schema:"channel_flag_maxfamilyclients_inherited"`
	ChannelNeededTalkPower               int           `schema:"channel_needed_talk_power"`
	ChannelNamePhonetic                  string        `schema:"channel_name_phonetic"`
	ChannelIconId                        string        `schema:"channel_icon_id"`
}

// notifychannelcreated
type ChannelCreatedEvent struct {
	ChannelId                            int    `schema:"cid"`
	ChannelParentId                      int    `schema:"cpid"`
	ChannelName                          string `schema:"channel_name"`
	ChannelTopic                         string `schema:"channel_topic"`
	ChannelCodec                         string `schema:"channel_codec"`
	ChannelCodecQuality                  string `schema:"channel_codec_quality"`
	ChannelMaxClients                    int    `schema:"channel_maxclients"`
	ChannelMaxFamilyClients              int    `schema:"channel_maxfamilyclients"`
	ChannelOrder                         int    `schema:"channel_order"`
	ChannelFlagPermanent                 string `schema:"channel_flag_permanent"`
	ChannelFlagSemiPermanent             string `schema:"channel_flag_semi_permanent"`
	ChannelFlagDefault                   string `schema:"channel_flag_default"`
	ChannelFlagPassword                  string `schema:"channel_flag_password"`
	ChannelCodecLatencyFactor            int    `schema:"channel_codec_latency_factor"`
	ChannelCodecIsUnencrypted            string `schema:"channel_codec_is_unencrypted"`
	ChannelDeleteDelay                   int    `schema:"channel_delete_delay"`
	ChannelFlagMaxclientsUnlimited       string `schema:"channel_flag_maxclients_unlimited"`
	ChannelFlagMaxFamilyClientsUnlimited string `schema:"channel_flag_maxfamilyclients_unlimited"`
	ChannelFlagMaxFamilyClientsInherited string `schema:"channel_flag_maxfamilyclients_inherited"`
	ChannelNeededTalkPower               int    `schema:"channel_needed_talk_power"`
	ChannelNamePhonetic                  string `schema:"channel_name_phonetic"`
	ChannelIconId                        string `schema:"channel_icon_id"`
	InvokerId                            int    `schema:"invokerid"`
	InvokerName                          string `schema:"invokername"`
	InvokerUid                           string `schema:"invokeruid"`
}

// notifychanneldeleted
type ChannelDeletedEvent struct {
	InvokerId   int    `schema:"invokerid"`
	InvokerName string `schema:"invokername"`
	InvokerUid  string `schema:"invokeruid"`
	ChannelId   int    `schema:"cid"`
}

//notifyclientmoved
type ClientMovedEvent struct {
	TargetChannelId int           `schema:"ctid"`
	ReasonId        EventReasonId `schema:"reasonid"`
	InvokerId       int           `schema:"invokerid"`
	InvokerName     string        `schema:"invokername"`
	InvokerUid      string        `schema:"invokeruid"`
	ClientId        []int         `schema:"clid"`
}

// notifytextmessage
type TextMessageEvent struct {
	TargetMode  int    `schema:"targetmode"`
	Message     string `schema:"msg"`
	Target      string `schema:"target"`
	InvokerId   int    `schema:"invokerid"`
	InvokerName string `schema:"invokername"`
	InvokerUid  string `schema:"invokeruid"`
}

func (e *TextMessageEvent) IsGm() bool {
	return e.InvokerId == 0 && e.InvokerName == "Server"
}

func (e *TextMessageEvent) IsPrivate() bool {
	return e.TargetMode == 1
}

func (e *TextMessageEvent) IsChannel() bool {
	return e.TargetMode == 2
}

func (e *TextMessageEvent) IsServer() bool {
	return e.TargetMode == 3
}

// notifytokenused
type TokenUsedEvent struct {
	ClientId       int      `schema:"clid"`
	ClientDbId     int      `schema:"cldbid"`
	ClientUid      string   `schema:"cluid"`
	Token          string   `schema:"token"`
	TokenCustomSet []string `schema:"tokencustomset"`
	Token1         string   `schema:"token1"`
	Token2         string   `schema:"token2"`
}
