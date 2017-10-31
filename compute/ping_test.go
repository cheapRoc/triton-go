package compute_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/joyent/triton-go/authentication"
	"github.com/joyent/triton-go/client"
	"github.com/joyent/triton-go/compute"
	"github.com/joyent/triton-go/testutils"
)

// test empty response
// test borked 404 response (TritonError)
// test borked 410 response (TritonError)
// test bad JSON decode

func pingSuccessFunc(req *http.Request) (*http.Response, error) {
	body := strings.NewReader(`{
	"ping": "pong",
	"cloudapi": {
		"versions": ["7.0.0", "7.1.0", "7.2.0", "7.3.0", "8.0.0"]
	}
}`)
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	return &http.Response{
		StatusCode: 200,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func buildClient() *compute.ComputeClient {
	testSigner, _ := authentication.NewTestSigner()
	httpClient := &client.Client{
		Authorizers: []authentication.Signer{testSigner},
		HTTPClient: &http.Client{
			Transport: testutils.DefaultMockTransport,
			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
	return &compute.ComputeClient{
		Client: httpClient,
	}
}

func TestPing(t *testing.T) {
	computeClient := buildClient()

	do := func(ctx context.Context, pc *compute.ComputeClient) (*compute.PingOutput, error) {
		defer testutils.DeactivateClient()

		ping, err := pc.Ping(ctx)
		if err != nil {
			return nil, err
		}
		return ping, nil
	}

	t.Run("successful", func(t *testing.T) {
		testutils.RegisterResponder("GET", "/--ping", pingSuccessFunc)

		resp, err := do(context.Background(), computeClient)
		if err != nil {
			t.Fatal(err)
		}

		if resp.Ping != "pong" {
			t.Errorf("ping was not pong: expected %s", resp.Ping)
		}

		versions := []string{"7.0.0", "7.1.0", "7.2.0", "7.3.0", "8.0.0"}
		if !reflect.DeepEqual(resp.CloudAPI.Versions, versions) {
			t.Errorf("ping did not contain CloudAPI versions: expected %s", versions)
		}
	})

	// t.Run("empty response", func(t *testing.T) {
	// 	testutils.RegisterResponder("GET", "/--ping", pingSuccessFunc)

	// 	resp, err := do(context.Background(), computeClient)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	if resp.Ping != "pong" {
	// 		t.Errorf("ping was not pong: expected %s", resp.Ping)
	// 	}

	// 	versions := []string{"7.0.0", "7.1.0", "7.2.0", "7.3.0", "8.0.0"}
	// 	if !reflect.DeepEqual(resp.CloudAPI.Versions, versions) {
	// 		t.Errorf("ping did not contain CloudAPI versions: expected %s", versions)
	// 	}
	// })

	// t.Run("404", func(t *testing.T) {
	// })

	// t.Run("410", func(t *testing.T) {
	// })

	// t.Run("JSON decode", func(t *testing.T) {
	// })

}
