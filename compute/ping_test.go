package compute_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/joyent/triton-go/client"
	"github.com/joyent/triton-go/compute"
	"github.com/joyent/triton-go/testutils"
)

// test empty response
// test borked 404 response (TritonError)
// test borked 410 response (TritonError)
// test bad JSON decode

func TestPing(t *testing.T) {
	httpClient := &client.Client{
		HTTPClient: &http.Client{
			Transport: http.DefaultTransport,
			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}

	c := &compute.ComputeClient{
		Client: httpClient,
	}

	do := func(ctx context.Context, pc *compute.ComputeClient) (*compute.PingOutput, error) {
		testutils.ActivateClient(false)

		defer testutils.DeactivateClient()
		ping, err := pc.Ping(ctx)
		if err != nil {
			return nil, err
		}
		return ping, nil
	}

	t.Run("successful", func(t *testing.T) {
		testutils.RegisterResponder("GET", "/--ping", func(req *http.Request) (*http.Response, error) {
			body := strings.NewReader(`{ 'ping': 'pong' }`)
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(body),
			}, nil
		})

		ctx := context.Background()
		resp, err := do(ctx, c)
		if err == nil {
			t.Fatal(err)
		}
		if resp.Ping != "pong" {
			t.Error("ping was not pong: expected %s", resp.Ping)
		}
	})

	// t.Run("empty response", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	compute.Client.SetResponse(nil)
	// 	resp, err := do(compute)
	// 	if err == nil {
	// 		t.Error(err)
	// 	}
	// })

	// t.Run("404", func(t *testing.T) {
	// })

	// t.Run("410", func(t *testing.T) {
	// })

	// t.Run("JSON decode", func(t *testing.T) {
	// })

}
