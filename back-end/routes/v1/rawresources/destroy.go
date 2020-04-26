package rawresources

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/hyperledger"
	RawResourcesModel "github.com/chaihanij/hyperledger-fabric-application/back-end/models/v1/rawresources"
)

func Destroy(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		if err := RawResourcesModel.Destroy(clients, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
	}
}
