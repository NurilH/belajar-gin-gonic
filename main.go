package main

import (
	"github.com/gin-gonic/gin"

	users "belajar-gin-gonic/users/delivery/http"
)

func main() {
	router := gin.Default()

	users.UsersNewDelivery()

	router.Run(":8000")
}
