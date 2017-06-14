package compute

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/errwrap"
)

type PackagesClient struct {
	*Compute
}

type Package struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Memory      int64  `json:"memory"`
	Disk        int64  `json:"disk"`
	Swap        int64  `json:"swap"`
	LWPs        int64  `json:"lwps"`
	VCPUs       int64  `json:"vcpus"`
	Version     string `json:"version"`
	Group       string `json:"group"`
	Description string `json:"description"`
	Default     bool   `json:"default"`
}

type ListPackagesInput struct {
	Name    string `json:"name"`
	Memory  int64  `json:"memory"`
	Disk    int64  `json:"disk"`
	Swap    int64  `json:"swap"`
	LWPs    int64  `json:"lwps"`
	VCPUs   int64  `json:"vcpus"`
	Version string `json:"version"`
	Group   string `json:"group"`
}

func (c *PackagesClient) ListPackages(ctx context.Context, input *ListPackagesInput) ([]*Package, error) {
	path := fmt.Sprintf("/%s/packages", c.client.AccountName)
	respReader, err := c.executeRequest(ctx, http.MethodGet, path, input)
	if respReader != nil {
		defer respReader.Close()
	}
	if err != nil {
		return nil, errwrap.Wrapf("Error executing ListPackages request: {{err}}", err)
	}

	var result []*Package
	decoder := json.NewDecoder(respReader)
	if err = decoder.Decode(&result); err != nil {
		return nil, errwrap.Wrapf("Error decoding ListPackages response: {{err}}", err)
	}

	return result, nil
}

type GetPackageInput struct {
	ID string
}

func (c *PackagesClient) GetPackage(ctx context.Context, input *GetPackageInput) (*Package, error) {
	path := fmt.Sprintf("/%s/packages/%s", c.client.AccountName, input.ID)
	respReader, err := c.executeRequest(ctx, http.MethodGet, path, nil)
	if respReader != nil {
		defer respReader.Close()
	}
	if err != nil {
		return nil, errwrap.Wrapf("Error executing GetPackage request: {{err}}", err)
	}

	var result *Package
	decoder := json.NewDecoder(respReader)
	if err = decoder.Decode(&result); err != nil {
		return nil, errwrap.Wrapf("Error decoding GetPackage response: {{err}}", err)
	}

	return result, nil
}