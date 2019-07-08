package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/michaeljdennis/harvest-go/config"
	"github.com/michaeljdennis/harvest-go/endpoint"
)

// APIClient ...
type APIClient struct {
	http.Client
	Config config.Config
}

// New returns new APIClient
func New(config config.Config) *APIClient {
	return &APIClient{
		Config: config,
	}
}

func (ac APIClient) createRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Add("Harvest-Account-ID", ac.Config.HarvestAccountID)
	req.Header.Add("Authorization", "Bearer "+ac.Config.HarvestToken)
	req.Header.Add("User-Agent", "Harvest Go (https://github.com/michaeljdennis/harvest-go)")

	return req, err
}

// Get takes endpoint e, makes an API request, and hydrates the properties of model m
func (ac APIClient) Get(e endpoint.Endpointer, m interface{}) {
	req, err := ac.createRequest("GET", e.GetURL(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := ac.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyStr := string(body)
	// fmt.Println("bodyStr:", bodyStr)

	err = json.NewDecoder(strings.NewReader(bodyStr)).Decode(&m)
	if err != nil {
		log.Fatalln(err)
	}
}
