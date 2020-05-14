package go_ts3

// apikeyadd `manage_scope`
type ApiKeyScope string

//noinspection GoUnusedConst
const (
	MANAGE ApiKeyScope = "manage"
	WRITE  ApiKeyScope = "write"
	READ   ApiKeyScope = "read"
)

type ApiKeyAddRequest struct {
	Scope      ApiKeyScope `schema:"scope,required"`
	Lifetime   int         `schema:"lifetime,omitempty"`
	ClientDbId int         `schema:"cldbid,omitempty"`
}

type ApiKeyAddResponse struct {
	Apikey     string `json:"apikey"`
	ClientDbId int    `json:"cldbid,string"`
	CreatedAt  int    `json:"created_at,string"`
	ExpiresAt  int    `json:"expires_at,string"`
	Id         int    `json:"id,string"`
	Scope      string `json:"scope"`
	ServerId   int    `json:"sid,string"`
	TimeLeft   int    `json:"time_left,string"`
}

func (c *TeamspeakHttpClient) ApiKeyAdd(request ApiKeyAddRequest) (*ApiKeyAddResponse, error) {
	var apiKeys []ApiKeyAddResponse
	err := c.requestWithParams("apikeyadd", request, &apiKeys)
	if err != nil {
		return nil, err
	}

	return &apiKeys[0], nil
}

// apikeydel `manage_scope`
type apiKeyDeleteRequest struct {
	Id int `schema:"id,required"`
}

func (c *TeamspeakHttpClient) ApiKeyDelete(id int) error {
	return c.requestWithParams("apikeydel", apiKeyDeleteRequest{Id: id}, nil)
}

// apikeylist `manage_scope`
type ApiKeyListResponse struct {
	ClientDbId int    `json:"cldbid,string"`
	CreatedAt  int    `json:"created_at,string"`
	ExpiresAt  int    `json:"expires_at,string"`
	ID         int    `json:"id,string"`
	Scope      string `json:"scope"`
	ServerId   int    `json:"sid,string"`
	TimeLeft   string `json:"time_left"`
}

type ApiKeyListRequest struct {
	ClientDbId string `schema:"cldbid,omitempty"`
	Start      int    `schema:"start,omitempty"`
	Duration   int    `schema:"duration,omitempty"`
	Count      bool   `schema:"count,omitempty"`
}

func (c *TeamspeakHttpClient) ApiKeyList(request ApiKeyListRequest) (*[]ApiKeyListResponse, error) {
	var apiKeys []ApiKeyListResponse

	err := c.requestWithParams("apikeylist", request, &apiKeys)
	if err != nil {
		return nil, err
	}

	return &apiKeys, nil
}
