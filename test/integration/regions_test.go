package integration

import (
	"errors"
	"os"
	"testing"

	"github.com/skandyla/go-vscale-client"
)

// Integration test, we verify communication with external API, but don't we got correct results
func TestRegionsService_List(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := vscale.NewClient(token)
	_, err = client.Region.List()
	if err != nil {
		t.Error(err)
		return
	}
}

func GetToken() (string, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return token, errors.New("Token is empty")
	}
	return token, nil
}
