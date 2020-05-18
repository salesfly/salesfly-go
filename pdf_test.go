package salesfly_test

import (
	"io/ioutil"
	"testing"

	"github.com/salesfly/salesfly-go"
	"github.com/stretchr/testify/assert"
)

func Test_PDF_HTML(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	options := salesfly.PDFOptions{
		DocumentURL: "https://example.com",
	}
	buffer, err := client.PDF.Create(&options)
	assert.NoError(t, err)
	assert.NotEmpty(t, buffer)

	ioutil.WriteFile("/tmp/test-go.pdf", buffer, 0644)
}
