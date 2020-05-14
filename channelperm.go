package go_ts3

// channeladdperm `manage_scope`
type channelAddPermissionRequest struct {
	ChannelId          int    `schema:"cid"`
	PermissionId       int    `schema:"permid,omitempty"`
	StringPermissionId string `schema:"permsid,omitempty"`
	PermissionValue    int    `schema:"permvalue"`
}

func (c *TeamspeakHttpClient) channelAddPermission(request channelAddPermissionRequest) error {
	return c.requestWithParams("channeladdperm", request, nil)
}

func (c *TeamspeakHttpClient) ChannelAddPermission(channelId int, permissionId int, permissionValue int) error {
	return c.channelAddPermission(channelAddPermissionRequest{
		ChannelId:       channelId,
		PermissionId:    permissionId,
		PermissionValue: permissionValue,
	})
}

func (c *TeamspeakHttpClient) ChannelAddStringPermission(channelId int, permissionId string, permissionValue int) error {
	return c.channelAddPermission(channelAddPermissionRequest{
		ChannelId:          channelId,
		StringPermissionId: permissionId,
		PermissionValue:    permissionValue,
	})
}

// channeldelperm `manage_scope`
type channelDeletePermissionRequest struct {
	ChannelId          int    `schema:"cid"`
	PermissionId       int    `schema:"permid,omitempty"`
	StringPermissionId string `schema:"permsid,omitempty"`
}

func (c *TeamspeakHttpClient) channelDeletePermission(request channelDeletePermissionRequest) error {
	return c.requestWithParams("channeldelperm", request, nil)
}

func (c *TeamspeakHttpClient) ChannelDeletePermission(channelId, permissionId int) error {
	return c.channelDeletePermission(channelDeletePermissionRequest{
		ChannelId:    channelId,
		PermissionId: permissionId,
	})
}

func (c *TeamspeakHttpClient) ChannelDeleteStringPermission(channelId int, permissionId string) error {
	return c.channelDeletePermission(channelDeletePermissionRequest{
		ChannelId:          channelId,
		StringPermissionId: permissionId,
	})
}

// channelpermlist `manage_scope, write_scope, read_scope`
type channelPermissionListRequest struct {
	ChannelId int  `schema:"cid"`
	AsString  bool `schema:"-permsid,omitempty"`
}

type ChannelPermissionListResponse struct {
	PermissionId      int `json:"permid,string"`
	PermissionNegated int `json:"permnegated,string"`
	PermissionSkip    int `json:"permskip,string"`
	PermissionValue   int `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ChannelPermissionList(channelId int) (*[]ChannelPermissionListResponse, error) {
	var perms []ChannelPermissionListResponse

	err := c.requestWithParams(
		"channelpermlist",
		channelPermissionListRequest{
			ChannelId: channelId,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

type ChannelStringPermissionListResponse struct {
	PermissionId      string `json:"permsid"`
	PermissionNegated int    `json:"permnegated,string"`
	PermissionSkip    int    `json:"permskip,string"`
	PermissionValue   int    `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ChannelStringPermissionList(channelId int) (*[]ChannelStringPermissionListResponse, error) {
	var perms []ChannelStringPermissionListResponse

	err := c.requestWithParams(
		"channelpermlist",
		channelPermissionListRequest{
			ChannelId: channelId,
			AsString:  true,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}
