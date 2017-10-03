package database

import (
	"log"

	"github.com/kardasjan/deepa/structures"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UpdateRouter Updates field in database
func UpdateRouter(session *mgo.Session, router structures.Router, siteID bson.ObjectId) {
	c := session.DB(dbName).C("sites")
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"n": 1}, "$set": bson.M{
			"router.status":     router.Status,
			"router.retryCount": router.RetryCount,
			"router.LastActive": router.LastActive}},
		Upsert:    false,
		Remove:    false,
		ReturnNew: false,
	}
	info, err := c.Find(bson.M{"_id": siteID}).Apply(change, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(info)
}
