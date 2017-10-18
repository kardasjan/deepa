package cctv

import (
	"log"

	"github.com/kardasjan/deepa/structures"
	"github.com/veqryn/go-email/email"
)

// Imggeo geovision pipe
func Imggeo(m *email.Message, sms *structures.SMSMessage) {
	for _, part := range m.MessagesAll() {
		mediaType, params, err := part.Header.ContentType()
		if err != nil {
			switch mediaType {
			case "text/plain":
				log.Println(part)
			default:
				log.Println(part)
			}
		} else {
			log.Println(err)
		}
		log.Println(params)
	}
}
