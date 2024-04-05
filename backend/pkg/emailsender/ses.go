package emailsender

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"

	"private-llm-backend/pkg/errorutil"
)

var _ EmailSender = (*sesEmailSender)(nil)

type sesEmailSender struct {
	client *ses.Client
}

func (s *sesEmailSender) SendEmail(ctx context.Context, to string, from string, fromName string, subject string, bodyHtml string) error {
	if fromName != "" {
		from = fmt.Sprintf("%s <%s>", fromName, from)
	}
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Data: &bodyHtml,
				},
			},
			Subject: &types.Content{
				Data: &subject,
			},
		},
		Source: &from,
	}
	_, err := s.client.SendEmail(ctx, input)
	if err != nil {
		return errorutil.WithDetail(err, errors.New("failed to send email"))
	}
	return nil
}

func NewSESEmailSender(ses *ses.Client) EmailSender {
	return &sesEmailSender{
		client: ses,
	}
}
