package portcheck

import (
	"log"
	"time"

	"github.com/kardasjan/deepa/helpers"
	"github.com/kardasjan/deepa/structures"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func updateRouter(router *structures.Router, siteID bson.ObjectId, session *mgo.Session) {
	doc := structures.Router{} // After updating here are new values
	c := session.DB("watch").C("sites")
	change := mgo.Change{
		Update:    bson.M{"router": router},
		ReturnNew: true,
	}
	info, err := c.Find(bson.M{"_id": siteID}).Apply(change, &doc)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(info)
	}
}

func sendRouter(message string, site *structures.Site, session *mgo.Session) {
	sms := &structures.SMSMessage{Retries: 0}
	sms.SiteName = site.Name
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(routerMsgType)
	helpers.EnqueueByMsgType(sms, session)
}
