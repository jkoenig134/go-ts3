package go_ts3_http

// clientaddperm `manage_scope`
func (c *TeamspeakHttpClient) ClientAddPerm() error {
	//TODO
	return nil
}

// clientdbdelete `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ClientDbDelete() error {
	//TODO
	return nil
}

// clientdbedit `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ClientDbEdit() error {
	//TODO
	return nil
}

// clientdbfind `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientDbFind() error {
	//TODO
	return nil
}

// clientdbinfo `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientDbInfo() error {
	//TODO
	return nil
}

// clientdblist `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientDbList() error {
	//TODO
	return nil
}

// clientdelperm `manage_scope`
func (c *TeamspeakHttpClient) ClientDeletePermission() error {
	//TODO
	return nil
}

// clientedit `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ClientEdit() error {
	//TODO
	return nil
}

// clientfind `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientFind() error {
	//TODO
	return nil
}

// clientgetdbidfromuid `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientGetDbIdFromUid() error {
	//TODO
	return nil
}

// clientgetids `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientGetIds() error {
	//TODO
	return nil
}

// clientgetnamefromdbid `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientGetNameFromDbId() error {
	//TODO
	return nil
}

// clientgetnamefromuid `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientGetNameFromUid() error {
	//TODO
	return nil
}

// clientgetuidfromclid `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientGetUidFromClientId() error {
	//TODO
	return nil
}

// clientinfo `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientInfo() error {
	//TODO
	return nil
}

// clientkick `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ClientKick() error {
	//TODO
	return nil
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
func (c *TeamspeakHttpClient) ClientMove() error {
	//TODO
	return nil
}

// clientpermlist `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientPermissionList() error {
	//TODO
	return nil
}

// clientpoke `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ClientPoke() error {
	//TODO
	return nil
}

// clientsetserverquerylogin `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ClientSetServerQueryLogin() error {
	//TODO
	return nil
}

// clientupdate `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ClientUpdate() error {
	//TODO
	return nil
}

// setclientchannelgroup `manage_scope, write_scope`
func (c *TeamspeakHttpClient) SetClientChannelGroup() error {
	//TODO
	return nil
}
