package main

import (
	"flag"
	"log"
	"os"

	"github.com/veqryn/go-email/email"
	"gopkg.in/mgo.v2"

	"github.com/kardasjan/deepa/cctv"
	"github.com/kardasjan/deepa/database"
	"github.com/kardasjan/deepa/ezs"
	"github.com/kardasjan/deepa/helpers"
	"github.com/kardasjan/deepa/portcheck"
	"github.com/kardasjan/deepa/structures"
)

const (
	dbName             = "watch"
	contactsCollection = "contacts"
	sitesCollection    = "sites"
	msgTypesCollection = "messageTypes"
	logFile            = "/var/log/watch.log"
)

func main() {
	// Connect to Database
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Set logging file
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Gather Flags
	prdx := flag.Bool("paradox", false, "Incoming message is from Paradox IP module")
	gv := flag.Bool("geovision", false, "Incoming message is from Geovision CCTV software (Supports Analog & Digital)")

	testdb := flag.Bool("testdb", false, "Fills database with test data")
	prdxTest := flag.Bool("paradoxTest", false, "Test paradox functionality")
	gvTest := flag.Bool("gvTest", false, "Test geovision functionality")

	port := flag.Bool("port", false, "Start port checking")
	portTest := flag.Bool("portTest", false, "Test port functionality")

	restTest := flag.Bool("restTest", false, "Test REST API")

	flag.Parse()

	// Generate SMS Message object
	smsMsg := &structures.SMSMessage{Retries: 0}

	if *prdx {
		m := helpers.RecieveMessage()
		paradox(smsMsg, session, m)
	}
	if *prdxTest {
		m := helpers.TestParadoxTest()
		paradox(smsMsg, session, m)
	}
	if *gv {
		m := helpers.RecieveMessage()
		geovision(smsMsg, session, m)
	}
	if *gvTest {
		m := helpers.TestVideoLost()
		geovision(smsMsg, session, m)
	}
	if *testdb {
		log.Println("Inserting test data into DB")
		database.TestPopulate(session)
	}
	if *port {
		portcheck.Startup(session)
	}
	if *portTest {
		log.Println("portTest")
	}
	if *restTest {
		helpers.SendSMSRest("+420725114175", "Manual deepa test")
	}
}

func paradox(sms *structures.SMSMessage, session *mgo.Session, m *email.Message) {
	log.Println("Paradox")

	// Process Paradox message, data in sms.Message struct
	ezs.Paradox(m, sms)

	// Get site object
	site := database.GetSiteBySlug(sms.Site, session)
	sms.SiteName = site.Name

	// Enqueue
	helpers.EnqueueByMsgType(sms, session)
}

func geovision(sms *structures.SMSMessage, session *mgo.Session, m *email.Message) {
	log.Println("Geovision")

	sms.Site = helpers.GetLastElement(m.Header.Get("subject"), "_")
	log.Println("Site: " + sms.Site)

	// Process Geovision message
	cctv.Geovision(m, sms)

	// Get site object
	site := database.GetSiteBySlug(sms.Site, session)
	sms.SiteName = site.Name

	// Enqueue
	helpers.EnqueueByMsgType(sms, session)
}
