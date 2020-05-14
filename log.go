package go_ts3

// logadd `manage_scope, write_scope`
type TeamspeakLogLevel int

//noinspection GoUnusedConst
const (
	FATAL   TeamspeakLogLevel = 1
	ERROR   TeamspeakLogLevel = 2
	WARNING TeamspeakLogLevel = 3
	INFO    TeamspeakLogLevel = 4
)

type LogAddRequest struct {
	LogLevel   TeamspeakLogLevel `schema:"loglevel,required"`
	LogMessage string            `schema:"logmsg,required"`
}

func (c *TeamspeakHttpClient) LogAdd(request LogAddRequest) error {
	return c.requestWithParams(
		"logadd",
		request,
		nil,
	)
}

// logview `manage_scope, write_scope, read_scope`
type Log struct {
	FileSize     int    `json:"file_size,string"`
	LogMessage   string `json:"l"`
	LastPosition int    `json:"last_pos,string"`
}

type logViewRequest struct {
	Lines    int `schema:"lines,omitempty"`
	Reverse  int `schema:"reverse,omitempty"`
	Instance int `schema:"instance,omitempty"`
	BeginPos int `schema:"begin_pos,omitempty"`
}

func (c *TeamspeakHttpClient) logView(request logViewRequest) (*[]Log, error) {
	var logs []Log

	err := c.requestWithParams(
		"logview",
		request,
		&logs,
	)

	if err != nil {
		return nil, err
	}

	return &logs, nil
}

func (c *TeamspeakHttpClient) LogViewInstance(lines int, reverse bool, beginPos int) (*[]Log, error) {
	return c.logView(logViewRequest{
		Lines:    lines,
		Reverse:  boolToInt(reverse),
		Instance: 1,
		BeginPos: beginPos,
	})
}

func (c *TeamspeakHttpClient) LogViewVirtualServer(lines int, reverse bool, beginPos int) (*[]Log, error) {
	return c.logView(logViewRequest{
		Lines:    lines,
		Reverse:  boolToInt(reverse),
		Instance: 0,
		BeginPos: beginPos,
	})
}
