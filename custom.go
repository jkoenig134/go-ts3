package go_ts3

// custominfo `manage_scope, write_scope, read_scope`
type customInfoRequest struct {
	ClientDbId int `schema:"cldbid"`
}

type CustomInfo struct {
	Ident string `json:"ident"`
	Value string `json:"value"`
}

func (c *TeamspeakHttpClient) CustomInfo(clientDbId int) (*[]CustomInfo, error) {
	var customs []CustomInfo

	err := c.requestWithParams("custominfo", customInfoRequest{ClientDbId: clientDbId}, &customs)
	if err != nil {
		return nil, err
	}

	return &customs, nil
}

// customsearch `manage_scope, write_scope, read_scope`
type customSearchRequest struct {
	Ident   string `schema:"ident"`
	Pattern string `schema:"pattern"`
}

type CustomSearchResponse struct {
	ClientDbId int    `json:"cldbid,string"`
	Ident      string `json:"ident"`
	Value      string `json:"value"`
}

func (c *TeamspeakHttpClient) CustomSearch(ident, pattern string) (*[]CustomSearchResponse, error) {
	var responses []CustomSearchResponse

	err := c.requestWithParams(
		"",
		customSearchRequest{
			Ident:   ident,
			Pattern: pattern,
		},
		&responses,
	)
	if err != nil {
		return nil, err
	}

	return &responses, nil
}

// customset `manage_scope, write_scope`
type customSetRequest struct {
	ClientDbId int    `schema:"cldbid"`
	Ident      string `schema:"ident"`
	Value      string `schema:"value"`
}

func (c *TeamspeakHttpClient) CustomSet(clientDbId int, ident, value string) error {
	return c.requestWithParams(
		"customset",
		customSetRequest{
			ClientDbId: clientDbId,
			Ident:      ident,
			Value:      value,
		},
		nil,
	)
}

// customdelete `manage_scope, write_scope`
type customDeleteRequest struct {
	ClientDbId int    `schema:"cldbid"`
	Ident      string `schema:"ident"`
}

func (c *TeamspeakHttpClient) CustomDelete(clientDbId int, ident string) error {
	return c.requestWithParams(
		"customdelete",
		customDeleteRequest{
			ClientDbId: clientDbId,
			Ident:      ident,
		},
		nil,
	)
}
