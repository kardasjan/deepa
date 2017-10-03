package helpers

import (
	"log"

	"github.com/kardasjan/deepa/database"
	"github.com/kardasjan/deepa/structures"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetMsgTypeReceivers Gathers all Contacts which are supposed to receive the messageType
func GetMsgTypeReceivers(messageType bson.ObjectId, slug string, session *mgo.Session) []structures.Contact {
	data := database.GetSiteBySlug(slug, session)
	sendTo := filterMsgTypeContacts(data, messageType)
	contacts := database.GetContactsByID(sendTo, session)

	var result []structures.Contact

	for _, contact := range contacts {
		result = append(result, contact)
	}
	return result
}

// GetServiceReceivers Get contacts who should be informed of service state change
func GetServiceReceivers(service bson.ObjectId, slug string, session *mgo.Session) []structures.Contact {
	data := database.GetSiteBySlug(slug, session)
	sendTo := filterServiceContacts(data, service)
	contacts := database.GetContactsByID(sendTo, session)

	var result []structures.Contact

	for _, contact := range contacts {
		result = append(result, contact)
	}
	return result
}

// Walk through contacts and filter which are supposed to recieve the message
func filterMsgTypeContacts(site *structures.Site, messageType bson.ObjectId) []bson.ObjectId {
	var contacts []bson.ObjectId
	for _, contact := range site.Contacts {
		log.Println("Comparing contact: " + contact.ContactID.Hex())
		for _, msgType := range contact.MessageTypes {
			log.Println("Comparing msgType: " + msgType.Hex())
			if messageType == msgType {
				log.Println("MsgType MATCH!")
				contacts = append(contacts, contact.ContactID)
				break
			}
		}
	}

	if len(contacts) <= 0 {
		log.Println("No contact should recieve this message")
		log.Fatal(contacts)
	}

	return contacts
}

func filterServiceContacts(site *structures.Site, serviceID bson.ObjectId) []bson.ObjectId {
	var contacts []bson.ObjectId
	for _, contact := range site.Contacts {
		log.Println("Comparing contact: " + contact.ContactID.Hex())
		for _, service := range contact.Services {
			log.Println("Comparing service: " + service.Hex())
			if serviceID == service {
				log.Println("Service MATCH!")
				contacts = append(contacts, contact.ContactID)
				break
			}
		}
	}

	if len(contacts) <= 0 {
		log.Println("No contact should recieve this message")
		log.Fatal(contacts)
	}

	return contacts
}
