package salesfly_test

import (
	"testing"

	"github.com/salesfly/salesfly-go"
	"github.com/stretchr/testify/assert"
)

func Test_GeoLocation_Get(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.GeoLocation.Get("8.8.8.8", nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "US", result.CountryCode)
}

func Test_GeoLocation_Get_with_options(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	options := &salesfly.GeoLocationOptions{
		DisplayFields: "country_code", // Return only country code
	}

	result, err := client.GeoLocation.Get("8.8.8.8", options)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "US", result.CountryCode)
}

func Test_GeoLocation_GetBulk(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	results, err := client.GeoLocation.GetBulk([]string{"8.8.8.8", "apple.com"}, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	for _, v := range results {
		assert.NotNil(t, v)
		assert.Equal(t, "US", v.CountryCode)
	}
}

func Test_GeoLocation_GetCurrent(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.GeoLocation.GetCurrent(nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func Test_GeoLocation_Get_Invalid_IP(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	_, err = client.GeoLocation.Get("x", nil)
	assert.Error(t, err)
	if err != nil {
		e := err.(*salesfly.Error)
		assert.Equal(t, "err-invalid-ip", e.Code)
	}
}
