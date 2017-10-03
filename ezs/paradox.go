package ezs

import (
	"log"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/veqryn/go-email/email"

	"github.com/kardasjan/deepa/helpers"
	"github.com/kardasjan/deepa/structures"
	"github.com/kardasjan/monday"
)

const (
	testEmailMsgType        = "5899ca12dcd5823070194399"
	onMsgType               = "5899ca43dcd582307019439b"
	offMsgType              = "5899ca48dcd582307019439c"
	systemErrorMsgType      = "5899caf3dcd582307019439d"
	moduleErrorMsgType      = "5899caf3dcd582307019439d"
	alarmMsgType            = "5899cafedcd582307019439e"
	webAccessBlockedMsgType = "5899cb07dcd582307019439f"
	unknownMsgType          = "58b45c6a1087ba31fabbd3e9"

	testText         = "Testovaci email"
	onText           = "Zapnuto do ostrahy"
	offText          = "Vypnuto z ostrahy"
	systemText       = "System -"
	poruchaLowerText = "porucha"
	poruchaUpperText = "Porucha"
	moduleText       = "Modul -"
	alertText        = "Poplach v objektu"
	webAccessText    = "Web pristup blokovan"

	messageText   = "Zprava: "
	subsystemText = "Podsystem: "
	timeText      = "Cas: "
	zoneText      = "Zona: "
	fromText      = "z: "
)

// Paradox Bootstrap function
func Paradox(m *email.Message, sms *structures.SMSMessage) {
	lines := strings.Split(string(m.Body[:]), "\n")
	sms.Site = getSite(lines[1])
	log.Println(sms.Site)
	log.Println("Processing...")
	processMessage(lines, sms)
}

// Find out what message is on the input
func processMessage(lines []string, sms *structures.SMSMessage) {
	if strings.Contains(lines[2], testText) {
		testEmail(sms)
	} else if strings.Contains(lines[2], onText) {
		onOff(sms, lines, true)
	} else if strings.Contains(lines[2], offText) {
		onOff(sms, lines, false)
	} else if strings.Contains(lines[2], systemText) || strings.Contains(lines[2], poruchaLowerText) || strings.Contains(lines[2], poruchaUpperText) {
		systemError(sms, lines)
	} else if strings.Contains(lines[2], moduleText) {
		moduleError(sms, lines)
	} else if strings.Contains(lines[2], alertText) {
		alarm(sms, lines)
	} else if strings.Contains(lines[2], webAccessText) {
		webAccessBlocked(sms, lines)
	} else {
		log.Printf("WARN: Message not recognized!")
		unknown(sms, lines)
	}

}

// Process Testovaci email message
func testEmail(sms *structures.SMSMessage) {
	sms.Body = "Testovací zpráva Paradox.\n"
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(testEmailMsgType)
}

// Process Zapnuto/Vypnuto message
func onOff(sms *structures.SMSMessage, lines []string, status bool) {
	sms.Body = helpers.GetLastElement(lines[2], messageText) + "\n"
	sms.Body += helpers.GetLastElement(lines[3], subsystemText) + "\n"
	sms.Body += "Z účtu: " + helpers.GetLastElement(lines[4], fromText) + "\n"
	sms.Datetime = parseDate(helpers.GetLastElement(lines[5], timeText))
	if status {
		sms.MsgType = bson.ObjectIdHex(onMsgType)
	} else {
		sms.MsgType = bson.ObjectIdHex(offMsgType)
	}
}

// Process System Error message
func systemError(sms *structures.SMSMessage, lines []string) {
	sms.Body = helpers.GetLastElement(lines[2], messageText) + "\n"
	sms.Datetime = parseDate(helpers.GetLastElement(lines[3], timeText))
	sms.MsgType = bson.ObjectIdHex(systemErrorMsgType)
}

// Process Module Error message
func moduleError(sms *structures.SMSMessage, lines []string) {
	sms.Body = helpers.GetLastElement(lines[2], messageText) + "\n"
	sms.Body += lines[3] + "\n"
	sms.Datetime = parseDate(helpers.GetLastElement(lines[4], timeText))
	sms.MsgType = bson.ObjectIdHex(moduleErrorMsgType)
}

// Process Poplach message
func alarm(sms *structures.SMSMessage, lines []string) {
	sms.Body = helpers.GetLastElement(lines[2], messageText) + "\n"
	sms.Body += helpers.GetLastElement(lines[3], subsystemText) + "\n"
	sms.Body += helpers.GetLastElement(lines[4], zoneText) + "\n"
	sms.Datetime = parseDate(helpers.GetLastElement(lines[6], timeText))
	sms.MsgType = bson.ObjectIdHex(alarmMsgType)
}

// Process Web přístup blokován message
func webAccessBlocked(sms *structures.SMSMessage, lines []string) {
	// TODO:
	// Receive MSG and save;
	// This content is not proper because it was not gathered from MSG body itself
	sms.Body = "Opakované přihlášení do webu zablokováno!"
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(webAccessBlockedMsgType)
}

func unknown(sms *structures.SMSMessage, lines []string) {
	sms.Body = "Neznámá zpráva Paradox!\n"
	sms.Body += strings.Join(lines, "\n")
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(unknownMsgType)
}

// Get site from which the message was sent
func getSite(line string) string {
	return helpers.GetLastElement(strings.ToLower(line), " ")
}

// Parse Paradox date into time object
func parseDate(dateTime string) time.Time {
	loc, _ := time.LoadLocation("Europe/Prague")
	result, err := monday.ParseInLocation("2 January 2006 15:04", dateTime, loc, monday.LocaleCsCZParadox)
	if err != nil {
		log.Println("Could not parse date!")
		log.Println(err.Error())
		return time.Now()
	}
	return result
}
