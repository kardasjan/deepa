package cctv

import (
	"log"

	"github.com/kardasjan/deepa/structures"
	"github.com/veqryn/go-email/email"
)

// Imggeo geovision pipe
func Imggeo(m *email.Message, sms *structures.SMSMessage) {
	for _, part := range m.MessagesAll() {
		mediaType, _, err := part.Header.ContentType()
		if err == nil {
			switch mediaType {
			case "multipart/mixed":
				log.Println(part)
			case "text/plain":
				log.Panicln("Text: " + string(part.Body))
			case "application/octet-stream":
				log.Println(part.Header)
			}
		} else {
			log.Println(err)
		}
	}
}
