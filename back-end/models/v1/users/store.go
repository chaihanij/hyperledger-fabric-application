package users

import "github.com/chaihanij/hyperledger-fabric-application/back-end/models"

func Store(firstName string, lastName string, email string, password string) (user *models.User, err error) {
	user, err = models.NewUser(firstName, lastName, email, password)
	if err != nil {
		return
	}

	mockUsers = append(mockUsers, *user)

	return
}
