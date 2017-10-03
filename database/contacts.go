package database

import (
	"log"

	"github.com/kardasjan/deepa/structures"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dbName = "watch"
var contactsColection = "contacts"

// GetContactsByID Get MANY contacts by IDs
func GetContactsByID(ids []bson.ObjectId, session *mgo.Session) []structures.Contact {
	var result []structures.Contact
	c := session.DB(dbName).C(contactsColection)
	var err = c.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&result)
	if err != nil {
		log.Println("func getContactsbyID probably did not return any Contacts")
		log.Fatal(err)
	}

	return result
}

// GetContactByID Get ONE contact by ID
func GetContactByID(id []bson.ObjectId, session *mgo.Session) []structures.Contact {
	// Do I need this function?
	var result []structures.Contact
	return result
}
