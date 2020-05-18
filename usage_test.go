package salesfly_test

import (
	"testing"

	"github.com/salesfly/salesfly-go"
	"github.com/stretchr/testify/assert"
)

func Test_GetUsage(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.Usage.Get()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEqual(t, 0, result.Allowed)
	assert.NotEqual(t, 0, result.Used)
}
