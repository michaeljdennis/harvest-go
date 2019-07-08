package endpoint

import (
	"encoding/json"
	"fmt"
)

// Endpointer ...
type Endpointer interface {
	GetURL() string
}

// Endpoint ...
type Endpoint struct {
	URL string
}

// NewEndpoint returns new Endpoint
func NewEndpoint(path string) *Endpoint {
	return &Endpoint{
		URL: "https://api.harvestapp.com/api/v2/" + path,
	}
}

// GetURL ...
func (e Endpoint) GetURL() string {
	return e.URL
}

// BaseModel ...
type BaseModel struct {
	PerPage      int               `json:"per_page"`
	TotalPages   int               `json:"total_pages"`
	TotalEntries int               `json:"total_entries"`
	NextPage     int               `json:"next_page"`
	PreviousPage int               `json:"previous_page"`
	Page         int               `json:"page"`
	Links        map[string]string `json:"links"`
}

// PrintModel ...
func PrintModel(m interface{}) {
	s, _ := json.MarshalIndent(m, "", "    ")
	fmt.Println(string(s))
}
