package database

import (
	"log"
	"time"

	"github.com/kardasjan/deepa/structures"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var contacts []structures.Contact

// TestPopulate Fills database with test data
func TestPopulate(session *mgo.Session) {
	clearDatabase(session)
	populateMsgTypes(session)
	populateContacts(session)
	kotorskaData(session)
	makovskehoData(session)
	ustaniceData(session)
	lekarnaMakData(session)
	vokacovaData(session)
	nevanovaData(session)
	poznanska1Data(session)
	poznanska2Data(session)
	lodzskaData(session)
	pejevoveData(session)
	travnickovaData(session)
	kosmickaData(session)
	pertoldovaData(session)
	levskehoData(session)
	mendelovaData(session)
	ciolkovskehoData(session)
	ohradniData(session)
	olstynskaData(session)
	famfulikovaData(session)
}

func populateContacts(session *mgo.Session) {
	contacts = append(contacts, makeContact("Jan Kardaš", "+420725114175", "jan.kardas@seznam.cz"))
	contacts = append(contacts, makeContact("Jiří Kardaš", "+420702014448", "jiri.kardas.jk@gmail.com"))

	c := session.DB(dbName).C("contacts")
	for _, contact := range contacts {
		err := c.Insert(contact)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func lekarnaMakData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Lékárna Makovského",
		IP:       "217.195.168.133",
		Slug:     "lekarnamak",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func vokacovaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Vokáčova 1180",
		IP:       "193.165.81.150",
		Slug:     "vokacova1180",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func nevanovaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Nevanova 1069",
		IP:       "109.239.69.227",
		Slug:     "nevanova1069",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func poznanska1Data(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Poznaňska 437",
		IP:       "193.165.118.122",
		Slug:     "poznanska437",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func poznanska2Data(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	//recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Poznaňská 428",
		IP:       "193.165.127.6",
		Slug:     "poznanska428",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func lodzskaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Lodžská 466",
		IP:       "128.0.183.218",
		Slug:     "lodzska466",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func pejevoveData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Pejevové 3121",
		IP:       "217.195.170.53",
		Slug:     "pejevove3121",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func travnickovaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Trávníčkova 1762",
		IP:       "217.195.169.177",
		Slug:     "travnickova1762",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func kosmickaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Kosmická 749",
		IP:       "217.195.169.129",
		Slug:     "kosmicka749",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func pertoldovaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Pertoldova 3380",
		IP:       "217.195.169.109",
		Slug:     "pertoldova3380",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func levskehoData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Levského 3201",
		IP:       "217.195.173.117",
		Slug:     "levskeho3201",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func mendelovaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Mendelova 542",
		IP:       "80.243.111.217",
		Slug:     "mendelova542",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func ciolkovskehoData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Ciolkovského 853",
		IP:       "80.243.111.29",
		Slug:     "ciolkovskeho853",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func ohradniData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Ohradní 1335",
		IP:       "193.165.56.98",
		Slug:     "ohradni1335",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func olstynskaData(session *mgo.Session) {
	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Famfulíkova 1140",
		IP:       "193.165.144.154",
		Slug:     "famfulikova1140",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func famfulikovaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "TP-Link",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Olštýnská 607",
		IP:       "217.195.169.149",
		Slug:     "olstynska607",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func ustaniceData(session *mgo.Session) {
	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "Mikrotik",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "U stanice 83",
		IP:       "94.142.237.66",
		Slug:     "ustanice83",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}

}

func makovskehoData(session *mgo.Session) {
	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	// Receive Jiří Kardaš
	var recieveTwo []bson.ObjectId
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "Mikrotik",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Makovského 1140",
		IP:       "80.243.111.177",
		Slug:     "makovskeho1140",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func kotorskaData(session *mgo.Session) {

	// Receive Jan Kardaš
	var recieveOne []bson.ObjectId
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca12dcd5823070194399")) // Test message
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca43dcd582307019439b")) // Zapnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899ca48dcd582307019439c")) // Vypnuto
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899caf3dcd582307019439d")) // Porucha
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cafedcd582307019439e")) // Poplach
	recieveOne = append(recieveOne, bson.ObjectIdHex("5899cb07dcd582307019439f")) // Web pristup
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d")) // Plný disk GV
	recieveOne = append(recieveOne, bson.ObjectIdHex("58a9fdda1087ba06c1615e08")) // Narusitel
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9")) // Neznámá Paradox
	recieveOne = append(recieveOne, bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea")) // Neznámá GV

	// Receive Jiří Kardaš
	var recieveTwo []bson.ObjectId
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899ca12dcd5823070194399"))
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899ca43dcd582307019439b"))
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("5899ca48dcd582307019439c"))
	recieveTwo = append(recieveTwo, bson.ObjectIdHex("58a9fdda1087ba06c1615e07")) // VideoLost

	var siteContacts []structures.SiteContact
	siteContacts = append(siteContacts, makeSiteContact(contacts[0], recieveOne))
	siteContacts = append(siteContacts, makeSiteContact(contacts[1], recieveTwo))

	router := structures.Router{
		Status:      1,
		RetryCount:  0,
		MaxRetries:  3,
		Description: "Mikrotik",
		CreatedAt:   time.Now(),
	}

	site := &structures.Site{
		ID:       bson.NewObjectId(),
		Name:     "Kotorská 1752",
		IP:       "193.165.81.174",
		Slug:     "kotorska1752",
		Status:   1,
		Router:   router,
		Contacts: siteContacts,
	}

	c := session.DB(dbName).C("sites")
	err := c.Insert(site)
	if err != nil {
		log.Fatal(err)
	}
}

func populateMsgTypes(session *mgo.Session) {

	var msgTypes [11]structures.MessageType

	// Paradox
	msgTypes[0] = makeMessageType(bson.ObjectIdHex("5899ca12dcd5823070194399"), "Testovací email", "Paradox")
	msgTypes[1] = makeMessageType(bson.ObjectIdHex("5899ca43dcd582307019439b"), "Zapnuto do ostrahy", "Paradox")
	msgTypes[2] = makeMessageType(bson.ObjectIdHex("5899ca48dcd582307019439c"), "Vypnuto z ostrahy", "Paradox")
	msgTypes[3] = makeMessageType(bson.ObjectIdHex("5899caf3dcd582307019439d"), "Porucha", "Paradox")
	msgTypes[4] = makeMessageType(bson.ObjectIdHex("5899cafedcd582307019439e"), "Poplach v objektu", "Paradox")
	msgTypes[5] = makeMessageType(bson.ObjectIdHex("5899cb07dcd582307019439f"), "Web přístup blokován", "Paradox")
	msgTypes[6] = makeMessageType(bson.ObjectIdHex("58b45c6a1087ba31fabbd3e9"), "Neznámá zpráva", "Paradox")

	// Geovision
	msgTypes[7] = makeMessageType(bson.ObjectIdHex("58a9fdda1087ba06c1615e07"), "Ztráta videa", "Geovision")
	msgTypes[8] = makeMessageType(bson.ObjectIdHex("58a9fdda1087ba06c1615e08"), "Narušitel", "Geovision")
	msgTypes[9] = makeMessageType(bson.ObjectIdHex("58a9fe5e1087ba0a24ab844d"), "Plný disk", "Geovision")
	msgTypes[10] = makeMessageType(bson.ObjectIdHex("58b45c6a1087ba31fabbd3ea"), "Neznámá zpráva", "Geovision")

	c := session.DB(dbName).C("messageTypes")
	for _, msgType := range msgTypes {
		err := c.Insert(msgType)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func makeMessageType(objectID bson.ObjectId, name string, technology string) structures.MessageType {
	object := structures.MessageType{
		ID:         objectID,
		Name:       name,
		Technology: technology,
	}
	return object
}

func makeSiteContact(contact structures.Contact, recieve []bson.ObjectId) structures.SiteContact {
	object := structures.SiteContact{
		ContactID:    contact.ID,
		MessageTypes: recieve,
	}
	return object
}

func makeContact(name string, phone string, email string) structures.Contact {
	object := structures.Contact{
		ID:    bson.NewObjectId(),
		Name:  name,
		Phone: phone,
		Email: email,
	}
	return object
}

func clearDatabase(session *mgo.Session) {
	names, _ := session.DB("watch").CollectionNames()

	if stringInSlice("sites", names) {
		c := session.DB(dbName).C("sites")
		err := c.DropCollection()
		if err != nil {
			panic(err)
		}
	}

	if stringInSlice("messageTypes", names) {
		c := session.DB(dbName).C("messageTypes")
		err := c.DropCollection()
		if err != nil {
			panic(err)
		}
	}

	if stringInSlice("contacts", names) {
		c := session.DB(dbName).C("contacts")
		err := c.DropCollection()
		if err != nil {
			panic(err)
		}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
