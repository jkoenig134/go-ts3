package go_ts3_http

// banadd `manage_scope, write_scope`
type BanAddRequest struct {
	IP           string `schema:"ip,omitempty"`
	Name         string `schema:"name,omitempty"`
	UID          string `schema:"uid,omitempty"`
	MyTsID       string `schema:"mytsid,omitempty"`
	Time         int    `schema:"time,omitempty"`
	BanReason    string `schema:"banreason,omitempty"`
	LastNickname string `schema:"lastnickname,omitempty"`
}

func (c *TeamspeakHttpClient) BanAdd(request BanAddRequest) error {
	return c.requestWithParams(
		"banadd",
		request,
		nil,
	)
}

// banclient `manage_scope, write_scope`
type BanClientRequest struct {
	ClientId  int    `schema:"clid,required"`
	Time      int    `schema:"time,omitempty"`
	BanReason string `schema:"banreason,omitempty"`
}

func (c *TeamspeakHttpClient) BanClient(request BanClientRequest) error {
	return c.requestWithParams(
		"banclient",
		request,
		nil,
	)
}

// bandel `manage_scope, write_scope`
type banDeleteRequest struct {
	BanId int `schema:"banid"`
}

func (c *TeamspeakHttpClient) BanDelete(banId int) error {
	return c.requestWithParams(
		"bandel",
		banDeleteRequest{BanId: banId},
		nil,
	)
}

// bandelall `manage_scope, write_scope`
func (c *TeamspeakHttpClient) BanDeleteAll() error {
	return c.request("bandelall", nil)
}

// banlist `manage_scope, write_scope, read_scope`
type Ban struct {
	BanId             string `json:"banid"`
	Created           string `json:"created"`
	Duration          string `json:"duration"`
	Enforcements      string `json:"enforcements"`
	InvokerClientDbId string `json:"invokercldbid"`
	InvokerName       string `json:"invokername"`
	InvokerUID        string `json:"invokeruid"`
	IP                string `json:"ip"`
	LastNickname      string `json:"lastnickname"`
	Mytsid            string `json:"mytsid"`
	Name              string `json:"name"`
	Reason            string `json:"reason"`
	UID               string `json:"uid"`
}

type BanListRequest struct {
	Start    int `schema:"start"`
	Duration int `schema:"duration"`
}

func (c *TeamspeakHttpClient) BanList(request BanListRequest) (*[]Ban, error) {
	var bans []Ban

	err := c.requestWithParams("banlist", request, &bans)
	if err != nil {
		return nil, err
	}

	return &bans, nil
}
