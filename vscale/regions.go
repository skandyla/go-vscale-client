package vscale

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Location struct {
	ID                string   `json:"id"`
	Description       string   `json:"description"`
	Active            bool     `json:"active"`
	PrivateNetworking bool     `json:"private_networking"`
	Rplans            []string `json:"rplans"`
	Templates         []string `json:"templates"`
}

// RegionsService is an interface for interfacing with the Locations
// endpoints of the Vscale API
type RegionsService interface {
	List() ([]Location, error)
}

// RegionsServiceOp handles communication with the Locations related methods of the vscale API
type RegionsServiceOp struct {
	client *Client
}

// List Locations
func (s *RegionsServiceOp) List() ([]Location, error) {
	req, err := http.NewRequest(http.MethodGet, s.client.BaseURL+"/v1/locations", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Token", s.client.Token)

	resp, err := s.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []Location
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}
