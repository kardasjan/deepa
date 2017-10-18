package cctv

import (
	"log"

	"github.com/kardasjan/deepa/structures"
	"github.com/veqryn/go-email/email"
)

// Imggeo Geovision multipart message with image attachments
func Imggeo(m *email.Message, sms *structures.SMSMessage, images []*email.Message) {
	for _, part := range m.MessagesAll() {
		mediaType, params, err := part.Header.ContentType()
		if err == nil {
			switch mediaType {
			case "multipart/mixed":
				// I don't need this, information about whole email
			case "text/plain":
				// Whatever the message type, send SMS message
				Geovision(part, sms)
			case "application/octet-stream":
				imgName := ""
				// Parameters are map allways consisting of single attachment name
				for _, param := range params {
					imgName = param
				}
				// Append image
				imagePart := email.NewPartAttachmentFromBytes(part.Body, imgName)
				images = append(images, imagePart)
			}
		} else {
			log.Println(err)
		}
	}
}
