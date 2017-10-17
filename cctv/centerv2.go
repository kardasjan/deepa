package cctv

import (
	"log"

	"github.com/kardasjan/deepa/structures"
	"github.com/veqryn/go-email/email"
)

// CenterV2 geovision pipe
func CenterV2(m *email.Message, sms *structures.SMSMessage) {
	log.Println(m)
	log.Println("Body: " + string(m.Body))
}
