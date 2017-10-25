package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/errwrap"
	"github.com/joyent/triton-go/client"
)

type CloudAPI struct {
	Versions []string `json:"versions"`
}

type Ping struct {
	Ping     string   `json:"ping"`
	CloudAPI CloudAPI `json:"cloudapi"`
}

func (c *ComputeClient) Ping(ctx context.Context) (*Ping, error) {
	reqInputs := client.RequestInput{
		Method: http.MethodGet,
		Path:   "/--ping",
	}
	response, err := c.Client.ExecuteRequestRaw(ctx, reqInputs)
	if response != nil {
		defer response.Body.Close()
	}

	if response == nil || response.StatusCode == http.StatusNotFound || response.StatusCode == http.StatusGone {
		return nil, &client.TritonError{
			StatusCode: response.StatusCode,
			Code:       "ResourceNotFound",
		}
	}
	if err != nil {
		return nil, errwrap.Wrapf("Error executing Get request: {{err}}",
			c.Client.DecodeError(response.StatusCode, response.Body))
	}

	var result *Ping
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&result); err != nil {
		return nil, errwrap.Wrapf("Error decoding Get response: {{err}}", err)
	}

	return result, nil
}
