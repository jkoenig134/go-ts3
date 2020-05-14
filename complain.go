package go_ts3

// complainadd `manage_scope, write_scope`
type complainAddRequest struct {
	TargetClientDbId int    `schema:"tcldbid"`
	Message          string `schema:"message"`
}

func (c *TeamspeakHttpClient) ComplainAdd(targetClientDbId int, message string) error {
	return c.requestWithParams(
		"complainadd",
		complainAddRequest{
			TargetClientDbId: targetClientDbId,
			Message:          message,
		},
		nil,
	)
}

// complaindel `manage_scope, write_scope`
type complainDeleteRequest struct {
	TargetClientDbId int `schema:"tcldbid"`
	SenderClientDbId int `schema:"fcldbid"`
}

func (c *TeamspeakHttpClient) ComplainDelete(targetClientDbId int, senderClientDbId int) error {
	return c.requestWithParams(
		"complaindel",
		complainDeleteRequest{
			TargetClientDbId: targetClientDbId,
			SenderClientDbId: senderClientDbId,
		},
		nil,
	)
}

// complaindelall `manage_scope, write_scope`
type complainDeleteAllRequest struct {
	TargetClientDbId int `schema:"tcldbid"`
}

func (c *TeamspeakHttpClient) ComplainDeleteAll(targetClientDbId int) error {
	return c.requestWithParams(
		"complaindelall",
		complainDeleteAllRequest{TargetClientDbId: targetClientDbId},
		nil,
	)
}

// complainlist `manage_scope, write_scope, read_scope`
type Complain struct {
	SenderClientDbId string `json:"fcldbid"`
	SenderName       string `json:"fname"`
	Message          string `json:"message"`
	Timestamp        string `json:"timestamp"`
	TargetClientDbId string `json:"tcldbid"`
	TargetName       string `json:"tname"`
}

func (c *TeamspeakHttpClient) ComplainList() (*[]Complain, error) {
	var complains []Complain

	err := c.request("complainlist", &complains)
	if err != nil {
		return nil, err
	}

	return &complains, err
}

type complainListByTargetRequest struct {
	TargetClientDbId int `schema:"tcldbid"`
}

func (c *TeamspeakHttpClient) ComplainListByTarget(targetClientDbId int) (*[]Complain, error) {
	var complains []Complain

	err := c.requestWithParams(
		"complainlist",
		complainListByTargetRequest{TargetClientDbId: targetClientDbId},
		&complains,
	)
	if err != nil {
		return nil, err
	}

	return &complains, err
}
