package go_ts3_http

import "errors"

// clientaddperm `manage_scope`
func (c *TeamspeakHttpClient) ClientAddPermission() error {
	//TODO
	return nil
}

// clientdbdelete `manage_scope, write_scope`
type clientDbDeleteRequest struct {
	ClientDbId int `schema:"cldbid,required"`
}

func (c *TeamspeakHttpClient) ClientDbDelete(clientDbId int) error {
	return c.requestWithParams("clientdbdelete", clientDbDeleteRequest{ClientDbId: clientDbId}, nil)
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
type clientDbInfoRequest struct {
	ClientDbId int `schema:"cldbid,required"`
}

type ClientDbInfo struct {
	ClientBase64HashClientUID  string `json:"client_base64HashClientUID"`
	ClientCreated              int    `json:"client_created,string"`
	ClientDatabaseID           int    `json:"client_database_id,string"`
	ClientDescription          string `json:"client_description"`
	ClientFlagAvatar           string `json:"client_flag_avatar"`
	ClientLastconnected        int    `json:"client_lastconnected,string"`
	ClientLastip               string `json:"client_lastip"`
	ClientMonthBytesDownloaded int    `json:"client_month_bytes_downloaded,string"`
	ClientMonthBytesUploaded   int    `json:"client_month_bytes_uploaded,string"`
	ClientNickname             string `json:"client_nickname"`
	ClientTotalBytesDownloaded int    `json:"client_total_bytes_downloaded,string"`
	ClientTotalBytesUploaded   int    `json:"client_total_bytes_uploaded,string"`
	ClientTotalconnections     int    `json:"client_totalconnections,string"`
	ClientUniqueIdentifier     string `json:"client_unique_identifier"`
}

func (c *TeamspeakHttpClient) ClientDbInfo(clientDbId int) (*ClientDbInfo, error) {
	var clientDbInfos []ClientDbInfo

	err := c.requestWithParams(
		"clientdbinfo",
		clientDbInfoRequest{ClientDbId: clientDbId},
		&clientDbInfos,
	)
	if err != nil {
		return nil, err
	}

	return &clientDbInfos[0], nil
}

// clientdblist `manage_scope, write_scope, read_scope`
type ClientDbListRequest struct {
	Start    int `schema:"start,omitempty"`
	Duration int `schema:"duration,omitempty"`
}

type DbClient struct {
	Cldbid                 int    `json:"cldbid,string"`
	ClientCreated          int    `json:"client_created,string"`
	ClientDescription      string `json:"client_description"`
	ClientLastConnected    int    `json:"client_lastconnected,string"`
	ClientLastIP           string `json:"client_lastip"`
	ClientLoginName        string `json:"client_login_name"`
	ClientNickname         string `json:"client_nickname"`
	ClientTotalConnections int    `json:"client_totalconnections,string"`
	ClientUniqueIdentifier string `json:"client_unique_identifier"`
}

func (c *TeamspeakHttpClient) ClientDbList(request ClientDbListRequest) (*[]DbClient, error) {
	var clients []DbClient

	err := c.requestWithParams("clientdblist", request, &clients)
	if err != nil {
		return nil, err
	}

	return &clients, nil
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
type clientFindRequest struct {
	Pattern string `schema:"pattern,required"`
}

type ClientFindResponse struct {
	ClientId       int    `json:"clid,string"`
	ClientNickname string `json:"client_nickname"`
}

func (c *TeamspeakHttpClient) ClientFind(pattern string) (*[]ClientFindResponse, error) {
	var clients []ClientFindResponse

	err := c.requestWithParams("clientfind", clientFindRequest{Pattern: pattern}, clients)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

// clientgetdbidfromuid `manage_scope, write_scope, read_scope`
type clientGetDbIdFromUidRequest struct {
	ClientUid string `schema:"cluid,required"`
}

type clientGetDbIdFromUidResponse struct {
	ClientDbId int `json:"cldbid,string"`
}

func (c *TeamspeakHttpClient) ClientGetDbIdFromUid(clientUid string) (*int, error) {
	var clients []clientGetDbIdFromUidResponse

	err := c.requestWithParams(
		"clientgetdbidfromuid",
		clientGetDbIdFromUidRequest{ClientUid: clientUid},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0].ClientDbId, nil
}

// clientgetids `manage_scope, write_scope, read_scope`
type clientGetIdsRequest struct {
	ClientUid string `schema:"cluid,required"`
}

type ClientGetIdsResponse struct {
	ClientUid string `json:"cluid"`
	ClientId  int    `json:"clid,string"`
	Name      string `json:"name"`
}

func (c *TeamspeakHttpClient) ClientGetIds(clientUid string) (*ClientGetIdsResponse, error) {
	var clients []ClientGetIdsResponse

	err := c.requestWithParams(
		"clientgetids",
		clientGetIdsRequest{ClientUid: clientUid},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientgetnamefromdbid `manage_scope, write_scope, read_scope`
type clientGetNameFromDbIdRequest struct {
	ClientDbId int `schema:"cldbid,required"`
}

type clientGetNameFromDbIdResponse struct {
	ClientUid  string `json:"cluid"`
	ClientDbId int    `json:"cldbid,string"`
	Name       string `json:"name"`
}

func (c *TeamspeakHttpClient) ClientGetNameFromDbId(clientDbId int) (*clientGetNameFromDbIdResponse, error) {
	var clients []clientGetNameFromDbIdResponse

	err := c.requestWithParams(
		"clientgetnamefromdbid",
		clientGetNameFromDbIdRequest{ClientDbId: clientDbId},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientgetnamefromuid `manage_scope, write_scope, read_scope`
type clientGetNameFromUidRequest struct {
	ClientUid string `schema:"cldbid,required"`
}

type clientGetNameFromUidResponse struct {
	ClientUid  string `json:"cluid"`
	ClientDbId int    `json:"cldbid,string"`
	Name       string `json:"name"`
}

func (c *TeamspeakHttpClient) ClientGetNameFromUid(clientUid string) (*clientGetNameFromUidResponse, error) {
	var clients []clientGetNameFromUidResponse

	err := c.requestWithParams(
		"clientgetnamefromdbid",
		clientGetNameFromUidRequest{ClientUid: clientUid},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientgetuidfromclid `manage_scope, write_scope, read_scope`
type clientGetUidFromClientIdRequest struct {
	ClientId int `schema:"clid,required"`
}

type clientGetUidFromClientIdResponse struct {
	ClientId  int    `json:"clid,string"`
	ClientUid string `json:"cluid"`
	NickName  string `json:"nickname"`
}

func (c *TeamspeakHttpClient) ClientGetUidFromClientId(clientId int) (*clientGetUidFromClientIdResponse, error) {
	var clients []clientGetUidFromClientIdResponse

	err := c.requestWithParams(
		"clientgetuidfromclid",
		clientGetUidFromClientIdRequest{ClientId: clientId},
		&clients,
	)
	if err != nil {
		return nil, err
	}

	return &clients[0], nil
}

// clientinfo `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientInfo() error {
	//TODO
	return nil
}

// clientkick `manage_scope, write_scope`
type ClientKickReason int

//noinspection GoUnusedConst
const (
	BLI ClientKickReason = 4
	BLA ClientKickReason = 5
)

type ClientKickRequest struct {
	ClientId      []int            `schema:"clid,required"`
	ReasonId      ClientKickReason `schema:"reasonid,required"`
	ReasonMessage string           `schema:"reasonmsg,omitempty"`
}

func (c *TeamspeakHttpClient) ClientKick(request ClientKickRequest) error {
	if len(request.ReasonMessage) > 40 {
		return errors.New("the message must not have more than 40 characters")
	}

	return c.requestWithParams("clientkick", request, nil)
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
type ClientMoveRequest struct {
	ClientId        []int  `schema:"clid,required"`
	ChannelId       int    `schema:"cid,required"`
	ChannelPassword string `schema:"cpw,omitempty"`
}

func (c *TeamspeakHttpClient) ClientMove(request ClientMoveRequest) error {
	return c.requestWithParams("clientmove", request, nil)
}

// clientpermlist `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) ClientPermissionList() error {
	//TODO
	return nil
}

// clientpoke `manage_scope, write_scope`
type clientPokeRequest struct {
	ClientId int    `schema:"clid,required"`
	Message  string `schema:"msg,required"`
}

func (c *TeamspeakHttpClient) ClientPoke(clientId int, message string) error {
	return c.requestWithParams("clientpoke", clientPokeRequest{ClientId: clientId, Message: message}, nil)
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
