package go_ts3_http

type gmRequest struct {
	Text string `schema:"msg,required"`
}

// gm `manage_scope`wsl
func (c *TeamspeakHttpClient) GlobalMessage(text string) error {
	return c.requestWithParams("gm", gmRequest{Text: text}, nil)
}
