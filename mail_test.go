package salesfly_test

import (
	"testing"

	"github.com/salesfly/salesfly-go"
	"github.com/stretchr/testify/assert"
)

func Test_Mail_Send0(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// OK
	m := salesfly.NewMessage("test2@demo2.org", "Test", "This is just a test: http://www.example.com.", "test2@demo2.org")
	m.SetHTML(`<html><body>This is just a test in HTML! <a href="http://www.example.com">Some link</a></body></html>`)
	m.AddAttachment("/Users/otto/me.png")
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)
}

func Test_Mail_Send(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// OK
	m := salesfly.NewMessage("ok@demo2.org", "Test", "This is just a test: http://www.example.com.", "ok@demo2.org", "okey@demo2.org")
	m.SetHTML(`<html><body>This is just a test in HTML! <a href="http://www.example.com">Some link</a></body></html>`)
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)
}

func Test_Mail_Send2(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// OK
	m := salesfly.NewMessage("ok@demo2.org", "Test", "This is just a test", "ok@demo2.org")
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)
}

func Test_Mail_Send_Hard_Bounce(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Hard bounce - permanent error
	m := salesfly.NewMessage("bounce@demo2.org", "Test", "This is just a test", "ok@demo2.org")
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)

}

func Test_Mail_Send_Graylisted(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Sender gray listed - transient error
	m := salesfly.NewMessage("gray@demo2.org", "Test", "This is just a test", "ok@demo2.org")
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)

}

func Test_Mail_Send_Mailbox_Full(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Mailbox temporary full
	m := salesfly.NewMessage("ok@demo2.org", "Test", "This is just a test", "full@demo2.org")
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)

}

func Test_Mail_Send_Unknown_Recipient(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Mailbox not found - permanment error
	m := salesfly.NewMessage("ok@demo2.org", "Test", "This is just a test", "none@demo2.org")
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)

}

func Test_Mail_Send_DomainError(t *testing.T) {
	client, err := salesfly.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	m := salesfly.NewMessage("ok@demo2.org", "Test", "This is just a test", "none@demo2.xxx")
	m.SetTestMode(true)
	receipt, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, receipt)
}
