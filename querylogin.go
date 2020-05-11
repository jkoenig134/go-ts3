package go_ts3_http

// queryloginadd `manage_scope, write_scope`
type QueryloginAddResponse struct {
	ClientDbId          int    `json:"cldbid,string"`
	ClientLoginName     string `json:"client_login_name"`
	ServerId            int    `json:"sid,string"`
	ClientLoginPassword string `json:"client_login_password"`
}

type queryloginAddRequest struct {
	ClientLoginName string `schema:"client_login_name,required"`
	ClientDbId      int    `schema:"cldbid,omitempty"`
}

func (c *TeamspeakHttpClient) QueryloginAddGlobal(clientLoginName string) (*QueryloginAddResponse, error) {
	var logins []QueryloginAddResponse

	oldServerId := c.serverID
	c.SetServerID(0)

	err := c.requestWithParams("queryloginadd", queryloginAddRequest{
		ClientLoginName: clientLoginName,
	}, &logins)

	c.SetServerID(oldServerId)

	if err != nil {
		return nil, err
	}

	return &logins[0], nil
}

func (c *TeamspeakHttpClient) QueryloginAddCurrent(clientLoginName string, clientDbId int) (*QueryloginAddResponse, error) {
	var logins []QueryloginAddResponse

	err := c.requestWithParams("queryloginadd", queryloginAddRequest{
		ClientLoginName: clientLoginName,
		ClientDbId:      clientDbId,
	}, &logins)
	if err != nil {
		return nil, err
	}

	return &logins[0], nil
}

// querylogindel `manage_scope, write_scope`
type queryloginDeleteRequest struct {
	ClientDbId int `schema:"cldbid,required"`
}

func (c *TeamspeakHttpClient) QueryloginDelete(clientDbId int) error {
	return c.requestWithParams("querylogindel", queryloginDeleteRequest{ClientDbId: clientDbId}, nil)
}

func (c *TeamspeakHttpClient) QueryloginDeleteGlobal(clientDbId int) error {
	oldServerId := c.serverID
	c.SetServerID(0)

	err := c.requestWithParams("querylogindel", queryloginDeleteRequest{ClientDbId: clientDbId}, nil)

	c.SetServerID(oldServerId)

	return err
}

// queryloginlist `manage_scope, write_scope, read_scope`
type Querylogin struct {
	ClientDbId      int    `json:"cldbid,string"`
	ClientLoginName string `json:"client_login_name"`
	ServerId        int    `json:"sid,string"`
}

type QueryloginListRequest struct {
	Pattern  string `schema:"pattern,omitempty"`
	Start    int    `schema:"start,omitempty"`
	Duration int    `schema:"duration,omitempty"`
}

func (c *TeamspeakHttpClient) QueryloginListGlobal(request QueryloginListRequest) (*[]Querylogin, error) {
	var logins []Querylogin

	oldServerId := c.serverID
	c.SetServerID(0)

	err := c.requestWithParams("queryloginlist", request, &logins)
	if err != nil {
		return nil, err
	}

	c.SetServerID(oldServerId)

	return &logins, nil
}

func (c *TeamspeakHttpClient) QueryloginList(request QueryloginListRequest) (*[]Querylogin, error) {
	var logins []Querylogin

	err := c.requestWithParams("queryloginlist", request, &logins)
	if err != nil {
		return nil, err
	}

	return &logins, nil
}
