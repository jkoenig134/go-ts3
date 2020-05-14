package go_ts3

// channelcreate `manage_scope, write_scope`
type ChannelCreateRequest struct {
	ChannelName                          string `schema:"channel_name,required"`
	ChannelTopic                         string `schema:"channel_topic,omitempty"`
	ChannelDescription                   string `schema:"channel_description,omitempty"`
	ChannelPassword                      string `schema:"channel_password,omitempty"`
	ChannelCodec                         string `schema:"channel_codec,omitempty"`
	ChannelCodecQuality                  string `schema:"channel_codec_quality,omitempty"`
	ChannelMaxclients                    string `schema:"channel_maxclients,omitempty"`
	ChannelMaxfamilyclients              string `schema:"channel_maxfamilyclients,omitempty"`
	ChannelOrder                         string `schema:"channel_order,omitempty"`
	ChannelFlagPermanent                 string `schema:"channel_flag_permanent,omitempty"`
	ChannelFlagSemiPermanent             string `schema:"channel_flag_semi_permanent,omitempty"`
	ChannelFlagTemporary                 string `schema:"channel_flag_temporary,omitempty"`
	ChannelFlagDefault                   string `schema:"channel_flag_default,omitempty"`
	ChannelFlagMaxclientsUnlimited       string `schema:"channel_flag_maxclients_unlimited,omitempty"`
	ChannelFlagMaxfamilyclientsUnlimited string `schema:"channel_flag_maxfamilyclients_unlimited,omitempty"`
	ChannelFlagMaxfamilyclientsInherited string `schema:"channel_flag_maxfamilyclients_inherited,omitempty"`
	ChannelNeededTalkPower               string `schema:"channel_needed_talk_power,omitempty"`
	ChannelNamePhonetic                  string `schema:"channel_name_phonetic,omitempty"`
	ChannelIconID                        string `schema:"channel_icon_id,omitempty"`
	ChannelCodecIsUnencrypted            string `schema:"channel_codec_is_unencrypted,omitempty"`
	ChannelParentId                      string `schema:"cpid,omitempty"`
}

type channelCreateResponse struct {
	ChannelId int `json:"cid,string"`
}

func (c *TeamspeakHttpClient) ChannelCreate(request ChannelCreateRequest) (*int, error) {
	var responses []channelCreateResponse

	err := c.requestWithParams("channelcreate", request, &responses)
	if err != nil {
		return nil, err
	}

	return &responses[0].ChannelId, nil
}

// channeldelete `manage_scope, write_scope`
type channelDeleteRequest struct {
	ChannelId int `schema:"cid"`
	Force     int `schema:"force"`
}

func (c *TeamspeakHttpClient) ChannelDelete(channelId int, force bool) error {
	return c.requestWithParams(
		"channeldelete",
		channelDeleteRequest{
			ChannelId: channelId,
			Force:     boolToInt(force),
		},
		nil,
	)
}

// channeledit `manage_scope, write_scope`
type ChannelEditRequest struct {
	ChannelId                            int    `schema:"cid,required"`
	ChannelName                          string `schema:"channel_name,omitempty"`
	ChannelTopic                         string `schema:"channel_topic,omitempty"`
	ChannelDescription                   string `schema:"channel_description,omitempty"`
	ChannelPassword                      string `schema:"channel_password,omitempty"`
	ChannelCodec                         string `schema:"channel_codec,omitempty"`
	ChannelCodecQuality                  string `schema:"channel_codec_quality,omitempty"`
	ChannelMaxclients                    string `schema:"channel_maxclients,omitempty"`
	ChannelMaxfamilyclients              string `schema:"channel_maxfamilyclients,omitempty"`
	ChannelOrder                         string `schema:"channel_order,omitempty"`
	ChannelFlagPermanent                 string `schema:"channel_flag_permanent,omitempty"`
	ChannelFlagSemiPermanent             string `schema:"channel_flag_semi_permanent,omitempty"`
	ChannelFlagTemporary                 string `schema:"channel_flag_temporary,omitempty"`
	ChannelFlagDefault                   string `schema:"channel_flag_default,omitempty"`
	ChannelFlagMaxclientsUnlimited       string `schema:"channel_flag_maxclients_unlimited,omitempty"`
	ChannelFlagMaxfamilyclientsUnlimited string `schema:"channel_flag_maxfamilyclients_unlimited,omitempty"`
	ChannelFlagMaxfamilyclientsInherited string `schema:"channel_flag_maxfamilyclients_inherited,omitempty"`
	ChannelNeededTalkPower               string `schema:"channel_needed_talk_power,omitempty"`
	ChannelNamePhonetic                  string `schema:"channel_name_phonetic,omitempty"`
	ChannelIconID                        string `schema:"channel_icon_id,omitempty"`
	ChannelCodecIsUnencrypted            string `schema:"channel_codec_is_unencrypted,omitempty"`
	ChannelParentId                      string `schema:"cpid,omitempty"`
}

func (c *TeamspeakHttpClient) ChannelEdit(request ChannelEditRequest) error {
	return c.requestWithParams("channeledit", request, nil)
}

// channelfind `manage_scope, write_scope, read_scope`
type channelFindRequest struct {
	Pattern string `schema:"pattern"`
}

type ChannelFindResponse struct {
	ChannelName string `json:"channel_name"`
	ChannelId   int    `json:"cid,string"`
}

func (c *TeamspeakHttpClient) ChannelFind(pattern string) (*[]ChannelFindResponse, error) {
	var channels []ChannelFindResponse

	err := c.requestWithParams(
		"channelfind",
		channelFindRequest{Pattern: pattern},
		&channels,
	)
	if err != nil {
		return nil, err
	}

	return &channels, nil
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
