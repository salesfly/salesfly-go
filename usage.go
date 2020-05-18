package salesfly

import "github.com/salesfly/salesfly-go/internal"

// Get returns the API usage for current month.
func (c *Usage) Get() (*APIUsage, error) {
	var resp Response
	_, err := c.client.get("/v1/usage", nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &APIUsage{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
