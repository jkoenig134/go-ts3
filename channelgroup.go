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
func (c *TeamspeakHttpClient) ChannelGroupDelete() error {
	//TODO
	return nil
}

// channelgroupdelperm `manage_scope`
func (c *TeamspeakHttpClient) ChannelGroupDeletePermission() error {
	//TODO
	return nil
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
