package structures

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SMSMessage Wrapper for SMS
type SMSMessage struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Body     string
	Datetime time.Time
	Site     string
	MsgType  bson.ObjectId
	SiteName string
	Retries  int
	Phone    string
}
