package salesfly

import "github.com/salesfly/salesfly-go/internal"

// Send sends a mail message
func (c *Mail) Send(message *MailMessage) (*MailReceipt, error) {
	var resp Response

	params := message.ToParameters()

	err := c.client.post("/v1/mail/send", params, &resp)
	if err != nil {
		return nil, err
	}

	result := &MailReceipt{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
