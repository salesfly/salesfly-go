package salesfly_test

import (
	"testing"

	"github.com/salesfly/salesfly-go"
	"github.com/stretchr/testify/assert"
)

func Test_Usage_Get(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.Usage.Get()
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
