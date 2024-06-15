package storage

import "golang-auth/models"

var users = map[string]models.User{}

func GetUser(username string) (models.User, bool) {
	user, exists := users[username]
	return user, exists
}

func AddUser(user models.User) {
	users[user.Username] = user
}
