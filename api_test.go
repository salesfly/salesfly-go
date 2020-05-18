package salesfly_test

import (
	"log"
	"os"
	"testing"

	"github.com/salesfly/salesfly-go"
	"github.com/stretchr/testify/assert"
)

var apiKey string

func init() {
	apiKey = os.Getenv("SALESFLY_APIKEY")
	if len(apiKey) == 0 {
		log.Fatalln("SALESFLY_APIKEY environment variable not defined")
	}
}

func Test_GetVersion(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	v := client.GetVersion()
	assert.NotEmpty(t, v)
}
