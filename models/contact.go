package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Contact Database structure
type Contact struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Phone string        `bson:"phone"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
}

// GetContactsByID Get MANY contacts by IDs
func GetContactsByID(ids []bson.ObjectId) ([]Contact, error) {
	var result []Contact
	c := GetDB().C("contacts")
	var err = c.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&result)
	if err != nil {
		return nil, errors.New("No contacts found")
	}
	return result, nil
}

// GetNumbersWhoShouldReceiveMessage Walk the DB and get all the phone numbers who should receive msgType for this the site
func GetNumbersWhoShouldReceiveMessage(site *Site, msgType bson.ObjectId) []string {
	sendTo := filterMsgTypeContacts(site, msgType) // Contact IDs who should receive
	contacts, _ := GetContactsByID(sendTo)         // Contact structures who should receive

	var numbers []string

	for _, contact := range contacts {
		numbers = append(numbers, contact.Phone)
	}
	return numbers
}

// Walk through contacts and filter which are supposed to recieve the message
func filterMsgTypeContacts(site *Site, messageType bson.ObjectId) []bson.ObjectId {
	var contacts []bson.ObjectId
	// For each contact in a Site
	for _, contact := range site.Contacts {
		// Compare to the MsgType
		for _, msgType := range contact.MessageTypes {
			// And if this msgType is present
			if messageType == msgType {
				// That means this contactID should get the message
				contacts = append(contacts, contact.ContactID)
				break
			}
		}
	}

	/*
		if len(contacts) <= 0 {
			log.Println("No contact should recieve this message")
		}
	*/

	return contacts
}
