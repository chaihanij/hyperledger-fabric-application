package rawresourcetypes

import (
	"encoding/json"
	"net/http"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
	"github.com/gorilla/mux"

	RawResourceTypesModel "github.com/chaihanij/hyperledger-fabric-application/back-end/models/v1/rawresourcetypes"
)

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var opts RawResourceTypesModel.UpdateOpts
		var rawresourcetype models.RawResourceType
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&rawresourcetype); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			opts.Replace = true
		}

		updatedRawResourceType, err := RawResourceTypesModel.Update(id, &rawresourcetype, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(updatedRawResourceType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
