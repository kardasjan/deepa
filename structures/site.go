package structures

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SiteContact Database structure
type SiteContact struct {
	ContactID    bson.ObjectId   `bson:"contactId,omitempty"`
	MessageTypes []bson.ObjectId `bson:"messageTypes,omitempty"`
	Services     []bson.ObjectId `bson:"services,omitempty"`
}

// Site Database structure
type Site struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Status   int           `bson:"status"`
	IP       string        `bson:"ip"`
	Slug     string        `bson:"slug"`
	Contacts []SiteContact `bson:"contacts"`
	Services []Service     `bson:"services"`
	Router   Router        `bson:"router"`
}

// Router Database structure
type Router struct {
	Status      int       `bson:"status"`
	RetryCount  int       `bson:"retryCount"`
	MaxRetries  int       `bson:"failCount"`
	Description string    `bson:"description"`
	Port        int       `bson:"port"`
	CreatedAt   time.Time `bson:"createdAt"`
	LastActive  time.Time `bson:"lastActive"`
}
