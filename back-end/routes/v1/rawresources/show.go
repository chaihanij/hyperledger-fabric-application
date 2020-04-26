package rawresources

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/hyperledger"
	RawResourcesModel "github.com/chaihanij/hyperledger-fabric-application/back-end/models/v1/rawresources"
)

func Show(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		rawResource, err := RawResourcesModel.Show(clients, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(rawResource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
