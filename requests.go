package go_ts3_http

import (
	"encoding/json"
	"fmt"
)

type WhoamiInfo struct {
	ClientChannelId               int    `json:"client_channel_id,string"`
	ClientDatabaseId              string `json:"client_database_id"`
	ClientId                      string `json:"client_id"`
	ClientLoginName               string `json:"client_login_name"`
	ClientNickname                string `json:"client_nickname"`
	ClientOriginServerId          string `json:"client_origin_server_id"`
	ClientUniqueIdentifier        string `json:"client_unique_identifier"`
	VirtualserverId               string `json:"virtualserver_id"`
	VirtualserverPort             string `json:"virtualserver_port"`
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
	ChannelId        string `json:"cid"`
	ClientId         string `json:"clid"`
	ClientDatabaseId string `json:"client_database_id"`
	ClientNickname   string `json:"client_nickname"`
	ClientType       string `json:"client_type"`
}

func (u *Client) IsBot() bool {
	return u.ClientType == "1"
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
	ChannelNeededSubscribePower string `json:"channel_needed_subscribe_power"`
	ChannelOrder                string `json:"channel_order"`
	ChannelId                   string `json:"cid"`
	PID                         string `json:"pid"`
	TotalClients                string `json:"total_clients"`
}

func (c *TeamspeakHttpClient) ChannelList(server int) (*[]Channel, error) {
	body, err := c.request(vServerUrl(server, "channellist"))
	if err != nil {
		return nil, err
	}

	fmt.Println(string(*body))

	var channels []Channel
	if err = json.Unmarshal(*body, &channels); err != nil {
		return nil, err
	}

	return &channels, nil
}

type Version struct {
	//"build":"1585305527","platform":"Linux","version":"3.12.1"
	Build    string `json:"build"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

func (c *TeamspeakHttpClient) Version() (*Version, error) {
	body, err := c.request("version")
	if err != nil {
		return nil, err
	}

	var whoami []Version
	if err = json.Unmarshal(*body, &whoami); err != nil {
		return nil, err
	}

	return &whoami[0], nil
}
