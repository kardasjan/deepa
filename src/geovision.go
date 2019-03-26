package src

import (
	"strings"
	"time"

	"github.com/kardasjan/deepa/helpers"
	"github.com/kardasjan/deepa/models"
	"github.com/kardasjan/deepa/utils"
	"github.com/kardasjan/monday"
	email "github.com/veqryn/go-email/email"
	"gopkg.in/mgo.v2/bson"
)

const (
	intruderMessage  = "Intruder detected"
	videoLostMessage = "Video Lost"
	fullDiskMessage  = "Disk storage space low detected."

	videoLostMsgType = bson.ObjectId("58a9fdda1087ba06c1615e07")
	intruderMsgType  = bson.ObjectId("58a9fdda1087ba06c1615e08")
	fullDiskMsgType  = bson.ObjectId("58a9fe5e1087ba0a24ab844d")
	unknownMsgType   = bson.ObjectId("58b45c6a1087ba31fabbd3ea")
)

// Geovision Create this struct if the technology uset is Geovision
type Geovision struct {
	Email email.Message
	Sms   *SMS
}

// Received Trigger this function upon email receive, it will process the email and compose SMS message
func (t Geovision) Received() (err error) {

	// If site does not exist then there is no point going further
	slug := utils.GetLastElement(t.Email.Header.Get("subject"), "_")
	site, err := models.GetSiteBySlug(slug)
	if err != nil {
		return err
	}
	t.Sms.Site = site

	bodyString := string(t.Email.Body)

	switch t.CheckMessageType() {
	case intruderMsgType:
		err = composeIntruderMessage(t.Sms, bodyString)
	case videoLostMsgType:
		err = composeVideoLostMessage(t.Sms, bodyString)
	case fullDiskMsgType:
		err = composeFullDiskMessage(t.Sms, bodyString)
	default:
		err = composeUnknownMessage(t.Sms, bodyString)
	}

	t.Sms.Technology = "Geovision"
	return
}

// CheckMessageType Check which message it is
func (t Geovision) CheckMessageType() bson.ObjectId {
	bodyString := string(t.Email.Body)
	if strings.Contains(bodyString, intruderMessage) {
		return intruderMsgType
	} else if strings.Contains(bodyString, videoLostMessage) {
		return videoLostMsgType
	} else if strings.Contains(bodyString, fullDiskMessage) {
		return fullDiskMsgType
	} else {
		return unknownMsgType
	}
}

// Local functions to create the SMS
func composeIntruderMessage(sms *SMS, body string) (err error) {
	camera := helpers.GetStringInBetween(body, "Information:", "-Intruder detected.")
	time := helpers.GetStringInBetween(body, "\non ", " Stredn")
	sms.Datetime, err = utils.ParseGeovisionDate(time)
	formattedDate := monday.Format(sms.Datetime, "2. January 2006 15:04", monday.LocaleCsCZ)

	sms.To = models.GetNumbersWhoShouldReceiveMessage(sms.Site, intruderMsgType)

	sms.Body = "Detekován narušitel!\n"
	sms.Body += "Kamera: " + camera + "\n"
	sms.Body += formattedDate + "\n"
	sms.Body += sms.Site.Name
	return
}

func composeVideoLostMessage(sms *SMS, body string) (err error) {
	camera := helpers.GetStringInBetween(body, "Lost Video Camera :\n", "\non ")
	time := helpers.GetStringInBetween(body, "\non ", " Stredn")
	sms.Datetime, err = utils.ParseGeovisionDate(time)
	formattedDate := monday.Format(sms.Datetime, "2. January 2006 15:04", monday.LocaleCsCZ)

	sms.To = models.GetNumbersWhoShouldReceiveMessage(sms.Site, videoLostMsgType)

	sms.Body = "Ztráta videa!\n"
	sms.Body += "Kamera: " + camera + "\n"
	sms.Body += formattedDate + "\n"
	sms.Body += sms.Site.Name
	return
}

func composeFullDiskMessage(sms *SMS, body string) (err error) {
	sms.Datetime = time.Now()
	formattedDate := monday.Format(sms.Datetime, "2. January 2006 15:04", monday.LocaleCsCZ)

	sms.To = models.GetNumbersWhoShouldReceiveMessage(sms.Site, fullDiskMsgType)

	sms.Body = "Disk je plný!\n"
	sms.Body += formattedDate + "\n"
	sms.Body += sms.Site.Name
	return
}

func composeUnknownMessage(sms *SMS, body string) (err error) {
	sms.Datetime = time.Now()
	formattedDate := monday.Format(sms.Datetime, "2. January 2006 15:04", monday.LocaleCsCZ)

	sms.To = models.GetNumbersWhoShouldReceiveMessage(sms.Site, unknownMsgType)

	sms.Body = "Neznámá zpráva GV!\n"
	sms.Body += body + "\n"
	sms.Body += formattedDate + "\n"
	sms.Body += sms.Site.Name
	return
}
