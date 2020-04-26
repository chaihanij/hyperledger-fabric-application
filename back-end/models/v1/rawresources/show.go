package rawresources

import (
	"encoding/json"
	"errors"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/hyperledger"
	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
)

func Show(clients *hyperledger.Clients, id string) (rawresource *models.RawResource, err error) {

	rawresources := new(models.RawResources)

	res, err := clients.Query("org1", "rawresources", "queryString", [][]byte{
		[]byte("{\"selector\":{ \"id\": { \"$eq\":\"" + id + "\" } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, rawresources); err != nil {
		return
	}

	list := *rawresources

	if len(list) == 0 {
		err = errors.New("unable to find rawresource with id " + id)
		return
	}

	rawresource = &list[0]

	return
}
