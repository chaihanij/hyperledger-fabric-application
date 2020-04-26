package rawresourcetypes

import (
	"net/http"

	"github.com/gorilla/mux"

	RawResourceTypesModel "github.com/chaihanij/hyperledger-fabric-application/back-end/models/v1/rawresourcetypes"
)

func Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		if err := RawResourceTypesModel.Destroy(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
	}
}
