package users

import (
	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
)

var mockUsers models.Users

func init() {
	usr, _ := models.NewUser("Nick", "Kotenberg", "nick@mail.com", "1234")

	mockUsers = models.Users{
		*usr,
	}
}
