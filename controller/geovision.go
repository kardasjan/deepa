package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kardasjan/deepa/src"
	"github.com/kardasjan/deepa/utils"
	email "github.com/veqryn/go-email/email"
)

// DecodeGeovision godoc
// @Summary Add person to DB
// @Description Add new person from body
// @Tags people
// @Accept  json
// @Produce  json
// @Param person body models.Person true "Person to add"
// @Success 200 {object} models.Person
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /people [post]
func (c *Controller) DecodeGeovision(ctx *gin.Context) {

	// veqryn\email message struct
	email := email.Message{}

	if err := ctx.ShouldBindJSON(&email); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	msg := src.Geovision{Email: email}
	err := msg.Received() // Trigger
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, msg.Sms)
}
