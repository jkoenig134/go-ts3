package go_ts3_http

//servercreate `manage_scope`
type ServerCreateRequest struct {
	VirtualserverName                                   string `schema:"virtualserver_name,required"`
	VirtualserverWelcomemessage                         string `schema:"virtualserver_welcomemessage,omitempty"`
	VirtualserverMaxclients                             string `schema:"virtualserver_maxclients,omitempty"`
	VirtualserverPassword                               string `schema:"virtualserver_password,omitempty"`
	VirtualserverHostmessage                            string `schema:"virtualserver_hostmessage,omitempty"`
	VirtualserverHostmessageMode                        string `schema:"virtualserver_hostmessage_mode,omitempty"`
	VirtualserverDefaultServerGroup                     string `schema:"virtualserver_default_server_group,omitempty"`
	VirtualserverDefaultChannelGroup                    string `schema:"virtualserver_default_channel_group,omitempty"`
	VirtualserverDefaultChannelAdminGroup               string `schema:"virtualserver_default_channel_admin_group,omitempty"`
	VirtualserverMaxDownloadTotalBandwidth              string `schema:"virtualserver_max_download_total_bandwidth,omitempty"`
	VirtualserverMaxUploadTotalBandwidth                string `schema:"virtualserver_max_upload_total_bandwidth,omitempty"`
	VirtualserverHostbannerURL                          string `schema:"virtualserver_hostbanner_url,omitempty"`
	VirtualserverHostbannerGfxURL                       string `schema:"virtualserver_hostbanner_gfx_url,omitempty"`
	VirtualserverHostbannerGfxInterval                  string `schema:"virtualserver_hostbanner_gfx_interval,omitempty"`
	VirtualserverComplainAutobanCount                   string `schema:"virtualserver_complain_autoban_count,omitempty"`
	VirtualserverComplainAutobanTime                    string `schema:"virtualserver_complain_autoban_time,omitempty"`
	VirtualserverComplainRemoveTime                     string `schema:"virtualserver_complain_remove_time,omitempty"`
	VirtualserverMinClientsInChannelBeforeForcedSilence string `schema:"virtualserver_min_clients_in_channel_before_forced_silence,omitempty"`
	VirtualserverPrioritySpeakerDimmModificator         string `schema:"virtualserver_priority_speaker_dimm_modificator,omitempty"`
	VirtualserverAntifloodPointsTickReduce              string `schema:"virtualserver_antiflood_points_tick_reduce,omitempty"`
	VirtualserverAntifloodPointsNeededCommandBlock      string `schema:"virtualserver_antiflood_points_needed_command_block,omitempty"`
	VirtualserverAntifloodPointsNeededPluginBlock       string `schema:"virtualserver_antiflood_points_needed_plugin_block,omitempty"`
	VirtualserverAntifloodPointsNeededIPBlock           string `schema:"virtualserver_antiflood_points_needed_ip_block,omitempty"`
	VirtualserverHostbannerMode                         string `schema:"virtualserver_hostbanner_mode,omitempty"`
	VirtualserverHostbuttonTooltip                      string `schema:"virtualserver_hostbutton_tooltip,omitempty"`
	VirtualserverHostbuttonGfxURL                       string `schema:"virtualserver_hostbutton_gfx_url,omitempty"`
	VirtualserverHostbuttonURL                          string `schema:"virtualserver_hostbutton_url,omitempty"`
	VirtualserverDownloadQuota                          string `schema:"virtualserver_download_quota,omitempty"`
	VirtualserverUploadQuota                            string `schema:"virtualserver_upload_quota,omitempty"`
	VirtualserverMachineID                              string `schema:"virtualserver_machine_id,omitempty"`
	VirtualserverPort                                   string `schema:"virtualserver_port,omitempty"`
	VirtualserverAutostart                              string `schema:"virtualserver_autostart,omitempty"`
	VirtualserverStatus                                 string `schema:"virtualserver_status,omitempty"`
	VirtualserverLogClient                              string `schema:"virtualserver_log_client,omitempty"`
	VirtualserverLogQuery                               string `schema:"virtualserver_log_query,omitempty"`
	VirtualserverLogChannel                             string `schema:"virtualserver_log_channel,omitempty"`
	VirtualserverLogPermissions                         string `schema:"virtualserver_log_permissions,omitempty"`
	VirtualserverLogServer                              string `schema:"virtualserver_log_server,omitempty"`
	VirtualserverLogFiletransfer                        string `schema:"virtualserver_log_filetransfer,omitempty"`
	VirtualserverMinClientVersion                       string `schema:"virtualserver_min_client_version,omitempty"`
	VirtualserverMinAndroidVersion                      string `schema:"virtualserver_min_android_version,omitempty"`
	VirtualserverMinIosVersion                          string `schema:"virtualserver_min_ios_version,omitempty"`
	VirtualserverMinWinphoneVersion                     string `schema:"virtualserver_min_winphone_version,omitempty"`
	VirtualserverNeededIdentitySecurityLevel            string `schema:"virtualserver_needed_identity_security_level,omitempty"`
	VirtualserverNamePhonetic                           string `schema:"virtualserver_name_phonetic,omitempty"`
	VirtualserverIconID                                 string `schema:"virtualserver_icon_id,omitempty"`
	VirtualserverReservedSlots                          string `schema:"virtualserver_reserved_slots,omitempty"`
	VirtualserverWeblistEnabled                         string `schema:"virtualserver_weblist_enabled,omitempty"`
	VirtualserverCodecEncryptionMode                    string `schema:"virtualserver_codec_encryption_mode,omitempty"`
}

type ServerCreateResponse struct {
	ServerId          int    `json:"sid,string"`
	VirtualServerPort int    `json:"virtualserver_port,string"`
	Token             string `json:"token"`
}

func (c *TeamspeakHttpClient) ServerCreate(request ServerCreateRequest) (*ServerCreateResponse, error) {
	var responses []ServerCreateResponse

	err := c.requestWithParams("servercreate", request, &responses)
	if err != nil {
		return nil, err
	}

	return &responses[0], nil
}

//serverdelete `manage_scope`
type serverDeleteRequest struct {
	ServerId int `schema:"sid,required"`
}

func (c *TeamspeakHttpClient) ServerDelete(serverId int) error {
	return c.requestWithParams(
		"serverdelete",
		serverDeleteRequest{ServerId: serverId},
		nil,
	)
}

//serveredit `manage_scope, write_scope`
type ServerEditRequest struct {
	VirtualserverName                                   string `schema:"virtualserver_name,omitempty"`
	VirtualserverWelcomemessage                         string `schema:"virtualserver_welcomemessage,omitempty"`
	VirtualserverMaxclients                             string `schema:"virtualserver_maxclients,omitempty"`
	VirtualserverPassword                               string `schema:"virtualserver_password,omitempty"`
	VirtualserverHostmessage                            string `schema:"virtualserver_hostmessage,omitempty"`
	VirtualserverHostmessageMode                        string `schema:"virtualserver_hostmessage_mode,omitempty"`
	VirtualserverDefaultServerGroup                     string `schema:"virtualserver_default_server_group,omitempty"`
	VirtualserverDefaultChannelGroup                    string `schema:"virtualserver_default_channel_group,omitempty"`
	VirtualserverDefaultChannelAdminGroup               string `schema:"virtualserver_default_channel_admin_group,omitempty"`
	VirtualserverMaxDownloadTotalBandwidth              string `schema:"virtualserver_max_download_total_bandwidth,omitempty"`
	VirtualserverMaxUploadTotalBandwidth                string `schema:"virtualserver_max_upload_total_bandwidth,omitempty"`
	VirtualserverHostbannerURL                          string `schema:"virtualserver_hostbanner_url,omitempty"`
	VirtualserverHostbannerGfxURL                       string `schema:"virtualserver_hostbanner_gfx_url,omitempty"`
	VirtualserverHostbannerGfxInterval                  string `schema:"virtualserver_hostbanner_gfx_interval,omitempty"`
	VirtualserverComplainAutobanCount                   string `schema:"virtualserver_complain_autoban_count,omitempty"`
	VirtualserverComplainAutobanTime                    string `schema:"virtualserver_complain_autoban_time,omitempty"`
	VirtualserverComplainRemoveTime                     string `schema:"virtualserver_complain_remove_time,omitempty"`
	VirtualserverMinClientsInChannelBeforeForcedSilence string `schema:"virtualserver_min_clients_in_channel_before_forced_silence,omitempty"`
	VirtualserverPrioritySpeakerDimmModificator         string `schema:"virtualserver_priority_speaker_dimm_modificator,omitempty"`
	VirtualserverAntifloodPointsTickReduce              string `schema:"virtualserver_antiflood_points_tick_reduce,omitempty"`
	VirtualserverAntifloodPointsNeededCommandBlock      string `schema:"virtualserver_antiflood_points_needed_command_block,omitempty"`
	VirtualserverAntifloodPointsNeededPluginBlock       string `schema:"virtualserver_antiflood_points_needed_plugin_block,omitempty"`
	VirtualserverAntifloodPointsNeededIPBlock           string `schema:"virtualserver_antiflood_points_needed_ip_block,omitempty"`
	VirtualserverHostbannerMode                         string `schema:"virtualserver_hostbanner_mode,omitempty"`
	VirtualserverHostbuttonTooltip                      string `schema:"virtualserver_hostbutton_tooltip,omitempty"`
	VirtualserverHostbuttonGfxURL                       string `schema:"virtualserver_hostbutton_gfx_url,omitempty"`
	VirtualserverHostbuttonURL                          string `schema:"virtualserver_hostbutton_url,omitempty"`
	VirtualserverDownloadQuota                          string `schema:"virtualserver_download_quota,omitempty"`
	VirtualserverUploadQuota                            string `schema:"virtualserver_upload_quota,omitempty"`
	VirtualserverMachineID                              string `schema:"virtualserver_machine_id,omitempty"`
	VirtualserverPort                                   string `schema:"virtualserver_port,omitempty"`
	VirtualserverAutostart                              string `schema:"virtualserver_autostart,omitempty"`
	VirtualserverStatus                                 string `schema:"virtualserver_status,omitempty"`
	VirtualserverLogClient                              string `schema:"virtualserver_log_client,omitempty"`
	VirtualserverLogQuery                               string `schema:"virtualserver_log_query,omitempty"`
	VirtualserverLogChannel                             string `schema:"virtualserver_log_channel,omitempty"`
	VirtualserverLogPermissions                         string `schema:"virtualserver_log_permissions,omitempty"`
	VirtualserverLogServer                              string `schema:"virtualserver_log_server,omitempty"`
	VirtualserverLogFiletransfer                        string `schema:"virtualserver_log_filetransfer,omitempty"`
	VirtualserverMinClientVersion                       string `schema:"virtualserver_min_client_version,omitempty"`
	VirtualserverMinAndroidVersion                      string `schema:"virtualserver_min_android_version,omitempty"`
	VirtualserverMinIosVersion                          string `schema:"virtualserver_min_ios_version,omitempty"`
	VirtualserverMinWinphoneVersion                     string `schema:"virtualserver_min_winphone_version,omitempty"`
	VirtualserverNeededIdentitySecurityLevel            string `schema:"virtualserver_needed_identity_security_level,omitempty"`
	VirtualserverNamePhonetic                           string `schema:"virtualserver_name_phonetic,omitempty"`
	VirtualserverIconID                                 string `schema:"virtualserver_icon_id,omitempty"`
	VirtualserverReservedSlots                          string `schema:"virtualserver_reserved_slots,omitempty"`
	VirtualserverWeblistEnabled                         string `schema:"virtualserver_weblist_enabled,omitempty"`
	VirtualserverCodecEncryptionMode                    string `schema:"virtualserver_codec_encryption_mode,omitempty"`
}

func (c *TeamspeakHttpClient) ServerEdit(request ServerEditRequest) error {
	return c.requestWithParams("serveredit", request, nil)
}

//serveridgetbyport `manage_scope`
type serverIdGetByPortRequest struct {
	VirtualserverPort int `schema:"virtualserver_port,required"`
}

type serverIdGetByPortResponse struct {
	ServerId int `json:"server_id,string"`
}

func (c *TeamspeakHttpClient) ServerIdGetByPort(port int) (*int, error) {
	var responses []serverIdGetByPortResponse

	err := c.requestWithParams(
		"",
		serverIdGetByPortRequest{VirtualserverPort: port},
		&responses,
	)
	if err != nil {
		return nil, err
	}

	return &responses[0].ServerId, nil
}

//serverinfo `manage_scope`
type ServerInfo struct {
	ConnectionBandwidthReceivedLastMinuteTotal          int     `json:"connection_bandwidth_received_last_minute_total,string"`
	ConnectionBandwidthReceivedLastSecondTotal          int     `json:"connection_bandwidth_received_last_second_total,string"`
	ConnectionBandwidthSentLastMinuteTotal              int     `json:"connection_bandwidth_sent_last_minute_total,string"`
	ConnectionBandwidthSentLastSecondTotal              int     `json:"connection_bandwidth_sent_last_second_total,string"`
	ConnectionBytesReceivedControl                      int     `json:"connection_bytes_received_control,string"`
	ConnectionBytesReceivedKeepalive                    int     `json:"connection_bytes_received_keepalive,string"`
	ConnectionBytesReceivedSpeech                       int     `json:"connection_bytes_received_speech,string"`
	ConnectionBytesReceivedTotal                        int     `json:"connection_bytes_received_total,string"`
	ConnectionBytesSentControl                          int     `json:"connection_bytes_sent_control,string"`
	ConnectionBytesSentKeepalive                        int     `json:"connection_bytes_sent_keepalive,string"`
	ConnectionBytesSentSpeech                           int     `json:"connection_bytes_sent_speech,string"`
	ConnectionBytesSentTotal                            int     `json:"connection_bytes_sent_total,string"`
	ConnectionFiletransferBandwidthReceived             int     `json:"connection_filetransfer_bandwidth_received,string"`
	ConnectionFiletransferBandwidthSent                 int     `json:"connection_filetransfer_bandwidth_sent,string"`
	ConnectionFiletransferBytesReceivedTotal            int     `json:"connection_filetransfer_bytes_received_total,string"`
	ConnectionFiletransferBytesSentTotal                int     `json:"connection_filetransfer_bytes_sent_total,string"`
	ConnectionPacketsReceivedControl                    int     `json:"connection_packets_received_control,string"`
	ConnectionPacketsReceivedKeepalive                  int     `json:"connection_packets_received_keepalive,string"`
	ConnectionPacketsReceivedSpeech                     int     `json:"connection_packets_received_speech,string"`
	ConnectionPacketsReceivedTotal                      int     `json:"connection_packets_received_total,string"`
	ConnectionPacketsSentControl                        int     `json:"connection_packets_sent_control,string"`
	ConnectionPacketsSentKeepalive                      int     `json:"connection_packets_sent_keepalive,string"`
	ConnectionPacketsSentSpeech                         int     `json:"connection_packets_sent_speech,string"`
	ConnectionPacketsSentTotal                          int     `json:"connection_packets_sent_total,string"`
	VirtualserverAntifloodPointsNeededCommandBlock      int     `json:"virtualserver_antiflood_points_needed_command_block,string"`
	VirtualserverAntifloodPointsNeededIPBlock           int     `json:"virtualserver_antiflood_points_needed_ip_block,string"`
	VirtualserverAntifloodPointsNeededPluginBlock       int     `json:"virtualserver_antiflood_points_needed_plugin_block,string"`
	VirtualserverAntifloodPointsTickReduce              int     `json:"virtualserver_antiflood_points_tick_reduce,string"`
	VirtualserverAskForPrivilegekey                     int     `json:"virtualserver_ask_for_privilegekey,string"`
	VirtualserverAutostart                              int     `json:"virtualserver_autostart,string"`
	VirtualserverChannelTempDeleteDelayDefault          int     `json:"virtualserver_channel_temp_delete_delay_default,string"`
	VirtualserverChannelsOnline                         int     `json:"virtualserver_channelsonline,string"`
	VirtualserverClientConnections                      int     `json:"virtualserver_client_connections,string"`
	VirtualserverClientsonline                          int     `json:"virtualserver_clientsonline,string"`
	VirtualserverCodecEncryptionMode                    int     `json:"virtualserver_codec_encryption_mode,string"`
	VirtualserverComplainAutoBanCount                   int     `json:"virtualserver_complain_autoban_count,string"`
	VirtualserverComplainAutoBanTime                    int     `json:"virtualserver_complain_autoban_time,string"`
	VirtualserverComplainRemoveTime                     int     `json:"virtualserver_complain_remove_time,string"`
	VirtualserverCreated                                int     `json:"virtualserver_created,string"`
	VirtualserverDefaultChannelAdminGroup               int     `json:"virtualserver_default_channel_admin_group,string"`
	VirtualserverDefaultChannelGroup                    int     `json:"virtualserver_default_channel_group,string"`
	VirtualserverDefaultServerGroup                     int     `json:"virtualserver_default_server_group,string"`
	VirtualserverDownloadQuota                          uint    `json:"virtualserver_download_quota,string"`
	VirtualserverFilebase                               string  `json:"virtualserver_filebase"`
	VirtualserverFlagPassword                           int     `json:"virtualserver_flag_password,string"`
	VirtualserverHostbannerGfxInterval                  int     `json:"virtualserver_hostbanner_gfx_interval,string"`
	VirtualserverHostbannerGfxURL                       string  `json:"virtualserver_hostbanner_gfx_url"`
	VirtualserverHostbannerMode                         int     `json:"virtualserver_hostbanner_mode,string"`
	VirtualserverHostbannerURL                          string  `json:"virtualserver_hostbanner_url"`
	VirtualserverHostbuttonGfxURL                       string  `json:"virtualserver_hostbutton_gfx_url"`
	VirtualserverHostbuttonTooltip                      string  `json:"virtualserver_hostbutton_tooltip"`
	VirtualserverHostbuttonURL                          string  `json:"virtualserver_hostbutton_url"`
	VirtualserverHostmessage                            string  `json:"virtualserver_hostmessage"`
	VirtualserverHostmessageMode                        int     `json:"virtualserver_hostmessage_mode,string"`
	VirtualserverIconID                                 int     `json:"virtualserver_icon_id,string"`
	VirtualserverID                                     int     `json:"virtualserver_id,string"`
	VirtualserverIP                                     string  `json:"virtualserver_ip"`
	VirtualserverLogChannel                             int     `json:"virtualserver_log_channel,string"`
	VirtualserverLogClient                              int     `json:"virtualserver_log_client,string"`
	VirtualserverLogFiletransfer                        int     `json:"virtualserver_log_filetransfer,string"`
	VirtualserverLogPermissions                         int     `json:"virtualserver_log_permissions,string"`
	VirtualserverLogQuery                               int     `json:"virtualserver_log_query,string"`
	VirtualserverLogServer                              int     `json:"virtualserver_log_server,string"`
	VirtualserverMachineID                              string  `json:"virtualserver_machine_id"`
	VirtualserverMaxDownloadTotalBandwidth              uint    `json:"virtualserver_max_download_total_bandwidth,string"`
	VirtualserverMaxUploadTotalBandwidth                uint    `json:"virtualserver_max_upload_total_bandwidth,string"`
	VirtualserverMaxclients                             int     `json:"virtualserver_maxclients,string"`
	VirtualserverMinAndroidVersion                      int     `json:"virtualserver_min_android_version,string"`
	VirtualserverMinClientVersion                       int     `json:"virtualserver_min_client_version,string"`
	VirtualserverMinClientsInChannelBeforeForcedSilence int     `json:"virtualserver_min_clients_in_channel_before_forced_silence,string"`
	VirtualserverMinIosVersion                          int     `json:"virtualserver_min_ios_version,string"`
	VirtualserverMonthBytesDownloaded                   int     `json:"virtualserver_month_bytes_downloaded,string"`
	VirtualserverMonthBytesUploaded                     int     `json:"virtualserver_month_bytes_uploaded,string"`
	VirtualserverName                                   string  `json:"virtualserver_name"`
	VirtualserverNamePhonetic                           string  `json:"virtualserver_name_phonetic"`
	VirtualserverNeededIdentitySecurityLevel            int     `json:"virtualserver_needed_identity_security_level,string"`
	VirtualserverNickname                               string  `json:"virtualserver_nickname"`
	VirtualserverPassword                               string  `json:"virtualserver_password"`
	VirtualserverPlatform                               string  `json:"virtualserver_platform"`
	VirtualserverPort                                   int     `json:"virtualserver_port,string"`
	VirtualserverPrioritySpeakerDimmModificator         float64 `json:"virtualserver_priority_speaker_dimm_modificator,string"`
	VirtualserverQueryClientConnections                 int     `json:"virtualserver_query_client_connections,string"`
	VirtualserverQueryclientsonline                     int     `json:"virtualserver_queryclientsonline,string"`
	VirtualserverReservedSlots                          int     `json:"virtualserver_reserved_slots,string"`
	VirtualserverStatus                                 string  `json:"virtualserver_status"`
	VirtualserverTotalBytesDownloaded                   int     `json:"virtualserver_total_bytes_downloaded,string"`
	VirtualserverTotalBytesUploaded                     int     `json:"virtualserver_total_bytes_uploaded,string"`
	VirtualserverTotalPacketlossControl                 float64 `json:"virtualserver_total_packetloss_control,string"`
	VirtualserverTotalPacketlossKeepalive               float64 `json:"virtualserver_total_packetloss_keepalive,string"`
	VirtualserverTotalPacketlossSpeech                  float64 `json:"virtualserver_total_packetloss_speech,string"`
	VirtualserverTotalPacketlossTotal                   float64 `json:"virtualserver_total_packetloss_total,string"`
	VirtualserverTotalPing                              float64 `json:"virtualserver_total_ping,string"`
	VirtualserverUniqueIdentifier                       string  `json:"virtualserver_unique_identifier"`
	VirtualserverUploadQuota                            uint    `json:"virtualserver_upload_quota,string"`
	VirtualserverUptime                                 int     `json:"virtualserver_uptime,string"`
	VirtualserverVersion                                string  `json:"virtualserver_version"`
	VirtualserverWeblistEnabled                         int     `json:"virtualserver_weblist_enabled,string"`
	VirtualserverWelcomemessage                         string  `json:"virtualserver_welcomemessage"`
}

func (c *TeamspeakHttpClient) ServerInfo() (*ServerInfo, error) {
	var serverInfos []ServerInfo

	err := c.request("serverinfo", &serverInfos)
	if err != nil {
		return nil, err
	}

	return &serverInfos[0], nil
}

//serverlist `manage_scope`
type Server struct {
	VirtualserverAutostart          int    `json:"virtualserver_autostart,string"`
	VirtualserverClientsonline      int    `json:"virtualserver_clientsonline,string"`
	VirtualserverID                 int    `json:"virtualserver_id,string"`
	VirtualserverMachineID          string `json:"virtualserver_machine_id"`
	VirtualserverMaxclients         int    `json:"virtualserver_maxclients,string"`
	VirtualserverName               string `json:"virtualserver_name"`
	VirtualserverPort               int    `json:"virtualserver_port,string"`
	VirtualserverQueryclientsonline int    `json:"virtualserver_queryclientsonline,string"`
	VirtualserverStatus             string `json:"virtualserver_status"`
	VirtualserverUptime             int    `json:"virtualserver_uptime,string"`
}

func (c *TeamspeakHttpClient) ServerList() (*[]Server, error) {
	var servers []Server

	err := c.request("serverlist", &servers)
	if err != nil {
		return nil, err
	}

	return &servers, nil
}

//serverprocessstop `manage_scope`
type ServerProcessStopRequest struct {
	ReasonMessage string `schema:"reasonmsg,omitempty"`
}

func (c *TeamspeakHttpClient) ServerProcessStop(request ServerProcessStopRequest) error {
	return c.requestWithParams("serverprocessstop", request, nil)
}

//serverrequestconnectioninfo `manage_scope, write_scope, read_scope`
type ServerRequestConnectionInfo struct {
	ConnectionBandwidthReceivedLastMinuteTotal string `json:"connection_bandwidth_received_last_minute_total"`
	ConnectionBandwidthReceivedLastSecondTotal string `json:"connection_bandwidth_received_last_second_total"`
	ConnectionBandwidthSentLastMinuteTotal     string `json:"connection_bandwidth_sent_last_minute_total"`
	ConnectionBandwidthSentLastSecondTotal     string `json:"connection_bandwidth_sent_last_second_total"`
	ConnectionBytesReceivedTotal               string `json:"connection_bytes_received_total"`
	ConnectionBytesSentTotal                   string `json:"connection_bytes_sent_total"`
	ConnectionConnectedTime                    string `json:"connection_connected_time"`
	ConnectionFiletransferBandwidthReceived    string `json:"connection_filetransfer_bandwidth_received"`
	ConnectionFiletransferBandwidthSent        string `json:"connection_filetransfer_bandwidth_sent"`
	ConnectionFiletransferBytesReceivedTotal   string `json:"connection_filetransfer_bytes_received_total"`
	ConnectionFiletransferBytesSentTotal       string `json:"connection_filetransfer_bytes_sent_total"`
	ConnectionPacketlossTotal                  string `json:"connection_packetloss_total"`
	ConnectionPacketsReceivedTotal             string `json:"connection_packets_received_total"`
	ConnectionPacketsSentTotal                 string `json:"connection_packets_sent_total"`
	ConnectionPing                             string `json:"connection_ping"`
}

func (c *TeamspeakHttpClient) ServerRequestConnectionInfo() (*[]ServerRequestConnectionInfo, error) {
	var infos []ServerRequestConnectionInfo

	err := c.request("serverrequestconnectioninfo", &infos)
	if err != nil {
		return nil, err
	}

	return &infos, nil
}

//serversnapshotcreate `manage_scope`
type ServerSnapshot struct {
	Data    string `json:"data"`
	Salt    string `json:"salt"`
	Version int    `json:"version,string"`
}

type ServerSnapshotCreateRequest struct {
	Password string `schema:"password,omitempty"`
}

func (c *TeamspeakHttpClient) ServerSnapshotCreate(request ServerSnapshotCreateRequest) (*ServerSnapshot, error) {
	var snapshots []ServerSnapshot

	err := c.requestWithParams("serversnapshotcreate", request, &snapshots)
	if err != nil {
		return nil, err
	}

	return &snapshots[0], nil
}

//serversnapshotdeploy `manage_scope`
func (original *ServerSnapshot) AsDeployRequest(password *string, mapping bool, keepFiles bool) ServerSnapshotDeployRequest {
	request := ServerSnapshotDeployRequest{
		Data:      original.Data,
		Salt:      original.Salt,
		Version:   original.Version,
		Mapping:   mapping,
		KeepFiles: keepFiles,
	}

	if password != nil {
		request.Password = *password
	}

	return request
}

type ServerSnapshotDeployRequest struct {
	Data      string `schema:"data,required"`
	Salt      string `schema:"salt,required"`
	Version   int    `schema:"version,required"`
	Password  string `schema:"password,omitempty"`
	Mapping   bool   `schema:"-mapping,omitempty"`
	KeepFiles bool   `schema:"-keepfiles,omitempty"`
}

type ServerSnapshotDeployResponse struct {
	OCId int `json:"ocid,string"`
	NCId int `json:"ncid,string"`
}

func (c *TeamspeakHttpClient) ServerSnapshotDeploy(request ServerSnapshotDeployRequest) (*[]ServerSnapshotDeployResponse, error) {
	var deployments []ServerSnapshotDeployResponse

	err := c.requestWithParams("serversnapshotdeploy", request, &deployments)
	if err != nil {
		return nil, err
	}

	return &deployments, nil
}

//serverstart `manage_scope`
type serverStartRequest struct {
	ServerId int `schema:"sid,required"`
}

func (c *TeamspeakHttpClient) ServerStart(serverId int) error {
	return c.requestWithParams("serverstart", serverStartRequest{ServerId: serverId}, nil)
}

//serverstop `manage_scope`
type ServerStopRequest struct {
	ServerId      int    `schema:"sid,required"`
	ReasonMessage string `schema:"reasonmsg,omitempty"`
}

func (c *TeamspeakHttpClient) ServerStop(request ServerStopRequest) error {
	return c.requestWithParams("serverstop", request, nil)
}
