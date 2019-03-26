package utils

import (
	"strings"

	"github.com/kardasjan/deepa/structures"
	"github.com/kardasjan/monday"
)

// GetStringInBetween Returns empty string if no start string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	return str[s:e]
}

// GetLastElement Return last element of a string split by delimeter
func GetLastElement(str string, delimeter string) (last string) {
	split := strings.Split(str, delimeter)
	last = split[len(split)-1]
	return
}

// PrepareSMS Add all the parameters to the Body of a Message
func PrepareSMS(sms *structures.SMSMessage) {
	dt := monday.Format(sms.Datetime, "2. January 2006 15:04", monday.LocaleCsCZ)
	sms.Body += dt + "\n"
	sms.Body += sms.SiteName
}

// PrepareSMSWithEmail ...
func PrepareSMSWithEmail(sms *structures.SMSMessage) {
	dt := monday.Format(sms.Datetime, "2. January 2006 15:04", monday.LocaleCsCZ)
	sms.Body += dt + "\n"
	sms.Body += "Obrázky zaslány emailem" + "\n"
	sms.Body += sms.SiteName
}
