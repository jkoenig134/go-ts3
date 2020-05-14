package go_ts3

// channelclientaddperm `manage_scope`
type channelClientAddPermissionRequest struct {
	ChannelId          int    `schema:"cid"`
	ClientDbId         int    `schema:"cldbid"`
	PermissionId       int    `schema:"permid,omitempty"`
	StringPermissionId string `schema:"permsid,omitempty"`
	PermissionValue    int    `schema:"permvalue"`
}

func (c *TeamspeakHttpClient) channelClientAddPermission(request channelClientAddPermissionRequest) error {
	return c.requestWithParams("channelclientaddperm", request, nil)
}

func (c *TeamspeakHttpClient) ChannelClientAddPermission(channelId, clientDbId, permissionId, permissionValue int) error {
	return c.channelClientAddPermission(channelClientAddPermissionRequest{
		ChannelId:       channelId,
		ClientDbId:      clientDbId,
		PermissionId:    permissionId,
		PermissionValue: permissionValue,
	})
}

func (c *TeamspeakHttpClient) ChannelClientAddStringPermission(channelId, clientDbId int, permissionId string, permissionValue int) error {
	return c.channelClientAddPermission(channelClientAddPermissionRequest{
		ChannelId:          channelId,
		ClientDbId:         clientDbId,
		StringPermissionId: permissionId,
		PermissionValue:    permissionValue,
	})
}

// channelclientdelperm `manage_scope`
type channelClientDeletePermissionRequest struct {
	ChannelId          int    `schema:"cid"`
	ClientDbId         int    `schema:"cldbid"`
	PermissionId       int    `schema:"permid,omitempty"`
	StringPermissionId string `schema:"permsid,omitempty"`
}

func (c *TeamspeakHttpClient) channelClientDeletePermission(request channelClientDeletePermissionRequest) error {
	return c.requestWithParams("channelclientdelperm", request, nil)
}

func (c *TeamspeakHttpClient) ChannelClientDeletePermission(channelId, clientDbId, permissionId int) error {
	return c.channelClientDeletePermission(channelClientDeletePermissionRequest{
		ChannelId:    channelId,
		ClientDbId:   clientDbId,
		PermissionId: permissionId,
	})
}

func (c *TeamspeakHttpClient) ChannelClientDeleteStringPermission(channelId, clientDbId int, permissionId string) error {
	return c.channelClientDeletePermission(channelClientDeletePermissionRequest{
		ChannelId:          channelId,
		ClientDbId:         clientDbId,
		StringPermissionId: permissionId,
	})
}

// channelclientpermlist `manage_scope, write_scope, read_scope`
type channelClientPermissionListRequest struct {
	ChannelId  int  `schema:"cid"`
	ClientDbId int  `schema:"cldbid"`
	AsString   bool `schema:"-permsid,omitempty"`
}

type ChannelClientPermissionListResponse struct {
	PermissionId      int `json:"permid,string"`
	PermissionNegated int `json:"permnegated,string"`
	PermissionSkip    int `json:"permskip,string"`
	PermissionValue   int `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ChannelClientPermissionList(channelId, clientDbId int) (*[]ChannelClientPermissionListResponse, error) {
	var perms []ChannelClientPermissionListResponse

	err := c.requestWithParams(
		"channelpermlist",
		channelClientPermissionListRequest{
			ChannelId:  channelId,
			ClientDbId: clientDbId,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

type ChannelClientStringPermissionListResponse struct {
	PermissionId      string `json:"permsid"`
	PermissionNegated int    `json:"permnegated,string"`
	PermissionSkip    int    `json:"permskip,string"`
	PermissionValue   int    `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ChannelClientStringPermissionList(channelId, clientDbId int) (*[]ChannelClientStringPermissionListResponse, error) {
	var perms []ChannelClientStringPermissionListResponse

	err := c.requestWithParams(
		"channelpermlist",
		channelClientPermissionListRequest{
			ChannelId:  channelId,
			ClientDbId: clientDbId,
			AsString:   true,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}
