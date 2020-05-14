package go_ts3

import (
	"fmt"
	"strings"
)

// tokenadd `manage_scope, write_scope`
// tokencustomset is currently not possible
type TokenAddRequest struct {
	TokenType        int    `schema:"tokentype"`
	TokenMajorId     int    `schema:"tokenid1"`
	TokenMinorId     int    `schema:"tokenid2"`
	TokenDescription string `schema:"tokendescription"`
	TokenCustomSet   string `schema:"tokencustomset,omitempty"`
}

func NewGroupToken(serverGroupId int, description string) TokenAddRequest {
	return TokenAddRequest{
		TokenType:        0,
		TokenMajorId:     serverGroupId,
		TokenMinorId:     0,
		TokenDescription: description,
	}
}

func NewCustomSetGroupToken(serverGroupId int, description string, customSet map[string]string) TokenAddRequest {
	return TokenAddRequest{
		TokenType:        0,
		TokenMajorId:     serverGroupId,
		TokenMinorId:     0,
		TokenDescription: description,
		TokenCustomSet:   generateCustomSet(customSet),
	}
}

func NewChannelToken(channelGroupId, channelId int, description string) TokenAddRequest {
	return TokenAddRequest{
		TokenType:        1,
		TokenMajorId:     channelGroupId,
		TokenMinorId:     channelId,
		TokenDescription: description,
	}
}

func NewCustomSetChannelToken(channelGroupId, channelId int, description string, customSet map[string]string) TokenAddRequest {
	return TokenAddRequest{
		TokenType:        1,
		TokenMajorId:     channelGroupId,
		TokenMinorId:     channelId,
		TokenDescription: description,
		TokenCustomSet:   generateCustomSet(customSet),
	}
}

func generateCustomSet(values map[string]string) string {
	var entries []string

	for ident, value := range values {
		entries = append(entries, fmt.Sprintf("ident=%s value=%s", ident, value))
	}

	return strings.Join(entries, "\\p")
}

type TokenAddResponse struct {
	Token string `schema:"token"`
}

func (c *TeamspeakHttpClient) TokenAdd(request TokenAddRequest) (*TokenAddResponse, error) {
	var tokens []TokenAddResponse

	err := c.requestWithParams("tokenadd", request, &tokens)
	if err != nil {
		return nil, err
	}

	return &tokens[0], nil
}

// tokendelete `manage_scope, write_scope`
type tokenDeleteRequest struct {
	TokenKey string `schema:"token,required"`
}

func (c *TeamspeakHttpClient) TokenDelete(tokenKey string) error {
	return c.requestWithParams("tokendelete", tokenDeleteRequest{TokenKey: tokenKey}, nil)
}

// tokenlist `manage_scope, write_scope, read_scope`
type Token struct {
	Token            string `json:"token"`
	TokenDescription string `json:"token_description"`
	TokenCreated     int    `json:"token_created,string"`
	TokenType        int    `json:"token_type,string"`
	TokenId1         int    `json:"token_id1,string"`
	TokenId2         int    `json:"token_id2,string"`
	TokenCustomset   string `json:"token_customset"`
}

func (c *TeamspeakHttpClient) TokenList() (*[]Token, error) {
	var tokens []Token

	err := c.request("tokenlist", &tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

// tokenuse `manage_scope, write_scope`
type tokenUseRequest struct {
	TokenKey string `schema:"token,required"`
}

func (c *TeamspeakHttpClient) TokenUse(tokenKey string) error {
	return c.requestWithParams("tokenuse", tokenUseRequest{TokenKey: tokenKey}, nil)
}
