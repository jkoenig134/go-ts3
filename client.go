package go_ts3

import "errors"

// clientaddperm `manage_scope`
type clientAddPermissionRequest struct {
	ClientDbId      int `schema:"cldbid"`
	PermissionId    int `schema:"permid"`
	PermissionValue int `schema:"permvalue"`
	PermSkip        int `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ClientAddPermission(clientDbId, permissionId, permissionValue int, permSkip bool) error {
	return c.requestWithParams(
		"clientaddperm",
		clientAddPermissionRequest{
			ClientDbId:      clientDbId,
			PermissionId:    permissionId,
			PermissionValue: permissionValue,
			PermSkip:        boolToInt(permSkip),
		},
		nil,
	)
}

type clientAddStringPermissionRequest struct {
	ClientDbId      int    `schema:"cldbid"`
	PermissionName  string `schema:"permsid"`
	PermissionValue int    `schema:"permvalue"`
	PermSkip        int    `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ClientAddStringPermission(clientDbId int, permissionName string, permissionValue int, permSkip bool) error {
	return c.requestWithParams(
		"clientaddperm",
		clientAddStringPermissionRequest{
			ClientDbId:      clientDbId,
			PermissionName:  permissionName,
			PermissionValue: permissionValue,
			PermSkip:        boolToInt(permSkip),
		},
		nil,
	)
}

// clientdbdelete `manage_scope, write_scope`
type clientDbDeleteRequest struct {
	ClientDbId int `schema:"cldbid,required"`
}

func (c *TeamspeakHttpClient) ClientDbDelete(clientDbId int) error {
	return c.requestWithParams("clientdbdelete", clientDbDeleteRequest{ClientDbId: clientDbId}, nil)
}

// clientdbedit `manage_scope, write_scope`
type ClientDbEditRequest struct {
	ClientDbId               int    `schema:"cldbid,required"`
	ClientNickname           string `schema:"client_nickname"`
	ClientIsTalker           string `schema:"client_is_talker"`
	ClientDescription        string `schema:"client_description"`
	ClientIsChannelCommander string `schema:"client_is_channel_commander"`
	ClientIconID             string `schema:"client_icon_id"`
}

func (c *TeamspeakHttpClient) ClientDbEdit(request ClientDbEditRequest) error {
	return c.requestWithParams("clientdbedit", request, nil)
}

// clientdbfind `manage_scope, write_scope, read_scope`
type clientDbFindRequest struct {
	IsUid   bool   `schema:"-uid,omitempty"`
	Pattern string `schema:"pattern"`
}

type ClientDbFindResponse struct {
	ClientDbId []int `json:"cldbid,string"`
}

func (c *TeamspeakHttpClient) clientDbFind(request clientDbFindRequest) (*[]ClientDbFindResponse, error) {
	var clients []ClientDbFindResponse

	err := c.requestWithParams("clientdbfind", request, &clients)
	if err != nil {
		return nil, err
	}

	return &clients, err
}

func (c *TeamspeakHttpClient) ClientDbFindName(pattern string) (*[]ClientDbFindResponse, error) {
	return c.clientDbFind(clientDbFindRequest{
		IsUid:   false,
		Pattern: pattern,
	})
}

func (c *TeamspeakHttpClient) ClientDbFindUid(uid int) (*[]ClientDbFindResponse, error) {
	return c.clientDbFind(clientDbFindRequest{
		IsUid:   true,
		Pattern: string(uid),
	})
}

// clientdbinfo `manage_scope, write_scope, read_scope`
type clientDbInfoRequest struct {
	ClientDbId int `schema:"cldbid,required"`
}

type ClientDbInfo struct {
	ClientBase64HashClientUID  string `json:"client_base64HashClientUID"`
	ClientCreated              int    `json:"client_created,string"`
	ClientDatabaseID           int    `json:"client_database_id,string"`
	ClientDescription          string `json:"client_description"`
	ClientFlagAvatar           string `json:"client_flag_avatar"`
	ClientLastconnected        int    `json:"client_lastconnected,string"`
	ClientLastip               string `json:"client_lastip"`
	ClientMonthBytesDownloaded int    `json:"client_month_bytes_downloaded,string"`
	ClientMonthBytesUploaded   int    `json:"client_month_bytes_uploaded,string"`
	ClientNickname             string `json:"client_nickname"`
	ClientTotalBytesDownloaded int    `json:"client_total_bytes_downloaded,string"`
	ClientTotalBytesUploaded   int    `json:"client_total_bytes_uploaded,string"`
	ClientTotalconnections     int    `json:"client_totalconnections,string"`
	ClientUniqueIdentifier     string `json:"client_unique_identifier"`
}

func (c *TeamspeakHttpClient) ClientDbInfo(clientDbId int) (*ClientDbInfo, error) {
	var clientDbInfos []ClientDbInfo

	err := c.requestWithParams(
		"clientdbinfo",
		clientDbInfoRequest{ClientDbId: clientDbId},
		&clientDbInfos,
	)
	if err != nil {
		return nil, err
	}

	return &clientDbInfos[0], nil
}

// clientdblist `manage_scope, write_scope, read_scope`
type ClientDbListRequest struct {
	Start    int `schema:"start,omitempty"`
	Duration int `schema:"duration,omitempty"`
}

type DbClient struct {
	Cldbid                 int    `json:"cldbid,string"`
	ClientCreated          int    `json:"client_created,string"`
	ClientDescription      string `json:"client_description"`
	ClientLastConnected    int    `json:"client_lastconnected,string"`
	ClientLastIP           string `json:"client_lastip"`
	ClientLoginName        string `json:"client_login_name"`
	ClientNickname         string `json:"client_nickname"`
	ClientTotalConnections int    `json:"client_totalconnections,string"`
	ClientUniqueIdentifier string `json:"client_unique_identifier"`
}

func (c *TeamspeakHttpClient) ClientDbList(request ClientDbListRequest) (*[]DbClient, error) {
	var clients []DbClient

	err := c.requestWithParams("clientdblist", request, &clients)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

// clientdelperm `manage_scope`
type ClientDeletePermission struct {
	ClientDbId int      `schema:"cldbid,required"`
	PermId     []int    `schema:"permid"`
	PermsId    []string `schema:"permsid"`
}

func (c *TeamspeakHttpClient) ClientDeletePermission(request ClientDeletePermission) error {
	return c.requestWithParams("clientdelperm", request, nil)
}

// clientedit `manage_scope, write_scope`
type ClientEditRequest struct {
	ClientId                 int    `schema:"clid,required"`
	ClientNickname           string `schema:"client_nickname"`
	ClientIsTalker           string `schema:"client_is_talker"`
	ClientDescription        string `schema:"client_description"`
	ClientIsChannelCommander string `schema:"client_is_channel_commander"`
	ClientIconID             string `schema:"client_icon_id"`
}

func (c *TeamspeakHttpClient) ClientEdit(request ClientEditRequest) error {
	return c.requestWithParams("clientedit", request, nil)
}

// clientfind `manage_scope, write_scope, read_scope`
type clientFindRequest struct {
	Pattern string `schema:"pattern,required"`
}

type ClientFindResponse struct {
	ClientId       int    `json:"clid,string"`
	ClientNickname string `json:"client_nickname"`
}

func (c *TeamspeakHttpClient) ClientFind(pattern string) (*[]ClientFindResponse, error) {
	var clients []ClientFindResponse

	err := c.requestWithParams("clientfind", clientFindRequest{Pattern: pattern}, clients)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

// clientgetdbidfromuid `manage_scope, write_scope, read_scope`
type clientGetDbIdFromUidRequest struct {
	ClientUid string `schema:"cluid,required"`
}

type clientGetDbIdFromUidResponse struct {
	ClientDbId int `json:"cldbid,string"`
}

func (c *TeamspeakHttpClient) ClientGetDbIdFromUid(clientUid string) (*int, error) {
	var clients []clientGetDbIdFromUidResponse

	err := c.requestWithParams(
		"clientgetdbidfromuid",
		clientGetDbIdFromUidRequest{ClientUid: clientUid},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0].ClientDbId, nil
}

// clientgetids `manage_scope, write_scope, read_scope`
type clientGetIdsRequest struct {
	ClientUid string `schema:"cluid,required"`
}

type ClientGetIdsResponse struct {
	ClientUid string `json:"cluid"`
	ClientId  int    `json:"clid,string"`
	Name      string `json:"name"`
}

func (c *TeamspeakHttpClient) ClientGetIds(clientUid string) (*ClientGetIdsResponse, error) {
	var clients []ClientGetIdsResponse

	err := c.requestWithParams(
		"clientgetids",
		clientGetIdsRequest{ClientUid: clientUid},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientgetnamefromdbid `manage_scope, write_scope, read_scope`
type clientGetNameFromDbIdRequest struct {
	ClientDbId int `schema:"cldbid,required"`
}

type clientGetNameFromDbIdResponse struct {
	ClientUid  string `json:"cluid"`
	ClientDbId int    `json:"cldbid,string"`
	Name       string `json:"name"`
}

func (c *TeamspeakHttpClient) ClientGetNameFromDbId(clientDbId int) (*clientGetNameFromDbIdResponse, error) {
	var clients []clientGetNameFromDbIdResponse

	err := c.requestWithParams(
		"clientgetnamefromdbid",
		clientGetNameFromDbIdRequest{ClientDbId: clientDbId},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientgetnamefromuid `manage_scope, write_scope, read_scope`
type clientGetNameFromUidRequest struct {
	ClientUid string `schema:"cldbid,required"`
}

type clientGetNameFromUidResponse struct {
	ClientUid  string `json:"cluid"`
	ClientDbId int    `json:"cldbid,string"`
	Name       string `json:"name"`
}

func (c *TeamspeakHttpClient) ClientGetNameFromUid(clientUid string) (*clientGetNameFromUidResponse, error) {
	var clients []clientGetNameFromUidResponse

	err := c.requestWithParams(
		"clientgetnamefromdbid",
		clientGetNameFromUidRequest{ClientUid: clientUid},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientgetuidfromclid `manage_scope, write_scope, read_scope`
type clientGetUidFromClientIdRequest struct {
	ClientId int `schema:"clid,required"`
}

type clientGetUidFromClientIdResponse struct {
	ClientId  int    `json:"clid,string"`
	ClientUid string `json:"cluid"`
	NickName  string `json:"nickname"`
}

func (c *TeamspeakHttpClient) ClientGetUidFromClientId(clientId int) (*clientGetUidFromClientIdResponse, error) {
	var clients []clientGetUidFromClientIdResponse

	err := c.requestWithParams(
		"clientgetuidfromclid",
		clientGetUidFromClientIdRequest{ClientId: clientId},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientinfo `manage_scope, write_scope, read_scope`
type clientInfoRequest struct {
	ClientId int `schema:"clid"`
}

type ClientInfo struct {
	ClientId                                   int    `json:"cid,string"`
	ClientAway                                 int    `json:"client_away,string"`
	ClientAwayMessage                          string `json:"client_away_message"`
	ClientBadges                               string `json:"client_badges"`
	ClientBase64HashClientUID                  string `json:"client_base64HashClientUID"`
	ClientChannelGroupID                       int    `json:"client_channel_group_id,string"`
	ClientChannelGroupInheritedChannelID       int    `json:"client_channel_group_inherited_channel_id,string"`
	ClientCountry                              string `json:"client_country"`
	ClientCreated                              int    `json:"client_created,string"`
	ClientDatabaseID                           int    `json:"client_database_id,string"`
	ClientDefaultChannel                       string `json:"client_default_channel"`
	ClientDefaultToken                         string `json:"client_default_token"`
	ClientDescription                          string `json:"client_description"`
	ClientFlagAvatar                           string `json:"client_flag_avatar"`
	ClientIconID                               int    `json:"client_icon_id,string"`
	ClientIdleTime                             int    `json:"client_idle_time,string"`
	ClientInputHardware                        int    `json:"client_input_hardware,string"`
	ClientInputMuted                           int    `json:"client_input_muted,string"`
	ClientIntegrations                         string `json:"client_integrations"`
	ClientIsChannelCommander                   int    `json:"client_is_channel_commander,string"`
	ClientIsPrioritySpeaker                    int    `json:"client_is_priority_speaker,string"`
	ClientIsRecording                          int    `json:"client_is_recording,string"`
	ClientIsTalker                             int    `json:"client_is_talker,string"`
	ClientLastconnected                        int    `json:"client_lastconnected,string"`
	ClientLoginName                            string `json:"client_login_name"`
	ClientMetaData                             string `json:"client_meta_data"`
	ClientMonthBytesDownloaded                 int    `json:"client_month_bytes_downloaded,string"`
	ClientMonthBytesUploaded                   int    `json:"client_month_bytes_uploaded,string"`
	ClientMyteamspeakAvatar                    string `json:"client_myteamspeak_avatar"`
	ClientMyteamspeakID                        string `json:"client_myteamspeak_id"`
	ClientNeededServerqueryViewPower           int    `json:"client_needed_serverquery_view_power,string"`
	ClientNickname                             string `json:"client_nickname"`
	ClientNicknamePhonetic                     string `json:"client_nickname_phonetic"`
	ClientOutputHardware                       int    `json:"client_output_hardware,string"`
	ClientOutputMuted                          int    `json:"client_output_muted,string"`
	ClientOutputonlyMuted                      int    `json:"client_outputonly_muted,string"`
	ClientPlatform                             string `json:"client_platform"`
	ClientSecurityHash                         string `json:"client_security_hash"`
	ClientServergroups                         int    `json:"client_servergroups,string"`
	ClientSignedBadges                         string `json:"client_signed_badges"`
	ClientTalkPower                            int    `json:"client_talk_power,string"`
	ClientTalkRequest                          int    `json:"client_talk_request,string"`
	ClientTalkRequestMsg                       string `json:"client_talk_request_msg"`
	ClientTotalBytesDownloaded                 int    `json:"client_total_bytes_downloaded,string"`
	ClientTotalBytesUploaded                   int    `json:"client_total_bytes_uploaded,string"`
	ClientTotalconnections                     int    `json:"client_totalconnections,string"`
	ClientType                                 int    `json:"client_type,string"`
	ClientUniqueIdentifier                     string `json:"client_unique_identifier"`
	ClientUnreadMessages                       int    `json:"client_unread_messages,string"`
	ClientVersion                              string `json:"client_version"`
	ClientVersionSign                          string `json:"client_version_sign"`
	ConnectionBandwidthReceivedLastMinuteTotal int    `json:"connection_bandwidth_received_last_minute_total,string"`
	ConnectionBandwidthReceivedLastSecondTotal int    `json:"connection_bandwidth_received_last_second_total,string"`
	ConnectionBandwidthSentLastMinuteTotal     int    `json:"connection_bandwidth_sent_last_minute_total,string"`
	ConnectionBandwidthSentLastSecondTotal     int    `json:"connection_bandwidth_sent_last_second_total,string"`
	ConnectionBytesReceivedTotal               int    `json:"connection_bytes_received_total,string"`
	ConnectionBytesSentTotal                   int    `json:"connection_bytes_sent_total,string"`
	ConnectionClientIP                         string `json:"connection_client_ip"`
	ConnectionConnectedTime                    int    `json:"connection_connected_time,string"`
	ConnectionFiletransferBandwidthReceived    int    `json:"connection_filetransfer_bandwidth_received,string"`
	ConnectionFiletransferBandwidthSent        int    `json:"connection_filetransfer_bandwidth_sent,string"`
	ConnectionPacketsReceivedTotal             int    `json:"connection_packets_received_total,string"`
	ConnectionPacketsSentTotal                 int    `json:"connection_packets_sent_total,string"`
}

func (c *TeamspeakHttpClient) ClientInfo(clientId int) (*[]ClientInfo, error) {
	var clients []ClientInfo

	err := c.requestWithParams(
		"clientinfo",
		clientInfoRequest{ClientId: clientId},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

// clientkick `manage_scope, write_scope`
type ClientKickReason int

//noinspection GoUnusedConst
const (
	BLI ClientKickReason = 4
	BLA ClientKickReason = 5
)

type ClientKickRequest struct {
	ClientId      []int            `schema:"clid,required"`
	ReasonId      ClientKickReason `schema:"reasonid,required"`
	ReasonMessage string           `schema:"reasonmsg,omitempty"`
}

func (c *TeamspeakHttpClient) ClientKick(request ClientKickRequest) error {
	if len(request.ReasonMessage) > 40 {
		return errors.New("the message must not have more than 40 characters")
	}

	return c.requestWithParams("clientkick", request, nil)
}

// clientlist `manage_scope, write_scope, read_scope`
type Client struct {
	ChannelId        int    `json:"cid,string"`
	ClientId         int    `json:"clid,string"`
	ClientDatabaseId int    `json:"client_database_id,string"`
	ClientNickname   string `json:"client_nickname"`
	ClientType       int    `json:"client_type,string"`
}

func (u *Client) IsBot() bool {
	return u.ClientType == 1
}

func (c *TeamspeakHttpClient) ClientList() (*[]Client, error) {
	var users []Client
	err := c.request("clientlist", &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

// clientmove `manage_scope, write_scope`
type ClientMoveRequest struct {
	ClientId        []int  `schema:"clid,required"`
	ChannelId       int    `schema:"cid,required"`
	ChannelPassword string `schema:"cpw,omitempty"`
}

func (c *TeamspeakHttpClient) ClientMove(request ClientMoveRequest) error {
	return c.requestWithParams("clientmove", request, nil)
}

// clientpermlist `manage_scope, write_scope, read_scope`
type clientPermissionListRequest struct {
	ClientDbId int  `schema:"cldbid"`
	AsString   bool `schema:"-permsid,omitempty"`
}

type ClientPermissionListResponse struct {
	PermissionId      int `json:"permid,string"`
	PermissionNegated int `json:"permnegated,string"`
	PermissionSkip    int `json:"permskip,string"`
	PermissionValue   int `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ClientPermissionList(clientDbId int) (*[]ClientPermissionListResponse, error) {
	var perms []ClientPermissionListResponse

	err := c.requestWithParams(
		"clientpermlist",
		clientPermissionListRequest{
			ClientDbId: clientDbId,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

type ClientStringPermissionListResponse struct {
	PermissionId      string `json:"permsid"`
	PermissionNegated int    `json:"permnegated,string"`
	PermissionSkip    int    `json:"permskip,string"`
	PermissionValue   int    `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ClientStringPermissionList(clientDbId int) (*[]ClientStringPermissionListResponse, error) {
	var perms []ClientStringPermissionListResponse

	err := c.requestWithParams(
		"clientpermlist",
		clientPermissionListRequest{
			ClientDbId: clientDbId,
			AsString:   true,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

// clientpoke `manage_scope, write_scope`
type clientPokeRequest struct {
	ClientId int    `schema:"clid,required"`
	Message  string `schema:"msg,required"`
}

func (c *TeamspeakHttpClient) ClientPoke(clientId int, message string) error {
	return c.requestWithParams("clientpoke", clientPokeRequest{ClientId: clientId, Message: message}, nil)
}

// clientsetserverquerylogin `manage_scope, write_scope`
type clientSetServerQueryLoginRequest struct {
	ClientLoginName string `schema:"client_login_name"`
}

func (c *TeamspeakHttpClient) ClientSetServerQueryLogin(clientLoginName string) error {
	return c.requestWithParams(
		"clientsetserverquerylogin",
		clientSetServerQueryLoginRequest{ClientLoginName: clientLoginName},
		nil,
	)
}

// clientupdate `manage_scope, write_scope`
type ClientUpdateRequest struct {
	ClientNickname           string `schema:"client_nickname"`
	ClientIsTalker           string `schema:"client_is_talker"`
	ClientDescription        string `schema:"client_description"`
	ClientIsChannelCommander string `schema:"client_is_channel_commander"`
	ClientIconID             string `schema:"client_icon_id"`
}

func (c *TeamspeakHttpClient) ClientUpdate(request ClientUpdateRequest) error {
	return c.requestWithParams("clientupdate", request, nil)
}

// setclientchannelgroup `manage_scope, write_scope`
type setClientChannelGroupRequest struct {
	ChannelGroupId int `schema:"cgid"`
	ChannelId      int `schema:"cid"`
	ClientDbId     int `schema:"cldbid"`
}

func (c *TeamspeakHttpClient) SetClientChannelGroup(channelGroupId, channelId, clientDbId int) error {
	return c.requestWithParams(
		"setclientchannelgroup",
		setClientChannelGroupRequest{
			ChannelGroupId: channelGroupId,
			ChannelId:      channelId,
			ClientDbId:     clientDbId,
		},
		nil,
	)
}
