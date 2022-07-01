package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword()

func VerifyPassword()

func SignUp()

func Login()

func GetUser()

func GetUsers()
