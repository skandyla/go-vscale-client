package vscale

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// From outdated SDK
//type Scalet struct {
//	Hostname        string        `json:"hostname,omitempty"`
//	Locked          bool          `json:"locked,omitempty"`
//	Location        string        `json:"location,omitempty"`
//	Rplan           string        `json:"rplan,omitempty"`
//	Name            string        `json:"name,omitempty"`
//	Active          bool          `json:"active,omitempty"`
//	Keys            []Keys        `json:"keys,omitempty"`
//	PublicAddresses PublicAddress `json:"public_address,omitempty"`
//	Status          string        `json:"status,omitempty"`
//	MadeFrom        string        `json:"made_from,omitempty"`
//	CTID            int64         `json:"ctid,omitempty"`
//}

//Autogenerated
type Scalet struct {
	Ctid              int           `json:"ctid,omitempty"`
	Name              string        `json:"name,omitempty"`
	Status            string        `json:"status,omitempty"`
	Location          string        `json:"location,omitempty"`
	Rplan             string        `json:"rplan,omitempty"`
	Keys              []Keys        `json:"keys,omitempty"`
	Tags              []interface{} `json:"tags,omitempty"`
	PublicAddress     PublicAddress `json:"public_address,omitempty"`
	PrivateAddress    interface{}   `json:"private_address,omitempty"`
	MadeFrom          string        `json:"made_from,omitempty"`
	Hostname          string        `json:"hostname,omitempty"`
	Created           string        `json:"created,omitempty"`
	Active            bool          `json:"active,omitempty"`
	Locked            bool          `json:"locked,omitempty"`
	Deleted           interface{}   `json:"deleted,omitempty"`
	BlockReason       interface{}   `json:"block_reason,omitempty"`
	BlockReasonCustom interface{}   `json:"block_reason_custom,omitempty"`
	DateBlock         interface{}   `json:"date_block,omitempty"`
}

type Keys struct {
	Name string `json:"name,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

type PublicAddress struct {
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Address string `json:"address,omitempty"`
}

// ScaletsService is an interface for interfacing with the scalets
// endpoints of the Vscale API
type ScaletsService interface {
	List() ([]Scalet, error)
	Create(createRequest *ScaletCreateRequest) (Scalet, error)
}

// ScaletsServiceOp handles communication with the Scalets related methods of the vscale API
type ScaletsServiceOp struct {
	client *Client
}

// List Scalets
func (s *ScaletsServiceOp) List() ([]Scalet, error) {
	req, err := http.NewRequest(http.MethodGet, s.client.BaseURL+"/v1/scalets", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Token", s.client.Token)

	resp, err := s.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []Scalet
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}

type ScaletCreateRequest struct {
	MakeFrom string  `json:"make_from,omitempty"`
	Rplan    string  `json:"rplan,omitempty"`
	DoStart  bool    `json:"do_start,omitempty"`
	Name     string  `json:"name,omitempty"`
	Keys     []int64 `json:"keys,omitempty"`
	Password string  `json:"password,omitempty"`
	Location string  `json:"location,omitempty"`
}

// Create Scalet
func (s *ScaletsServiceOp) Create(createRequest *ScaletCreateRequest) (Scalet, error) {
	bytesRepresentation, err := json.Marshal(createRequest)
	if err != nil {
		log.Fatalln(err)
	}

	scalet := Scalet{}

	req, err := http.NewRequest(http.MethodPost, s.client.BaseURL+"/v1/scalets", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return scalet, err
	}
	req.Header.Add("X-Token", s.client.Token)

	resp, err := s.client.client.Do(req)
	if err != nil {
		return scalet, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return scalet, err
	}

	if err = json.Unmarshal(body, &scalet); err != nil {
		log.Fatal(err)
	}
	return scalet, nil
}
