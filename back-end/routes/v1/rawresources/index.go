package rawresources

import (
	"encoding/json"
	"net/http"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/hyperledger"
	RawResourcesModel "github.com/chaihanij/hyperledger-fabric-application/back-end/models/v1/rawresources"
)

func Index(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rawresources, err := RawResourcesModel.Index(clients)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(rawresources)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
	}
}
