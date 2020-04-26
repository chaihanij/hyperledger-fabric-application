package users

import (
	"encoding/json"
	"net/http"

	UsersModel "github.com/chaihanij/hyperledger-fabric-application/back-end/models/v1/users"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := UsersModel.Index()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
	}
}
