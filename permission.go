package go_ts3_http

// permfind `manage_scope, write_scope, read_scope`
type permissionFindRequest struct {
	PermissionId       int    `schema:"permid,omitempty"`
	StringPermissionId string `schema:"permsid,omitempty"`
}

type PermissionFindResponse struct {
	Id                  int                 `json:"p,string"`
	PermissionGroupType PermissionGroupType `json:"t,string"`
	MajorId             int                 `json:"id1,string"`
	MinorId             int                 `json:"id2,string"`
}

func (c *TeamspeakHttpClient) permissionFind(request permissionFindRequest) (*PermissionFindResponse, error) {
	var perms []PermissionFindResponse

	err := c.requestWithParams("permfind", request, &perms)
	if err != nil {
		return nil, err
	}

	return &perms[0], nil
}

func (c *TeamspeakHttpClient) PermissionFind(permissionId int) (*PermissionFindResponse, error) {
	return c.permissionFind(permissionFindRequest{PermissionId: permissionId})
}

func (c *TeamspeakHttpClient) PermissionStringFind(permissionId string) (*PermissionFindResponse, error) {
	return c.permissionFind(permissionFindRequest{StringPermissionId: permissionId})
}

// permget `manage_scope, write_scope, read_scope`
type permissionGetRequest struct {
	PermissionId       int    `schema:"permid,omitempty"`
	StringPermissionId string `schema:"permsid,omitempty"`
}

type PermissionGetResponse struct {
	PermissionId       int    `schema:"permid"`
	StringPermissionId string `schema:"permsid"`
	PermissionValue    int    `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) permissionGet(request permissionGetRequest) (*PermissionGetResponse, error) {
	var perms []PermissionGetResponse

	err := c.requestWithParams("permget", request, &perms)
	if err != nil {
		return nil, err
	}

	return &perms[0], nil
}

func (c *TeamspeakHttpClient) PermissionGet(permissionId int) (*PermissionGetResponse, error) {
	return c.permissionGet(permissionGetRequest{PermissionId: permissionId})
}

func (c *TeamspeakHttpClient) StringPermissionGet(permissionId string) (*PermissionGetResponse, error) {
	return c.permissionGet(permissionGetRequest{StringPermissionId: permissionId})
}

// permidgetbyname `manage_scope, write_scope, read_scope`
type permissionGetByNameRequest struct {
	Name string `schema:"permsid"`
}

type permissionGetByNameResponse struct {
	Id int `json:"permid,string"`
}

func (c *TeamspeakHttpClient) PermissionGetByName(name string) (*int, error) {
	var ids []permissionGetByNameResponse

	err := c.requestWithParams("permidgetbyname", permissionGetByNameRequest{Name: name}, &ids)
	if err != nil {
		return nil, err
	}

	return &ids[0].Id, nil
}

// permissionlist `manage_scope, write_scope, read_scope`
type Permission struct {
	PermissionId          int    `json:"permid,string"`
	PermissionName        string `json:"permname"`
	PermissionDescription string `json:"permdesc"`
}

func (c *TeamspeakHttpClient) PermissionList() (*[]Permission, error) {
	var permissions []Permission

	err := c.request("permissionlist", &permissions)
	if err != nil {
		return nil, err
	}

	return &permissions, nil
}

// permoverview `manage_scope, write_scope, read_scope`
type permissionOverviewRequest struct {
	ChannelId    int `schema:"cid"`
	ClientDbId   int `schema:"cldbid"`
	PermissionId int `schema:"permid"`
}

type PermissionOverview struct {
	Id                  int                 `json:"p,string"`
	PermissionGroupType PermissionGroupType `json:"t,string"`
	MajorId             int                 `json:"id1,string"`
	MinorId             int                 `json:"id2,string"`
	Value               int                 `json:"v,string"`
	Negated             int                 `json:"n,string"`
	Skipped             int                 `json:"s,string"`
}

func (c *TeamspeakHttpClient) PermissionOverview(channelId, clientDbId int) (*[]PermissionOverview, error) {
	var permissions []PermissionOverview

	err := c.requestWithParams(
		"permoverview",
		permissionOverviewRequest{
			ChannelId:    channelId,
			ClientDbId:   clientDbId,
			PermissionId: 0,
		},
		&permissions,
	)
	if err != nil {
		return nil, err
	}

	return &permissions, nil
}

// permreset `manage_scope`
type permissionResetResponse struct {
	Token string `json:"token"`
}

func (c *TeamspeakHttpClient) PermissionReset() (*string, error) {
	var tokens []permissionResetResponse

	err := c.request("permreset", &tokens)
	if err != nil {
		return nil, err
	}

	return &tokens[0].Token, nil
}
