package go_ts3

// channelgroupadd `manage_scope, write_scope`
func (c *TeamspeakHttpClient) ChannelGroupAdd() error {
	//TODO
	return nil
}

// channelgroupaddperm `manage_scope`
func (c *TeamspeakHttpClient) ChannelGroupAddPerm() error {
	//TODO
	return nil
}

// channelgroupclientlist `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ChannelGroupClientList() error {
	//TODO
	return nil
}

// channelgroupcopy `manage_scope`
func (c *TeamspeakHttpClient) ChannelGroupCopy() error {
	//TODO
	return nil
}

// channelgroupdel `manage_scope`
type channelGroupDeleteRequest struct {
	ChannelGroupId int `schema:"cgid"`
	Force          int `schema:"force"`
}

func (c *TeamspeakHttpClient) ChannelGroupDelete(channelGroupId int, force bool) error {
	return c.requestWithParams("channelgroupdel", channelGroupDeleteRequest{
		ChannelGroupId: channelGroupId,
		Force:          boolToInt(force),
	}, nil)
}

// channelgroupdelperm `manage_scope`
type channelGroupDeletePermissionRequest struct {
	ChannelGroupId     int    `schema:"cgid"`
	PermissionId       int    `schema:"permid,omitempty"`
	StringPermissionId string `schema:"permsid,omitempty"`
}

func (c *TeamspeakHttpClient) channelGroupDeletePermission(request channelGroupDeletePermissionRequest) error {
	return c.requestWithParams("channelgroupdelperm", request, nil)
}

func (c *TeamspeakHttpClient) ChannelGroupDeletePermission(channelGroupId, permissionId int) error {
	return c.channelGroupDeletePermission(channelGroupDeletePermissionRequest{
		ChannelGroupId: channelGroupId,
		PermissionId:   permissionId,
	})
}

func (c *TeamspeakHttpClient) ChannelGroupDeleteStringPermission(channelGroupId int, permissionId string) error {
	return c.channelGroupDeletePermission(channelGroupDeletePermissionRequest{
		ChannelGroupId:     channelGroupId,
		StringPermissionId: permissionId,
	})
}

// channelgrouplist `manage_scope, write_scope, read_scope`
type ChannelGroup struct {
	Name           string                      `json:"name"`
	ChannelGroupId int                         `json:"cgid,string"`
	Type           PermissionGroupDatabaseType `json:"type,string"`
	IconId         int                         `json:"iconid,string"`
	Namemode       int                         `json:"namemode,string"`
	Savedb         int                         `json:"savedb,string"`
	Sortid         int                         `json:"sortid,string"`
	NMemberAddp    int                         `json:"n_member_addp,string"`
	NMemberRemovep int                         `json:"n_member_removep,string"`
	NModifyp       int                         `json:"n_modifyp,string"`
}

func (c *TeamspeakHttpClient) ChannelGroupList() (*[]ChannelGroup, error) {
	var groups []ChannelGroup

	err := c.request("channelgrouplist", &groups)
	if err != nil {
		return nil, err
	}

	return &groups, err
}

// channelgrouppermlist `manage_scope, write_scope, read_scope`
type channelGroupPermissionListRequest struct {
	ChannelGroupId int  `schema:"cgid"`
	AsString       bool `schema:"-permsid,omitempty"`
}

type ChannelGroupPermissionListResponse struct {
	PermissionId      int `json:"permid"`
	PermissionNegated int `json:"permnegated,string"`
	PermissionSkip    int `json:"permskip,string"`
	PermissionValue   int `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ChannelGroupPermissionList(channelGroupId int) (*[]ChannelGroupPermissionListResponse, error) {
	var perms []ChannelGroupPermissionListResponse

	err := c.requestWithParams(
		"channelgrouppermlist",
		channelGroupPermissionListRequest{
			ChannelGroupId: channelGroupId,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

type ChannelGroupStringPermissionListResponse struct {
	PermissionId      string `json:"permsid"`
	PermissionNegated int    `json:"permnegated,string"`
	PermissionSkip    int    `json:"permskip,string"`
	PermissionValue   int    `json:"permvalue,string"`
}

func (c *TeamspeakHttpClient) ChannelGroupStringPermissionList(channelGroupId int) (*[]ChannelGroupStringPermissionListResponse, error) {
	var perms []ChannelGroupStringPermissionListResponse

	err := c.requestWithParams(
		"channelgrouppermlist",
		channelGroupPermissionListRequest{
			ChannelGroupId: channelGroupId,
			AsString:       true,
		},
		&perms,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

// channelgrouprename `manage_scope`
type channelGroupRenameRequest struct {
	ChannelGroupId int    `schema:"cgid"`
	Name           string `schema:"name"`
}

func (c *TeamspeakHttpClient) ChannelGroupRename(channelGroupId int, name string) error {
	return c.requestWithParams(
		"channelgrouprename",
		channelGroupRenameRequest{ChannelGroupId: channelGroupId, Name: name},
		nil,
	)
}
