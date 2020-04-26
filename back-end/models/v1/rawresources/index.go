package rawresources

import (
	"encoding/json"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/hyperledger"
	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
)

func Index(clients *hyperledger.Clients) (rawresources *models.RawResources, err error) {
	rawresources = new(models.RawResources)

	res, err := clients.Query("org1", "rawresources", "queryString", [][]byte{
		[]byte("{\"selector\":{ \"visible\": { \"$eq\":true } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, rawresources); err != nil {
		return
	}

	return
}
