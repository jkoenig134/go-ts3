package go_ts3_http

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type WhoamiBody struct {
	ClientChannelId               string `json:"client_channel_id"`
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

type Whoami struct {
	WhoamiInfo []WhoamiBody `json:"body"`
	Status     Status       `json:"status"`
}

//TODO remove
type Placeholder struct {
	Status Status `json:"status"`
}

func (c *Client) Whoami() (*Whoami, error) {
	whoami := &Whoami{}
	err := c.request("whoami", whoami)
	if err != nil {
		return nil, err
	}

	return whoami, nil
}

type User struct {
	ChannelId        string `json:"cid"`
	ClientId         string `json:"clid"`
	ClientDatabaseId string `json:"client_database_id"`
	ClientNickname   string `json:"client_nickname"`
	ClientType       string `json:"client_type"`
}

func (u *User) IsBot() bool {
	return u.ClientType == "1"
}

type ClientList struct {
	Users  []User `json:"body"`
	Status Status `json:"status"`
}

func (c *Client) ClientList(server int) (*ClientList, error) {
	clientList := &ClientList{}
	err := c.request(vServerUrl(server, "clientlist"), clientList)
	if err != nil {
		return nil, err
	}

	return clientList, nil
}

type Channel struct {
	ChannelName                 string `json:"channel_name"`
	ChannelNeededSubscribePower string `json:"channel_needed_subscribe_power"`
	ChannelOrder                string `json:"channel_order"`
	ChannelId                   string `json:"cid"`
	PID                         string `json:"pid"`
	TotalClients                string `json:"total_clients"`
}

type ChannelList struct {
	Channels []Channel `json:"body"`
	Status   Status    `json:"status"`
}

func (c *Client) ChannelList(server int) (*ChannelList, error) {
	channelList := &ChannelList{}

	err := c.request(vServerUrl(server, "channellist"), channelList)
	if err != nil {
		return nil, err
	}

	return channelList, nil
}
