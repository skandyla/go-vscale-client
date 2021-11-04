package vscale

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// UT List function implementation
func TestRegions_List(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/v1/locations", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[
  {
    "id": "spb0",
    "description": "",
    "active": true,
    "private_networking": false,
    "rplans": [
      "small",
      "medium",
      "large",
      "huge",
      "monster"
    ],
    "templates": [
      "ubuntu_20.04_64_001_ajenti",
      "debian_10_64_001_fastpanel",
      "CentOS_8_64_001_master"
    ]
  }]`)
	})

	resource := NewClient("")

	// specify params of our testing server
	resource.client = server.Client()
	resource.BaseURL = server.URL

	regions, err := resource.Region.List()
	if err != nil {
		t.Errorf("Regions.List returned error: %v", err)
	}

	expectedRegions := []Location{
		{
			ID:                "spb0",
			Description:       "",
			Active:            true,
			PrivateNetworking: false,
			Rplans:            []string{"small", "medium", "large", "huge", "monster"},
			Templates:         []string{"ubuntu_20.04_64_001_ajenti", "debian_10_64_001_fastpanel", "CentOS_8_64_001_master"},
		},
	}

	// comparing test results:
	// 1 reflect:
	//if !reflect.DeepEqual(regions, expectedRegions) {
	//	t.Errorf("Regions.List returned regions %+v, expected %+v", regions, expectedRegions)
	//}

	// 2 stretchr/testify:
	//assert.Equal(t, expectedRegions, regions)

	// 3 google cmp:
	diff := cmp.Diff(expectedRegions, regions)
	if diff != "" {
		t.Fatalf(diff)
	}
}
