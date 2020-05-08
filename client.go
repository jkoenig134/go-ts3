package go_ts3_http

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/jkoenig134/schema"
	"github.com/valyala/fasthttp"
	"net/url"
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
	config     Config
	httpClient fasthttp.Client
	eventBus   EventBus.Bus
}

func NewClient(config Config) TeamspeakHttpClient {
	return TeamspeakHttpClient{
		config,
		fasthttp.Client{},
		EventBus.New(),
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

func (c *TeamspeakHttpClient) request(path string, parameters interface{}, v interface{}) error {
	stringParams := ""

	if parameters != nil {
		form := url.Values{}
		err := schema.NewEncoder().Encode(parameters, form)

		if err != nil {
			return err
		}

		encoded := form.Encode()

		if encoded != "" {
			fmt.Println(form.Encode())
			stringParams = fmt.Sprintf("?%s", encoded)
		}
	}

	requestUrl := fmt.Sprintf("%s/%s%s", c.config.baseUrl, path, stringParams)
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
		return fmt.Errorf(
			"Query returned non 0 exit code: '%d'. Message: '%s' ",
			tsResponse.Status.Code,
			tsResponse.Status.Message)
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

func vServerUrl(id int, path string) string {
	return fmt.Sprintf("%d/%s", id, path)
}
