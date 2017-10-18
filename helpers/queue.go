package helpers

import (
	"log"

	"github.com/veqryn/go-email/email"

	"github.com/kardasjan/deepa/structures"
	"gopkg.in/mgo.v2"
)

// EnqueueByMsgType Assign all recievers to a message and push them to Database
func EnqueueByMsgType(sms *structures.SMSMessage, session *mgo.Session) {
	// Gather phone numbers and send the message
	contacts := GetMsgTypeReceivers(sms.MsgType, sms.Site, session)
	EnqueueAll(contacts, sms, session)
}

// EnqueueByMsgTypeWithEmail ...
func EnqueueByMsgTypeWithEmail(sms *structures.SMSMessage, session *mgo.Session, images []*email.Message) {
	contacts := GetMsgTypeReceivers(sms.MsgType, sms.Site, session)
	EnqueueAllWithEmail(contacts, sms, images, session)
}

// EnqueueByService ...
func EnqueueByService(sms *structures.SMSMessage, service *structures.Service, session *mgo.Session) {
	contacts := GetServiceReceivers(service.ID, sms.Site, session)
	EnqueueAll(contacts, sms, session)
}

// EnqueueAll ...
func EnqueueAll(contacts []structures.Contact, sms *structures.SMSMessage, session *mgo.Session) {
	PrepareSMS(sms)
	log.Println("SMS Prepared!")

	for _, contact := range contacts {
		sms.Phone = contact.Phone
		SendSMSRest(sms.Phone, sms.Body)
	}
}

// EnqueueAllWithEmail ...
func EnqueueAllWithEmail(contacts []structures.Contact, sms *structures.SMSMessage, images []*email.Message, session *mgo.Session) {
	PrepareSMSWithEmail(sms)
	log.Println("SMS with Email Prepared!")

	for _, contact := range contacts {
		sms.Phone = contact.Phone
		SendSMSRest(sms.Phone, sms.Body)
		SendEmail(contact.Email, sms.Body, images)
	}
}
