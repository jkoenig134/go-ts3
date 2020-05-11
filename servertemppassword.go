package go_ts3_http

// servertemppasswordadd `manage_scope, write_scope`
type serverTempPasswordAddRequest struct {
	Password              string `schema:"pw"`
	Description           string `schema:"desc"`
	Duration              int    `schema:"duration"`
	TargetChannelId       int    `schema:"tcid"`
	TargetChannelPassword string `schema:"tcpw"`
}

func (c *TeamspeakHttpClient) ServerTempPasswordAdd(password string, description string, duration int, targetChannelId int, targetChannelPassword string) error {
	return c.requestWithParams(
		"servertemppasswordadd",
		serverTempPasswordAddRequest{
			Password:              password,
			Description:           description,
			Duration:              duration,
			TargetChannelId:       targetChannelId,
			TargetChannelPassword: targetChannelPassword,
		},
		nil,
	)
}

// servertemppassworddel `manage_scope, write_scope`
type serverTempPasswordDeleteRequest struct {
	Password string `schema:"pw"`
}

func (c *TeamspeakHttpClient) ServerTempPasswordDelete(password string) error {
	return c.requestWithParams(
		"servertemppassworddel",
		serverTempPasswordDeleteRequest{Password: password},
		nil,
	)
}

// servertemppasswordlist `manage_scope, write_scope, read_scope`
type ServerTempPassword struct {
	Nickname          string `json:"nickname"`
	UID               string `json:"uid"`
	Description       string `json:"desc"`
	CleartextPassword string `json:"pw_clear"`
	Start             int    `json:"start,string"`
	End               int    `json:"end,string"`
	TargetChannelId   int    `json:"tcid,string"`
}

func (c *TeamspeakHttpClient) ServerTempPasswordList() (*[]ServerTempPassword, error) {
	var passwords []ServerTempPassword

	err := c.request("servertemppasswordlist", passwords)
	if err != nil {
		return nil, err
	}

	return &passwords, nil
}
