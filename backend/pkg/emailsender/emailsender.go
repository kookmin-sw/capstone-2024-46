package emailsender

import "context"

type EmailSender interface {
	SendEmail(ctx context.Context, to string, from string, fromName string, subject string, bodyHtml string) error
}
