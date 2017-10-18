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
	ip150 := flag.Bool("ip150", false, "Incoming message is from Paradox IP150 module")
	iprs := flag.Bool("iprs", false, "Incoming message is from Paradox IPRS")
	gv := flag.Bool("geovision", false, "Incoming message is from Geovision CCTV software (Supports Analog & Digital)")
	centerv2 := flag.Bool("centerv2", false, "Incoming message is from Geovision CenterV2")
	imggeo := flag.Bool("imggeo", false, "Incoming message is from Geovision and consists image attachments")

	testdb := flag.Bool("testdb", false, "Fills database with test data")
	prdxTest := flag.Bool("paradoxTest", false, "Test paradox functionality")
	gvTest := flag.Bool("gvTest", false, "Test geovision functionality")

	port := flag.Bool("port", false, "Start port checking")             // Needs testing
	portTest := flag.Bool("portTest", false, "Test port functionality") // TODO

	restTest := flag.Bool("restTest", false, "Test REST API")

	flag.Parse()

	// Generate SMS Message object
	smsMsg := &structures.SMSMessage{Retries: 0}

	if *ip150 {
		m := helpers.RecieveMessage()
		runIP150(smsMsg, session, m)
	}
	if *iprs {
		m := helpers.RecieveMessage()
		runIPRS(smsMsg, session, m)
	}
	if *prdxTest {
		m := helpers.TestParadoxTest()
		runIP150(smsMsg, session, m)
	}
	if *gv {
		m := helpers.RecieveMessage()
		runGeovision(smsMsg, session, m)
	}
	if *centerv2 {
		m := helpers.RecieveMessage()
		runCenterV2(smsMsg, session, m)
	}
	if *imggeo {
		m := helpers.RecieveMessage()
		runImggeo(smsMsg, session, m)
	}
	if *gvTest {
		m := helpers.TestVideoLost()
		runGeovision(smsMsg, session, m)
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

func runIP150(sms *structures.SMSMessage, session *mgo.Session, m *email.Message) {
	log.Println("IP150")

	// Process Paradox IP150 message, data in sms.Message struct
	ezs.IP150(m, sms)

	// Get site object
	site := database.GetSiteBySlug(sms.Site, session)
	sms.SiteName = site.Name

	// Enqueue
	helpers.EnqueueByMsgType(sms, session)
}

func runIPRS(sms *structures.SMSMessage, session *mgo.Session, m *email.Message) {
	log.Println("IPRS")
	ezs.IPRS(m, sms)
}

func runCenterV2(sms *structures.SMSMessage, session *mgo.Session, m *email.Message) {
	log.Println("CenterV2")
	cctv.CenterV2(m, sms)
}

func runImggeo(sms *structures.SMSMessage, session *mgo.Session, m *email.Message) {
	log.Println("ImgGeo")
	cctv.Imggeo(m, sms)
}

func runGeovision(sms *structures.SMSMessage, session *mgo.Session, m *email.Message) {
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
