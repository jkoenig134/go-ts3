package go_ts3_http

import (
	"fmt"
	"net/url"
)

/*
stringParams, err := c.encodeStruct(parameters)
	if err != nil {
		return err
	}
*/

func (c *TeamspeakHttpClient) encodeStruct(v interface{}) (string, error) {
	if v == nil {
		empty := ""
		return empty, nil
	}

	form := url.Values{}
	err := c.encoder.Encode(v, form)

	if err != nil {
		return "", err
	}

	encoded := form.Encode()

	if encoded != "" {
		encoded = fmt.Sprintf("?%s", encoded)
	}

	return encoded, nil
}
