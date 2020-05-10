package go_ts3_http

type HostInfo struct {
	ConnectionBandwidthReceivedLastMinuteTotal int `json:"connection_bandwidth_received_last_minute_total,string"`
	ConnectionBandwidthReceivedLastSecondTotal int `json:"connection_bandwidth_received_last_second_total,string"`
	ConnectionBandwidthSentLastMinuteTotal     int `json:"connection_bandwidth_sent_last_minute_total,string"`
	ConnectionBandwidthSentLastSecondTotal     int `json:"connection_bandwidth_sent_last_second_total,string"`
	ConnectionBytesReceivedTotal               int `json:"connection_bytes_received_total,string"`
	ConnectionBytesSentTotal                   int `json:"connection_bytes_sent_total,string"`
	ConnectionFiletransferBandwidthReceived    int `json:"connection_filetransfer_bandwidth_received,string"`
	ConnectionFiletransferBandwidthSent        int `json:"connection_filetransfer_bandwidth_sent,string"`
	ConnectionFiletransferBytesReceivedTotal   int `json:"connection_filetransfer_bytes_received_total,string"`
	ConnectionFiletransferBytesSentTotal       int `json:"connection_filetransfer_bytes_sent_total,string"`
	ConnectionPacketsReceivedTotal             int `json:"connection_packets_received_total,string"`
	ConnectionPacketsSentTotal                 int `json:"connection_packets_sent_total,string"`
	HostTimestampUtc                           int `json:"host_timestamp_utc,string"`
	InstanceUptime                             int `json:"instance_uptime,string"`
	VirtualserversRunningTotal                 int `json:"virtualservers_running_total,string"`
	VirtualserversTotalChannelsOnline          int `json:"virtualservers_total_channels_online,string"`
	VirtualserversTotalClientsOnline           int `json:"virtualservers_total_clients_online,string"`
	VirtualserversTotalMaxclients              int `json:"virtualservers_total_maxclients,string"`
}

// hostinfo `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) HostInfo() (*HostInfo, error) {
	var hostInfo []HostInfo
	err := c.request("hostinfo", &hostInfo)
	if err != nil {
		return nil, err
	}

	return &hostInfo[0], nil
}

type Version struct {
	Build    string `json:"build"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

// version `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) Version() (*Version, error) {
	var version []Version
	err := c.request("version", &version)
	if err != nil {
		return nil, err
	}

	return &version[0], nil
}

type WhoamiInfo struct {
	ClientChannelId               int    `json:"client_channel_id,string"`
	ClientDatabaseId              int    `json:"client_database_id,string"`
	ClientId                      int    `json:"client_id,string"`
	ClientLoginName               string `json:"client_login_name"`
	ClientNickname                string `json:"client_nickname"`
	ClientOriginServerId          int    `json:"client_origin_server_id,string"`
	ClientUniqueIdentifier        string `json:"client_unique_identifier"`
	VirtualserverId               int    `json:"virtualserver_id,string"`
	VirtualserverPort             int    `json:"virtualserver_port,string"`
	VirtualserverStatus           string `json:"virtualserver_status"`
	VirtualserverUniqueIdentifier string `json:"virtualserver_unique_identifier"`
}

// whoami `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) Whoami() (*WhoamiInfo, error) {
	var whoami []WhoamiInfo
	err := c.request("whoami", &whoami)
	if err != nil {
		return nil, err
	}

	return &whoami[0], nil
}

type Subsystem string

//noinspection GoUnusedConst
const (
	VOICE        Subsystem = "voice"
	QUERY        Subsystem = "query"
	FILETRANSFER Subsystem = "filetransfer"
)

type BindingListRequest struct {
	Subsystem Subsystem `schema:"subsystem,omitempty"`
}

type Binding struct {
	IP string `json:"ip"`
}

// bindinglist `manage_scope, read_scope`
func (c *TeamspeakHttpClient) BindingList(request BindingListRequest) (*[]Binding, error) {
	var bindings []Binding
	err := c.requestWithParams("bindinglist", request, &bindings)
	if err != nil {
		return nil, err
	}

	return &bindings, nil
}
