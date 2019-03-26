package src

import (
	"time"

	"github.com/kardasjan/deepa/models"

	"gopkg.in/mgo.v2/bson"
)

type Technology interface {
	Recieved()
	CheckMessageType() string
}

type SMS struct {
	To         []string
	Body       string
	Technology string
	Site       *models.Site
	Datetime   time.Time
	MsgType    bson.ObjectId
}
