package cctv

import (
	"log"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/veqryn/go-email/email"

	"github.com/kardasjan/deepa/helpers"
	"github.com/kardasjan/deepa/msgTypes"
	"github.com/kardasjan/deepa/structures"
	"github.com/kardasjan/monday"
)

// Geovision geovision pipe
func Geovision(m *email.Message, sms *structures.SMSMessage) {
	body := m.Body // []byte
	log.Println(string(body))

	if isIntruder(string(body)) {
		intruder(string(body), sms)
	} else if isVideoLost(string(body)) {
		videoLost(string(body), sms)
	} else if isFullDisk(string(body)) {
		fullDisk(string(body), sms)
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
	sms.MsgType = bson.ObjectIdHex(msgTypes.GeovisionIntruder)
}

func videoLost(body string, sms *structures.SMSMessage) {
	camera := helpers.GetStringInBetween(body, "Lost Video Camera :\n", "\non ")
	time := helpers.GetStringInBetween(body, "\non ", " Stredn")

	sms.Body = "Ztráta videa!\n"
	sms.Body += "Kamera: " + camera + "\n"
	sms.Datetime = parseDate(time)
	sms.MsgType = bson.ObjectIdHex(msgTypes.GeovisionVideoLost)
}

func fullDisk(body string, sms *structures.SMSMessage) {
	sms.Body = "Disk je plný!\n"
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(msgTypes.GeovisionFullDisk)
}

func unknown(body string, sms *structures.SMSMessage) {
	sms.Body = "Neznámá zpráva GV!\n"
	sms.Body += body
	sms.Datetime = time.Now()
	sms.MsgType = bson.ObjectIdHex(msgTypes.GeovisionUnknown)
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
