package rawresourcetypes

import (
	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
)

var mockRawResourceTypes models.RawResourceTypes

func init() {
	iron, _ := models.NewRawResourceType("Iron")
	copper, _ := models.NewRawResourceType("Copper")
	platinum, _ := models.NewRawResourceType("Platinum")

	mockRawResourceTypes = models.RawResourceTypes{
		*iron,
		*copper,
		*platinum,
	}
}
