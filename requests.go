package go_ts3_http

import (
	"encoding/json"
)

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

func (c *TeamspeakHttpClient) Whoami() (*WhoamiInfo, error) {
	body, err := c.request("whoami")
	if err != nil {
		return nil, err
	}

	var whoami []WhoamiInfo
	if err = json.Unmarshal(*body, &whoami); err != nil {
		return nil, err
	}

	return &whoami[0], nil
}

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

func (c *TeamspeakHttpClient) ClientList(server int) (*[]Client, error) {
	body, err := c.request(vServerUrl(server, "clientlist"))
	if err != nil {
		return nil, err
	}

	var users []Client
	if err = json.Unmarshal(*body, &users); err != nil {
		return nil, err
	}

	return &users, nil
}

type Channel struct {
	ChannelName                 string `json:"channel_name"`
	ChannelNeededSubscribePower int    `json:"channel_needed_subscribe_power,string"`
	ChannelOrder                int    `json:"channel_order,string"`
	ChannelId                   int    `json:"cid,string"`
	PID                         int    `json:"pid,string"`
	TotalClients                int    `json:"total_clients,string"`
}

func (c *TeamspeakHttpClient) ChannelList(server int) (*[]Channel, error) {
	body, err := c.request(vServerUrl(server, "channellist"))
	if err != nil {
		return nil, err
	}

	var channels []Channel
	if err = json.Unmarshal(*body, &channels); err != nil {
		return nil, err
	}

	return &channels, nil
}

type Version struct {
	Build    string `json:"build"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

func (c *TeamspeakHttpClient) Version() (*Version, error) {
	body, err := c.request("version")
	if err != nil {
		return nil, err
	}

	var version []Version
	if err = json.Unmarshal(*body, &version); err != nil {
		return nil, err
	}

	return &version[0], nil
}
