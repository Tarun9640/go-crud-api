package storage


import "crud-go/models"

// users slice acts like a fake database
var Users = []models.User{
	{ID: 1, Name: "Tarun", Email: "tarun@gmail.com"},
	{ID: 2, Name: "Rahul", Email: "rahul@gmail.com"},
}