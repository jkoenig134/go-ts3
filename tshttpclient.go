package go_ts3

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/jkoenig134/schema"
	"github.com/valyala/fasthttp"
)

type Config struct {
	baseUrl string
	apiKey  string
}

func NewConfig(baseUrl string, apiKey string) Config {
	return Config{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

type TeamspeakHttpClient struct {
	config      Config
	httpClient  fasthttp.Client
	eventBus    EventBus.Bus
	encoder     schema.Encoder
	serverID    int
	eventClient *EventClient
}

func (c *TeamspeakHttpClient) SetServerID(serverID int) {
	c.serverID = serverID

	if c.eventClient != nil {
		_ = c.eventClient.SwitchServer(serverID)
	}
}

func NewClient(config Config) TeamspeakHttpClient {
	return TeamspeakHttpClient{
		config,
		fasthttp.Client{},
		EventBus.New(),
		*schema.NewEncoder(),
		1,
		nil,
	}
}

type status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type tsResponse struct {
	Body   interface{} `json:"body"`
	Status status      `json:"status"`
}

func (t *tsResponse) asError() error {
	return fmt.Errorf(
		"Query returned non 0 exit code: '%d'. Message: '%s' ",
		t.Status.Code,
		t.Status.Message)
}

func (c *TeamspeakHttpClient) requestWithParams(path string, paramStruct interface{}, v interface{}) error {
	param, err := c.encodeStruct(paramStruct)
	if err != nil {
		return err
	}

	merged := fmt.Sprintf("%s%s", path, param)
	return c.request(merged, v)
}

func (c *TeamspeakHttpClient) request(path string, v interface{}) error {
	requestUrl := fmt.Sprintf("%s/%d/%s", c.config.baseUrl, c.serverID, path)
	request := fasthttp.AcquireRequest()
	request.Header.Set("x-api-key", c.config.apiKey)
	request.SetRequestURI(requestUrl)
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	err := c.httpClient.Do(request, response)
	if err != nil {
		return err
	}

	stringBody := string(response.Body())
	code := response.StatusCode()

	if code != 200 {
		err = fmt.Errorf("Error Code: %d\n%s", code, stringBody)
		return err
	}

	tsResponse := &tsResponse{}
	err = json.Unmarshal(response.Body(), tsResponse)
	if err != nil {
		return err
	}

	if tsResponse.Status.Code != 0 {
		return tsResponse.asError()
	}

	if v == nil {
		return nil
	}

	jsonBody, err := json.Marshal(tsResponse.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonBody, v); err != nil {
		return err
	}

	return nil
}
