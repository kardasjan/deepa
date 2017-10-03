package database

import (
	"log"

	"github.com/kardasjan/deepa/structures"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SMSRetryLimit ...
const SMSRetryLimit = 3
const maxBufferSize = 10 // Send only 10 messages at once

// QueueMessage Push SMS to Queue
func QueueMessage(sms *structures.SMSMessage, session *mgo.Session) (err error) {
	log.Println("--- insertMessage ")
	err = session.DB(dbName).C("queue").Insert(sms)
	if err != nil {
		log.Println("insertMessage: ", err)
		return
	}
	return
}

// GetPendingMessages Fetch database
func GetPendingMessages(session *mgo.Session) (messages []structures.SMSMessage, err error) {
	log.Println("--- getPendingMessages ")
	err = session.DB(dbName).C("queue").Find(bson.M{"retries": bson.M{"$lte": SMSRetryLimit}}).Sort("-retries").Limit(maxBufferSize).All(&messages)
	return
}

// MessageSent Trigger if message was sent and OK
func MessageSent(sms *structures.SMSMessage, session *mgo.Session) (err error) {
	log.Println("--- messageSent ")
	err = session.DB(dbName).C("queue").Remove(bson.M{"_id": sms.ID})
	if err != nil {
		log.Println("removeMessage: ", err)
		return
	}
	err = session.DB(dbName).C("sent").Insert(sms)
	return
}

// MessageNotSent Usually when message was not send, increment retires
func MessageNotSent(sms *structures.SMSMessage, session *mgo.Session) (err error) {
	log.Println("--- updateMessageStatus ")
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"retries": sms.Retries,
		}},
	}
	_, err = session.DB(dbName).C("queue").Find(bson.M{"_id": sms.ID}).Apply(change, nil)
	return
}
