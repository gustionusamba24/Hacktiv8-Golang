package services

import (
	"math/rand"
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/controllers/views"
	"sesi6-gin/httpserver/repositories/models"

	"golang.org/x/crypto/bcrypt"
)


func CreateUser(req *params.UserCreateRequest) *views.Response {
	// step : (4) buat model
	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestError(err)
	}

	user.ID = rand.Intn(100000)
	user.Password = string(hash)
	user.Username = req.Username
	
	user.SaveUser()
	// step : (5) kirim ke repositories
	// err = repositories.CreateUser(&model)

	// step : (7) buat sebuah views
	v := views.SuccessCreateResponse(user, "created success!")
	// step : (8) kembalikan views ke controller
	return v
}

func ReadUser() *views.Response {
	data := models.SaveData
	v := views.SuccessCreateResponse(data, "created success!")

	// step : (8) kembalikan views ke controller
	return v
}