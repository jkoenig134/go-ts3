package go_ts3_http

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
	var users []Client
	err := c.request(vServerUrl(server, "clientlist"), &users)
	if err != nil {
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
	var channels []Channel
	err := c.request(vServerUrl(server, "channellist"), &channels)
	if err != nil {
		return nil, err
	}

	return &channels, nil
}
