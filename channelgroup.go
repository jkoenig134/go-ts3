package go_ts3

// channelgroupadd `manage_scope, write_scope`
type channelGroupAddRequest struct {
	Name              string            `schema:"name,required"`
	GroupDatabaseType GroupDatabaseType `schema:"type,omitempty"`
}

type channelGroupAddResponse struct {
	ChannelGroupId int `json:"cgid,string"`
}

func (c *TeamspeakHttpClient) channelGroupAdd(request channelGroupAddRequest) (*int, error) {
	var ids []channelGroupAddResponse

	err := c.requestWithParams(
		"channelgroupadd",
		request,
		&ids,
	)
	if err != nil {
		return nil, err
	}

	return &ids[0].ChannelGroupId, nil
}

func (c *TeamspeakHttpClient) ChannelGroupAdd(name string) (*int, error) {
	return c.channelGroupAdd(channelGroupAddRequest{Name: name})
}

func (c *TeamspeakHttpClient) ChannelGroupAddWithType(name string, groupType GroupDatabaseType) (*int, error) {
	return c.channelGroupAdd(channelGroupAddRequest{Name: name, GroupDatabaseType: groupType})
}

// channelgroupaddperm `manage_scope`
type channelGroupAddPermissionRequest struct {
	ChannelGroupId  int `schema:"cgid"`
	PermissionId    int `schema:"permid"`
	PermissionValue int `schema:"permvalue"`
	PermNegated     int `schema:"permnegated"`
	PermSkip        int `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ChannelGroupAddPermission(
	channelGroupId, permissionId, permissionValue int,
	permSkip, permNegated bool,
) error {
	return c.requestWithParams(
		"channelgroupaddperm",
		channelGroupAddPermissionRequest{
			ChannelGroupId:  channelGroupId,
			PermissionId:    permissionId,
			PermissionValue: permissionValue,
			PermNegated:     boolToInt(permNegated),
			PermSkip:        boolToInt(permSkip),
		},
		nil,
	)
}

type channelGroupAddStringPermissionRequest struct {
	ChannelGroupId     int    `schema:"cgid"`
	PermissionStringId string `schema:"permsid"`
	PermissionValue    int    `schema:"permvalue"`
	PermNegated        int    `schema:"permnegated"`
	PermSkip           int    `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ChannelGroupAddStringPermission(
	channelGroupId int,
	permissionStringId string,
	permissionValue int,
	permSkip, permNegated bool,
) error {
	return c.requestWithParams(
		"channelgroupaddperm",
		channelGroupAddStringPermissionRequest{
			ChannelGroupId:     channelGroupId,
			PermissionStringId: permissionStringId,
			PermissionValue:    permissionValue,
			PermNegated:        boolToInt(permNegated),
			PermSkip:           boolToInt(permSkip),
		},
		nil,
	)
}

// channelgroupclientlist `manage_scope, write_scope, read_scope`
type ChannelGroupClientListRequest struct {
	ChannelId  int `schema:"cid,omitempty"`
	ClientDbId int `schema:"cldbid,omitempty"`
	GroupId    int `schema:"cgid,omitempty"`
}

type ChannelGroupClientListResponse struct {
	ChannelGroupId int `json:"cgid,string"`
	ChannelId      int `json:"cid,string"`
	ClientDbId     int `json:"cldbid,string"`
}

func (c *TeamspeakHttpClient) ChannelGroupClientList(request ChannelGroupClientListRequest) (*[]ChannelGroupClientListResponse, error) {
	var clients []ChannelGroupClientListResponse

	err := c.requestWithParams("channelgroupclientlist", request, &clients)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

// channelgroupcopy `manage_scope`
type channelGroupCopyRequest struct {
	SourceGroupId int
	TargetGroupId int
	Name          string
	Type          GroupDatabaseType
}

type channelGroupCopyResponse struct {
	ChannelGroupId int `json:"cgid,string"`
}

func (c *TeamspeakHttpClient) ChannelGroupCopy(sourceGroupId, targetGroupId int, name string, groupType GroupDatabaseType) (*int, error) {
	var ids []channelGroupCopyResponse

	err := c.requestWithParams(
		"channelgroupcopy",
		channelGroupCopyRequest{
			SourceGroupId: sourceGroupId,
			TargetGroupId: targetGroupId,
			Name:          name,
			Type:          groupType,
		},
		&ids,
	)
	if err != nil {
		return nil, err
	}

	return &ids[0].ChannelGroupId, nil
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
	Name           string            `json:"name"`
	ChannelGroupId int               `json:"cgid,string"`
	Type           GroupDatabaseType `json:"type,string"`
	IconId         int               `json:"iconid,string"`
	Namemode       int               `json:"namemode,string"`
	Savedb         int               `json:"savedb,string"`
	Sortid         int               `json:"sortid,string"`
	NMemberAddp    int               `json:"n_member_addp,string"`
	NMemberRemovep int               `json:"n_member_removep,string"`
	NModifyp       int               `json:"n_modifyp,string"`
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
