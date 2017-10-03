package structures

import "gopkg.in/mgo.v2/bson"

// MessageType Database structure
type MessageType struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	Technology string        `bson:"technology"`
}
