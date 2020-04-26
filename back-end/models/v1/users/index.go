package users

import "github.com/chaihanij/hyperledger-fabric-application/back-end/models"

func Index() (users *models.Users, err error) {
	users = &mockUsers

	return
}
