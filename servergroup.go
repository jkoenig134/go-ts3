package go_ts3

type ServerGroupType int

//noinspection GoUnusedConst
const (
	ServerGroupTypeChannelGuest    ServerGroupType = 10
	ServerGroupTypeServerGuest     ServerGroupType = 15
	ServerGroupTypeQueryGuest      ServerGroupType = 20
	ServerGroupTypeChannelVoice    ServerGroupType = 25
	ServerGroupTypeServerNormal    ServerGroupType = 30
	ServerGroupTypeChannelOperator ServerGroupType = 35
	ServerGroupTypeChannelAdmin    ServerGroupType = 40
	ServerGroupTypeServerAdmin     ServerGroupType = 45
	ServerGroupTypeQueryAdmin      ServerGroupType = 50
)

type GroupDatabaseType int

//noinspection GoUnusedConst
const (
	GroupDatabaseTypeTemplate GroupDatabaseType = 0
	GroupDatabaseTypeRegular  GroupDatabaseType = 1
	GroupDatabaseTypeQuery    GroupDatabaseType = 2
)

// servergroupadd `manage_scope`
type serverGroupAddRequest struct {
	Name              string            `schema:"name,required"`
	GroupDatabaseType GroupDatabaseType `schema:"type,omitempty"`
}

type serverGroupAddResponse struct {
	ServerGroupId int `json:"sgid,string"`
}

func (c *TeamspeakHttpClient) serverGroupAdd(request serverGroupAddRequest) (*int, error) {
	var ids []serverGroupAddResponse

	err := c.requestWithParams(
		"servergroupadd",
		request,
		&ids,
	)
	if err != nil {
		return nil, err
	}

	return &ids[0].ServerGroupId, nil
}

func (c *TeamspeakHttpClient) ServerGroupAdd(name string) (*int, error) {
	return c.serverGroupAdd(serverGroupAddRequest{Name: name})
}

func (c *TeamspeakHttpClient) ServerGroupAddWithType(name string, groupType GroupDatabaseType) (*int, error) {
	return c.serverGroupAdd(serverGroupAddRequest{Name: name, GroupDatabaseType: groupType})
}

// servergroupaddclient `manage_scope`
type serverGroupAddClientRequest struct {
	ServerGroupId int `schema:"sgid,required"`
	ClientDbId    int `schema:"cldbid,required"`
}

func (c *TeamspeakHttpClient) ServerGroupAddClient(serverGroupId, clientDbId int) error {
	return c.requestWithParams(
		"servergroupaddclient",
		serverGroupAddClientRequest{
			ServerGroupId: serverGroupId,
			ClientDbId:    clientDbId,
		},
		nil,
	)
}

// servergroupaddperm `manage_scope`
type serverGroupAddPermissionRequest struct {
	ServerGroupId   int `schema:"sgid"`
	PermissionId    int `schema:"permid"`
	PermissionValue int `schema:"permvalue"`
	PermNegated     int `schema:"permnegated"`
	PermSkip        int `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ServerGroupAddPermission(
	serverGroupId, permissionId, permissionValue int,
	permSkip, permNegated bool,
) error {
	return c.requestWithParams(
		"servergroupaddperm",
		serverGroupAddPermissionRequest{
			ServerGroupId:   serverGroupId,
			PermissionId:    permissionId,
			PermissionValue: permissionValue,
			PermNegated:     boolToInt(permNegated),
			PermSkip:        boolToInt(permSkip),
		},
		nil,
	)
}

type serverGroupAddStringPermissionRequest struct {
	ServerGroupId      int    `schema:"sgid"`
	PermissionStringId string `schema:"permsid"`
	PermissionValue    int    `schema:"permvalue"`
	PermNegated        int    `schema:"permnegated"`
	PermSkip           int    `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ServerGroupAddStringPermission(
	serverGroupId int,
	permissionStringId string,
	permissionValue int,
	permSkip, permNegated bool,
) error {
	return c.requestWithParams(
		"servergroupaddperm",
		serverGroupAddStringPermissionRequest{
			ServerGroupId:      serverGroupId,
			PermissionStringId: permissionStringId,
			PermissionValue:    permissionValue,
			PermNegated:        boolToInt(permNegated),
			PermSkip:           boolToInt(permSkip),
		},
		nil,
	)
}

// servergroupautoaddperm `manage_scope`
type serverGroupAutoAddPermissionRequest struct {
	ServerGroupType ServerGroupType `schema:"sgtype"`
	PermissionId    int             `schema:"permid"`
	PermissionValue int             `schema:"permvalue"`
	PermNegated     int             `schema:"permnegated"`
	PermSkip        int             `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ServerGroupAutoAddPermission(
	serverGroupType ServerGroupType,
	permissionId, permissionValue int,
	permSkip, permNegated bool,
) error {
	return c.requestWithParams(
		"servergroupaddperm",
		serverGroupAutoAddPermissionRequest{
			ServerGroupType: serverGroupType,
			PermissionId:    permissionId,
			PermissionValue: permissionValue,
			PermNegated:     boolToInt(permNegated),
			PermSkip:        boolToInt(permSkip),
		},
		nil,
	)
}

type serverGroupAutoAddStringPermissionRequest struct {
	ServerGroupType    ServerGroupType `schema:"sgtype"`
	PermissionStringId string          `schema:"permsid"`
	PermissionValue    int             `schema:"permvalue"`
	PermNegated        int             `schema:"permnegated"`
	PermSkip           int             `schema:"permskip"`
}

func (c *TeamspeakHttpClient) ServerGroupAutoAddStringPermission(
	serverGroupType ServerGroupType,
	permissionStringId string,
	permissionValue int,
	permSkip, permNegated bool,
) error {
	return c.requestWithParams(
		"servergroupautoaddperm",
		serverGroupAutoAddStringPermissionRequest{
			ServerGroupType:    serverGroupType,
			PermissionStringId: permissionStringId,
			PermissionValue:    permissionValue,
			PermNegated:        boolToInt(permNegated),
			PermSkip:           boolToInt(permSkip),
		},
		nil,
	)
}

// servergroupautodelperm `manage_scope`
type serverGroupAutoDeletePermissionRequest struct {
	ServerGroupType ServerGroupType `schema:"sgtype"`
	PermissionId    int             `schema:"permid"`
}

func (c *TeamspeakHttpClient) ServerGroupAutoDeletePermission(serverGroupType ServerGroupType, permissionId int) error {
	return c.requestWithParams(
		"servergroupautodelperm",
		serverGroupAutoDeletePermissionRequest{
			ServerGroupType: serverGroupType,
			PermissionId:    permissionId,
		},
		nil,
	)
}

type serverGroupAutoDeleteStringPermissionRequest struct {
	ServerGroupType    ServerGroupType `schema:"sgtype"`
	StringPermissionId string          `schema:"permsid"`
}

func (c *TeamspeakHttpClient) ServerGroupAutoDeleteStringPermission(serverGroupType ServerGroupType, stringPermissionId string) error {
	return c.requestWithParams(
		"servergroupautodelperm",
		serverGroupAutoDeleteStringPermissionRequest{
			ServerGroupType:    serverGroupType,
			StringPermissionId: stringPermissionId,
		},
		nil,
	)
}

// servergroupclientlist `manage_scope`
type serverGroupClientListRequest struct {
	ServerGroupId int  `schema:"sgid"`
	Names         bool `schema:"names,omitempty"`
}

type ServerGroupClientList struct {
	ClientDbId             int    `json:"cldbid,string"`
	ClientNickname         string `json:"client_nickname"`
	ClientUniqueIdentifier string `json:"client_unique_identifier"`
}

func (c *TeamspeakHttpClient) ServerGroupClientList(serverGroupId int) (*[]ServerGroupClientList, error) {
	var clients []ServerGroupClientList

	err := c.requestWithParams(
		"servergroupclientlist",
		serverGroupClientListRequest{
			ServerGroupId: serverGroupId,
			Names:         true,
		},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

// servergroupcopy `manage_scope`
type serverGroupCopyRequest struct {
	SourceGroupId int               `schema:"ssgid"`
	TargetGroupId int               `schema:"tsgid"`
	Name          string            `schema:"name"`
	Type          GroupDatabaseType `schema:"type"`
}

type serverGroupCopyResponse struct {
	ServerGroupId int `json:"sgid,string"`
}

func (c *TeamspeakHttpClient) ServerGroupCopy(sourceGroupId, targetGroupId int, name string, groupType GroupDatabaseType) (*int, error) {
	var ids []serverGroupCopyResponse

	err := c.requestWithParams(
		"servergroupcopy",
		serverGroupCopyRequest{
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

	return &ids[0].ServerGroupId, nil
}

// servergroupdel `manage_scope`
type serverGroupDeleteRequest struct {
	ServerGroupId int `schema:"sgid"`
	Force         int `schema:"force"`
}

func (c *TeamspeakHttpClient) ServerGroupDelete(serverGroupId int, force bool) error {
	return c.requestWithParams(
		"servergroupdel",
		serverGroupDeleteRequest{
			ServerGroupId: serverGroupId,
			Force:         boolToInt(force),
		},
		nil,
	)
}

// servergroupdelclient `manage_scope`
type serverGroupDeleteClientRequest struct {
	ServerGroupId int `schema:"sgid,required"`
	ClientDbId    int `schema:"cldbid,required"`
}

func (c *TeamspeakHttpClient) ServerGroupDeleteClient(serverGroupId, clientDbId int) error {
	return c.requestWithParams(
		"servergroupdelclient",
		serverGroupDeleteClientRequest{
			ServerGroupId: serverGroupId,
			ClientDbId:    clientDbId,
		},
		nil,
	)
}

// servergroupdelperm `manage_scope`
type ServerGroupDeletePermissionRequest struct {
	ServerGroupId      int      `schema:"sgid,required"`
	PermissionId       []int    `schema:"permid,omitempty"`
	StringPermissionId []string `schema:"permsid,omitempty"`
}

func (c *TeamspeakHttpClient) ServerGroupDeletePermission(request ServerGroupDeletePermissionRequest) error {
	return c.requestWithParams("servergroupdelperm", request, nil)
}

// servergrouplist `manage_scope`
type ServerGroup struct {
	ServerGroupId  int               `json:"sgid,string"`
	Type           GroupDatabaseType `json:"type,string"`
	Name           string            `json:"name"`
	NameMode       int               `json:"namemode,string"`
	IconId         int               `json:"iconid,string"`
	NMemberAddp    int               `json:"n_member_addp,string"`
	NMemberRemovep int               `json:"n_member_removep,string"`
	NModifyp       int               `json:"n_modifyp,string"`
	SaveDb         int               `json:"savedb,string"`
	SortId         int               `json:"sortid,string"`
}

func (c *TeamspeakHttpClient) ServerGroupList() (*[]ServerGroup, error) {
	var serverGroups []ServerGroup

	err := c.request("servergrouplist", &serverGroups)
	if err != nil {
		return nil, err
	}

	return &serverGroups, nil
}

// servergrouppermlist `manage_scope`
type serverGroupPermsListRequest struct {
	ServerGroupId int  `schema:"sgid"`
	AsString      bool `schema:"-permsid,omitempty"`
}

type ServerGroupPermission struct {
	PermissionId      int `json:"permid,string"`
	PermissionValue   int `json:"permvalue,string"`
	PermissionNegated int `json:"permnegated,string"`
	PermSkip          int `json:"permskip,string"`
}

func (c *TeamspeakHttpClient) ServerGroupPermissionList(serverGroupId int) (*[]ServerGroupPermission, error) {
	var perms []ServerGroupPermission

	err := c.requestWithParams(
		"servergrouppermlist",
		serverGroupPermsListRequest{
			ServerGroupId: serverGroupId,
			AsString:      false,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

type ServerGroupStringPermission struct {
	StringPermissionId string `json:"permid"`
	PermissionValue    int    `json:"permvalue,string"`
	PermissionNegated  int    `json:"permnegated,string"`
	PermSkip           int    `json:"permskip,string"`
}

func (c *TeamspeakHttpClient) ServerGroupStringPermissionList(serverGroupId int) (*[]ServerGroupStringPermission, error) {
	var perms []ServerGroupStringPermission

	err := c.requestWithParams(
		"servergrouppermlist",
		serverGroupPermsListRequest{
			ServerGroupId: serverGroupId,
			AsString:      true,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &perms, nil
}

// servergrouprename `manage_scope`
type serverGroupRenameRequest struct {
	ServerGroupId int    `schema:"sgid"`
	Name          string `schema:"name"`
}

func (c *TeamspeakHttpClient) ServerGroupRename(serverGroupId int, name string) error {
	return c.requestWithParams(
		"servergrouprename",
		serverGroupRenameRequest{
			ServerGroupId: serverGroupId,
			Name:          name,
		},
		nil,
	)
}

// servergroupsbyclientid `manage_scope`
type serverGroupsByClientIdRequest struct {
	ClientDbId int `schema:"cldbid"`
}

type ServerGroupsByClientIdResponse struct {
	GroupName string `json:"name"`
	GroupId   int    `json:"sgid,string"`
}

func (c *TeamspeakHttpClient) ServerGroupsByClientId(clientDbId int) (*[]ServerGroupsByClientIdResponse, error) {
	var groups []ServerGroupsByClientIdResponse

	err := c.requestWithParams(
		"servergroupsbyclientid",
		serverGroupsByClientIdRequest{ClientDbId: clientDbId},
		&groups,
	)
	if err != nil {
		return nil, err
	}

	return &groups, nil
}
