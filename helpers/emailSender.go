package helpers

import (
	"net/smtp"

	"github.com/veqryn/go-email/email"
)

const (
	from    = "dohled@apsystem.cz"
	subject = "Zpráva narušitel"

	serverPort = ""

	identity = ""
	username = ""
	password = ""
	host     = ""
)

// SendEmail ...
func SendEmail(to string, body string, images []*email.Message) {
	header := email.Header{}
	header.SetFrom(from)
	header.SetTo(to)
	header.SetSubject(subject)

	msg := email.NewMessage(header, body, "", images...)

	msg.Send(serverPort, smtp.PlainAuth(identity, username, password, host))
}
