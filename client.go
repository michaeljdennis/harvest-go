package harvest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/michaeljdennis/harvest-go/endpoint"
)

// APIClient embeds an http.Client and has a pointer to Config.
type APIClient struct {
	http.Client
	Config *Config
}

// NewClient returns a pointer to a new APIClient.
func NewClient(config *Config) *APIClient {
	return &APIClient{
		Config: config,
	}
}

// createRequest returns an pointer to an http.Request with
// the appropriate authentication headers set.
func (ac APIClient) createRequest(method string, url string, body io.Reader) (*http.Request, error) {
	// Create request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set authentication headers
	req.Header.Add("Harvest-Account-ID", ac.Config.HarvestAccountID)
	req.Header.Add("Authorization", "Bearer "+ac.Config.HarvestToken)
	req.Header.Add("User-Agent", "Harvest Go (https://github.com/michaeljdennis/harvest-go)")

	return req, nil
}

// Get takes endpoint e, makes an API request, and hydrates the properties of model m
func (ac APIClient) Get(e endpoint.Endpointer, m interface{}) error {
	req, err := ac.createRequest("GET", e.GetURL(), nil)
	if err != nil {
		return err
	}

	resp, err := ac.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyStr := string(body)

	err = json.NewDecoder(strings.NewReader(bodyStr)).Decode(&m)
	if err != nil {
		return err
	}

	return nil
}
