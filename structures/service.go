package structures

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Service Database structure
type Service struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Name        string          `bson:"name"`
	Status      int             `bson:"status"`
	Description string          `bson:"description"`
	Port        int             `bson:"port"`
	CreatedAt   time.Time       `bson:"createdAt"`
	LastActive  time.Time       `bson:"LastActive"`
	SendTo      []bson.ObjectId `bson:"sendTo"`
	RetryCount  int             `bson:"retryCount"`
	MaxRetries  int             `bson:"MaxRetries"`
}
