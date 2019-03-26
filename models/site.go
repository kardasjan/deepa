package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// SiteContact Database structure
type SiteContact struct {
	ContactID    bson.ObjectId   `bson:"contactId,omitempty"`
	MessageTypes []bson.ObjectId `bson:"messageTypes,omitempty"`
}

// Site Database structure
type Site struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Status   int           `bson:"status"`
	IP       string        `bson:"ip"`
	Slug     string        `bson:"slug"`
	Contacts []SiteContact `bson:"contacts"`
}

// GetSiteBySlug Get Site object matching slug
func GetSiteBySlug(slug string) (*Site, error) {
	site := Site{}
	c := GetDB().C("sites")

	var err = c.Find(bson.M{"slug": slug}).One(&site)
	if err != nil {
		// There is no site with such slug!
		return nil, errors.New("Site with slug " + slug + " not found")
	}

	return &site, nil
}
