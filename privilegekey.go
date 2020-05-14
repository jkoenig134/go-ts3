package go_ts3

// privilegekeyadd `manage_scope, write_scope`
func (c *TeamspeakHttpClient) PrivilegekeyAdd(request TokenAddRequest) (*TokenAddResponse, error) {
	return c.TokenAdd(request)
}

// privilegekeydelete `manage_scope, write_scope`
func (c *TeamspeakHttpClient) PrivilegekeyDelete(tokenKey string) error {
	return c.TokenDelete(tokenKey)
}

// privilegekeylist `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) PrivilegekeyList() (*[]Token, error) {
	return c.TokenList()
}

// privilegekeyuse `manage_scope, write_scope`
func (c *TeamspeakHttpClient) PrivilegekeyUse(tokenKey string) error {
	return c.TokenUse(tokenKey)
}
