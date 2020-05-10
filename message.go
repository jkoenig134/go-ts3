package go_ts3_http

type Message struct {
	MessageId int    `json:"msgid,string"`
	ClientId  string `json:"cluid"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

// gm `manage_scope`
type gmRequest struct {
	Text string `schema:"msg,required"`
}

func (c *TeamspeakHttpClient) GlobalMessage(message string) error {
	return c.requestWithParams("gm", gmRequest{Text: message}, nil)
}

// messageadd `manage_scope, write_scope`
type messageAddRequest struct {
	ClientUid string `schema:"cluid"`
	Subject   string `schema:"subject"`
	Message   string `schema:"message"`
}

func (c *TeamspeakHttpClient) MessageAdd(clientUid string, subject string, message string) error {
	return c.requestWithParams("messageadd", messageAddRequest{
		ClientUid: clientUid,
		Subject:   subject,
		Message:   message,
	}, nil)
}

// messagedel `manage_scope, write_scope`
type messageDeleteRequest struct {
	MessageId int `schema:"msgid,required"`
}

func (c *TeamspeakHttpClient) MessageDel(messageId int) error {
	return c.requestWithParams("messageget", messageDeleteRequest{MessageId: messageId}, nil)
}

// messageget `manage_scope, write_scope, read_scope`
type messageGetRequest struct {
	MessageId int `schema:"msgid,required"`
}

func (c *TeamspeakHttpClient) MessageGet(messageId int) (*Message, error) {
	var messages []Message
	err := c.requestWithParams("messageget", messageGetRequest{MessageId: messageId}, &messages)
	if err != nil {
		return nil, err
	}

	return &messages[0], nil
}

// messagelist `manage_scope, write_scope, read_scope`
func (c *TeamspeakHttpClient) MessageList() (*[]Message, error) {
	var messages []Message

	err := c.request("messagelist", &messages)
	if err != nil {
		return nil, err
	}

	return &messages, nil
}

// messageupdateflag `manage_scope, write_scope`
type messageUpdateFlagRequest struct {
	MessageId string `schema:"msgid,required"`
	Flag      int    `schema:"flag,required"`
}

func (c *TeamspeakHttpClient) MessageSetRead(messageId string) error {
	return c.requestWithParams("messageupdateflag", messageUpdateFlagRequest{
		MessageId: messageId,
		Flag:      1,
	}, nil)
}

func (c *TeamspeakHttpClient) MessageSetUnread(messageId string) error {
	return c.requestWithParams("messageupdateflag", messageUpdateFlagRequest{
		MessageId: messageId,
		Flag:      1,
	}, nil)
}

// sendtextmessage `manage_scope, write_scope`
type sendMessageRequest struct {
	TargetMode int    `schema:"targetmode,required"`
	TargetId   int    `schema:"target,required"`
	Message    string `schema:"msg,required"`
}

func (c *TeamspeakHttpClient) SendChannelMessage(message string) error {
	return c.sendMessage(sendMessageRequest{
		TargetMode: 2,
		TargetId:   1,
		Message:    message,
	})
}

func (c *TeamspeakHttpClient) SendClientMessage(targetId int, message string) error {
	return c.sendMessage(sendMessageRequest{
		TargetMode: 1,
		TargetId:   targetId,
		Message:    message,
	})
}

func (c *TeamspeakHttpClient) sendMessage(request sendMessageRequest) error {
	return c.requestWithParams("sendtextmessage", request, nil)
}
