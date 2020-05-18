package salesfly_test

import (
	"os"
	"testing"

	"github.com/salesfly/salesfly-go"
	"github.com/stretchr/testify/assert"
)

func Test_SMS_Send(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	to := os.Getenv("MY_PHONE_NUMBER")
	if len(to) == 0 {
		to = "+12125551234"
	}

	receipt, err := client.SMS.Send("Objectia", to, "This is a test")
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)
}
