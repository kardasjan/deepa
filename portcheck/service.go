package portcheck

import (
	"log"
	"time"

	"github.com/kardasjan/deepa/helpers"
	"github.com/kardasjan/deepa/structures"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func updateService(service *structures.Service, session *mgo.Session) {
	c := session.DB("watch").C("services")
	err := c.Update(bson.M{"_id": service.ID}, service)
	log.Println(err)
}

func sendService(message string, service *structures.Service, site *structures.Site, session *mgo.Session) {
	sms := &structures.SMSMessage{Retries: 0}
	sms.SiteName = site.Name
	sms.Datetime = time.Now()
	helpers.EnqueueByService(sms, service, session)
}
