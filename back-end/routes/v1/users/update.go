package users

import (
	"encoding/json"
	"net/http"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
	"github.com/gorilla/mux"

	UsersModel "github.com/chaihanij/hyperledger-fabric-application/back-end/models/v1/users"
)

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var opts UsersModel.UpdateOpts
		var user models.User
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			opts.Replace = true
		}

		updatedUser, err := UsersModel.Update(id, &user, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(updatedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
