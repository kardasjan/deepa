package database

import (
	"log"

	"github.com/kardasjan/deepa/structures"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetSiteBySlug Get Site object matching slug
func GetSiteBySlug(slug string, session *mgo.Session) *structures.Site {
	site := new(structures.Site)
	c := session.DB(dbName).C("sites")

	var err = c.Find(bson.M{"slug": slug}).One(&site)
	if err != nil {
		log.Println("There is no site with such slug!")
		log.Fatal(err)
	}

	return site
}

// GetAllSites ...
func GetAllSites(session *mgo.Session) (result []structures.Site) {
	c := session.DB(dbName).C("sites")
	var err = c.Find(nil).All(&result)
	if err != nil {
		log.Println("There is no site with such slug!")
		log.Fatal(err)
	}
	return
}
