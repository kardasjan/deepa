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
			log.Println(mediaType)
			log.Println(part.Header)
		} else {
			log.Println(err)
		}
		for _, param := range params {
			log.Println("Param: " + param)
		}
	}
}
