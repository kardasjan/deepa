package cctv

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
	videolostMsgType = "58a9fdda1087ba06c1615e07"
	intruderMsgType  = "58a9fdda1087ba06c1615e08"
	fulldiskMsgType  = "58a9fe5e1087ba0a24ab844d"
	unknownMsgType   = "58b45c6a1087ba31fabbd3ea"

	intruderMessage  = "Intruder detected"
	videolostMessage = "Video Lost"
	fulldiskMessage  = "Disk storage space low detected."
)

// Geovision geovision pipe
func Geovision(m *email.Message, sms *structures.SMSMessage) {
	processMessage(m.Body, sms)
}

// Find out what message is on the input
func processMessage(body []byte, sms *structures.SMSMessage) {
	if strings.Contains(string(body), intruderMessage) {
		log.Println("Intruder!")
		intruder(string(body), sms)
	} else if strings.Contains(string(body), videolostMessage) {
		log.Println("Video Lost!")
		videolost(string(body), sms)
	} else if strings.Contains(string(body), fulldiskMessage) {
		log.Println("Full disk!")
		fulldisk(string(body), sms)
	} else {
		log.Printf("WARN: Message not recognized!")
		unknown(string(body), sms)
	}

}

func intruder(body string, sms *structures.SMSMessage) {
	camera := helpers.GetStringInBetween(body, "Information:", "-Intruder detected.")
	time := helpers.GetStringInBetween(body, "\non ", " Stredn")

	sms.Body = "Detekován narušitel!\n"
	sms.Body += "Kamera: " + camera + "\n"
	sms.Datetime = parseDate(time)
	sms.MsgType = bson.ObjectIdHex(intruderMsgType)
}

func videolost(body string, sms *structures.SMSMessage) {
	camera := helpers.GetStringInBetween(body, "Lost Video Camera :\n", "\non ")
	time := helpers.GetStringInBetween(body, "\non ", " Stredn")

	sms.Body = "Ztráta videa!\n"
	sms.Body += "Kamera: " + camera + "\n"
	sms.Datetime = parseDate(time)
	sms.MsgType = bson.ObjectIdHex(videolostMsgType)
}

func fulldisk(body string, sms *structures.SMSMessage) {
	sms.Body = "Disk je plný!\n"
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(fulldiskMsgType)
}

func unknown(body string, sms *structures.SMSMessage) {
	sms.Body = "Neznámá zpráva GV!\n"
	sms.Body += body
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(unknownMsgType)
}

// Parse Geovision date into time object
func parseDate(dateTime string) time.Time {
	parts := strings.SplitAfter(dateTime, ", ")
	loc, _ := time.LoadLocation("Europe/Prague")
	result, err := monday.ParseInLocation("02 Jan 06 15:04:05", parts[1], loc, monday.LocaleEnGB)
	if err != nil {
		log.Fatal(err)
		return time.Now()
	}
	return result
}
