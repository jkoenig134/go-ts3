package go_ts3_http

// channelcreate `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ChannelCreate() error {
	//TODO
	return nil
}

// channeldelete `manage_scope, write_scope`
type channelDeleteRequest struct {
	ChannelId int `schema:"cid"`
	Force     int `schema:"force"`
}

func (c *TeamspeakHttpClient) ChannelDelete(channelId int, force bool) error {
	forceInt := 0
	if force {
		forceInt = 1
	}

	return c.requestWithParams(
		"channeldelete",
		channelDeleteRequest{
			ChannelId: channelId,
			Force:     forceInt,
		},
		nil,
	)
}

// channeledit `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ChannelEdit() error {
	//TODO
	return nil
}

// channelfind `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ChannelFind() error {
	//TODO
	return nil
}

// channelinfo `manage_scope, write_scope, read_scope`
type channelInfoRequest struct {
	ChannelId int `schema:"cid,required"`
}

type ChannelInfo struct {
	ChannelBannerGfxURL                  string `json:"channel_banner_gfx_url"`
	ChannelBannerMode                    int    `json:"channel_banner_mode,string"`
	ChannelCodec                         int    `json:"channel_codec,string"`
	ChannelCodecIsUnencrypted            int    `json:"channel_codec_is_unencrypted,string"`
	ChannelCodecLatencyFactor            int    `json:"channel_codec_latency_factor,string"`
	ChannelCodecQuality                  int    `json:"channel_codec_quality,string"`
	ChannelDeleteDelay                   int    `json:"channel_delete_delay,string"`
	ChannelDescription                   string `json:"channel_description"`
	ChannelFilepath                      string `json:"channel_filepath"`
	ChannelFlagDefault                   int    `json:"channel_flag_default,string"`
	ChannelFlagMaxclientsUnlimited       int    `json:"channel_flag_maxclients_unlimited,string"`
	ChannelFlagMaxFamilyClientsInherited int    `json:"channel_flag_maxfamilyclients_inherited,string"`
	ChannelFlagMaxFamilyClientsUnlimited int    `json:"channel_flag_maxfamilyclients_unlimited,string"`
	ChannelFlagPassword                  int    `json:"channel_flag_password,string"`
	ChannelFlagPermanent                 int    `json:"channel_flag_permanent,string"`
	ChannelFlagSemiPermanent             int    `json:"channel_flag_semi_permanent,string"`
	ChannelForcedSilence                 int    `json:"channel_forced_silence,string"`
	ChannelIconID                        int    `json:"channel_icon_id,string"`
	ChannelMaxclients                    int    `json:"channel_maxclients,string"`
	ChannelMaxfamilyclients              int    `json:"channel_maxfamilyclients,string"`
	ChannelName                          string `json:"channel_name"`
	ChannelNamePhonetic                  string `json:"channel_name_phonetic"`
	ChannelNeededTalkPower               int    `json:"channel_needed_talk_power,string"`
	ChannelOrder                         int    `json:"channel_order,string"`
	ChannelPassword                      string `json:"channel_password"`
	ChannelSecuritySalt                  string `json:"channel_security_salt"`
	ChannelTopic                         string `json:"channel_topic"`
	ChannelUniqueIdentifier              string `json:"channel_unique_identifier"`
	Pid                                  int    `json:"pid,string"`
	SecondsEmpty                         int    `json:"seconds_empty,string"`
}

func (c *TeamspeakHttpClient) ChannelInfo(channelId int) (*ChannelInfo, error) {
	var channels []ChannelInfo

	err := c.requestWithParams(
		"channelinfo",
		channelInfoRequest{ChannelId: channelId},
		&channels,
	)

	if err != nil {
		return nil, err
	}

	return &channels[0], nil
}

// channellist `manage_scope, write_scope, read_scope`
type Channel struct {
	ChannelName                 string `json:"channel_name"`
	ChannelNeededSubscribePower int    `json:"channel_needed_subscribe_power,string"`
	ChannelOrder                int    `json:"channel_order,string"`
	ChannelId                   int    `json:"cid,string"`
	PID                         int    `json:"pid,string"`
	TotalClients                int    `json:"total_clients,string"`
}

func (c *TeamspeakHttpClient) ChannelList() (*[]Channel, error) {
	var channels []Channel

	err := c.request("channellist", &channels)
	if err != nil {
		return nil, err
	}

	return &channels, nil
}

// channelmove `manage_scope, write_scope`
type ChannelMoveRequest struct {
	ChannelId       int `schema:"cid,required"`
	ChannelParentId int `schema:"cpid,required"`
	Order           int `schema:"order,omitempty"`
}

func (c *TeamspeakHttpClient) ChannelMove(request ChannelMoveRequest) error {
	return c.requestWithParams(
		"channelmove",
		request,
		nil,
	)
}
