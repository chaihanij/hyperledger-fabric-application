package rawresources

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/chaihanij/hyperledger-fabric-application/back-end/hyperledger"
	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
)

func Store(clients *hyperledger.Clients, name string, typeID string, weight int, arrivalTime *time.Time) (rawresource *models.RawResource, err error) {
	rawresource = new(models.RawResource)

	rawresource.ID = uuid.NewV4().String()
	rawresource.Name = name
	rawresource.TypeID = typeID
	rawresource.Weight = weight

	if arrivalTime != nil {
		rawresource.ArrivalTime = arrivalTime
	}

	packet, err := json.Marshal(rawresource)
	if err != nil {
		return
	}

	if _, err = clients.Invoke("org1", "rawresources", "store", [][]byte{
		packet,
	}); err != nil {
		return
	}

	return
}
