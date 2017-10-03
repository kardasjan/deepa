package portcheck

import (
	"time"

	"github.com/kardasjan/deepa/database"

	mgo "gopkg.in/mgo.v2"
)

const (
	routerMsgType = ""
)

// Startup ...
func Startup(session *mgo.Session) {
	sites := database.GetAllSites(session)
	for _, site := range sites {
		// Router Check
		s := ping(site.IP, string(site.Router.Port))
		if site.Router.Status == 2 { // Status Disabled
			if s {
				site.Router.LastActive = time.Now()
			}
			updateRouter(&site.Router, site.ID, session)
			continue
		}
		if s {
			if site.Router.Status == 0 {
				sendRouter("Internet je OK!", &site, session)
				site.Router.Status = 1
			}
			site.Router.LastActive = time.Now()
			site.Router.RetryCount = 0
			updateRouter(&site.Router, site.ID, session)
		}
		if !s {
			site.Router.RetryCount++
			if site.Router.RetryCount == site.Router.MaxRetries {
				sendRouter("Internet nefunguje - Není možno dohledovat!", &site, session)
			}
			continue
		}

		for _, service := range site.Services {
			// Get service status true/false
			s := ping(site.IP, string(service.Port))
			if service.Status == 2 { // Status Disabled
				if s {
					service.LastActive = time.Now()
				}
				continue
			}
			if s {
				if service.Status == 0 {
					sendService("Služba "+service.Name+" je OK!", &service, &site, session)
					service.Status = 1
				}
				service.LastActive = time.Now()
				service.RetryCount = 0
				updateService(&service, session)
				continue
			}
			if !s {
				service.RetryCount++
				if service.RetryCount == service.MaxRetries {
					sendService("Služba "+service.Name+" neodpovídá!", &service, &site, session)
				}
				updateService(&service, session)
			}
		}
	}
}
