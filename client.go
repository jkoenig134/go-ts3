package go_ts3_http

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type Config struct {
	baseUrl string
	apiKey  string
	debug   bool
}

func NewConfig(baseUrl string, apiKey string, debug bool) Config {
	return Config{
		baseUrl: baseUrl,
		apiKey:  apiKey,
		debug:   debug,
	}
}

type Client struct {
	config     Config
	httpClient fasthttp.Client
}

func NewClient(config Config) Client {
	return Client{
		config,
		fasthttp.Client{},
	}
}

func (c *Client) request(path string, v interface{}) error {
	url := fmt.Sprintf("%s/%s", c.config.baseUrl, path)

	request := fasthttp.AcquireRequest()
	request.Header.Set("x-api-key", c.config.apiKey)
	request.SetRequestURI(url)
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

	if c.config.debug {
		fmt.Println(stringBody)
	}

	err = json.Unmarshal(response.Body(), v)
	if err != nil {
		return err
	}

	return nil
}

func vServerUrl(id int, path string) string {
	return fmt.Sprintf("%d/%s", id, path)
}
