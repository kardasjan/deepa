package cctv

import (
	"log"

	"github.com/kardasjan/deepa/structures"
	"github.com/veqryn/go-email/email"
)

// Imggeo geovision pipe
func Imggeo(m *email.Message, sms *structures.SMSMessage) {
	log.Println(m)
	log.Println("Body: " + string(m.Body))
}
