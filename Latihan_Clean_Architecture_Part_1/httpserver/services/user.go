package services

import (
	"fmt"
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/controllers/views"
	"sesi6-gin/httpserver/repositories/models"

	"golang.org/x/crypto/bcrypt"
)

var Users = []models.User{
	{
		ID:       1,
		Username: "bruh",
		Password: hashPassword("test"),
	},
	{
		ID:       2,
		Username: "rifqi",
		Password: hashPassword("rahasia"),
	},
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func CreateUser(req *params.UserCreateRequest) *views.Response {
	// step : (4) buat model
	var user models.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestError(err)
	}

	user.ID = len(Users) + 1
	user.Password = string(hashedPassword)
	user.Username = req.Username

	Users = append(Users, user)

	// step : (5) kirim ke repositories
	// err = repositories.CreateUser(&model)

	// step : (7) buat sebuah views
	v := views.SuccessCreateResponse(user, "created success!")

	// step : (8) kembalikan views ke controller
	return v
}
func GetUsers() *views.Response {
	v := views.SuccessCreateResponse(parseModelToUserGetAll(&Users), "succesfully get all the users!")
	return v
}

func GetUserByID(id int) *views.Response {
	if id > len(Users) || id < 0 {
		return views.FailedNotFound(fmt.Errorf("There is no user with id %d", id+1))
	}
	user := views.UserGet{
		ID:       id,
		Username: Users[id].Username,
	}
	v := views.SuccessCreateResponse(user, "succesfully get the user!")
	return v
}

func parseModelToUserGetAll(mod *[]models.User) *[]views.UserGet {
	var s []views.UserGet
	for _, st := range *mod {
		s = append(s, views.UserGet{
			ID:       st.ID,
			Username: st.Username,
		})
	}
	return &s
}
