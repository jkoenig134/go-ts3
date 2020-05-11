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

func (c *TeamspeakHttpClient) ClientList() (*[]Client, error) {
	var users []Client
	err := c.request("clientlist", &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
