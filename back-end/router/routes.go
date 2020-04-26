package router

import (
	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
	HomeHandler "github.com/chaihanij/hyperledger-fabric-application/back-end/routes/home"
	StatusHandler "github.com/chaihanij/hyperledger-fabric-application/back-end/routes/status"
)

func GetRoutes() models.Routes {
	return models.Routes{
		models.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: HomeHandler.Index},
		models.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: StatusHandler.Index},
	}
}
