package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/hyperledger"
	V1Router "github.com/chaihanij/hyperledger-fabric-application/back-end/routes/v1"
)

const (
	staticDir = "/static/"
)

type Service interface {
	GetRawRouter() *mux.Router
}

func GetRouter() Service {
	r := Router{
		RawRouter: mux.NewRouter().StrictSlash(true),
	}

	configPath := os.Getenv("HYPERLEDGER_CONFIG_PATH")

	if len(configPath) == 0 {
		panic("ENV var 'HYPERLEDGER_CONFIG_PATH' is not set. unable to connect to network")
	}

	clients := hyperledger.NewClientMap(
		"test-network",
		configPath,
	)

	_, err := clients.AddClient(
		"Admin",
		"org1",
		"mainchannel",
	)
	if err != nil {
		panic(err)
	}

	r.RawRouter.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	for _, route := range GetRoutes() {
		r.RawRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	for name, pack := range V1Router.GetRoutes(clients) {
		fmt.Println(name)
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}

	return r
}
