package structures

import "gopkg.in/mgo.v2/bson"

// Contact Database structure
type Contact struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Phone string        `bson:"phone"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
}
