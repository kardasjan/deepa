package ezs

import (
	"log"
	"strings"

	"github.com/kardasjan/deepa/structures"
	"github.com/veqryn/go-email/email"
)

// IPRS Bootstrap function
func IPRS(m *email.Message, sms *structures.SMSMessage) {
	lines := strings.Split(string(m.Body[:]), "\n")
	log.Println(string(m.Body[:]))
	log.Println(lines)
}
