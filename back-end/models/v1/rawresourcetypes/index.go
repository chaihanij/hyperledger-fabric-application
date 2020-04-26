package rawresourcetypes

import "github.com/chaihanij/hyperledger-fabric-application/back-end/models"

func Index() (rawresourcetypes *models.RawResourceTypes, err error) {
	rawresourcetypes = &mockRawResourceTypes

	return
}
