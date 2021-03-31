package controllers

import (
	"dasargo/Models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, user)
	}
}

func CreateUser(context *gin.Context) {
	 var user models.User
	 context.BindJSON(&user)
	 err := models.CreateUser(&user)
	 if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, user)
	}
}

func GetUserByID(context *gin.Context) {
	id := context.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, user)
	}
}

func UpdateUser(context *gin.Context) {

	var user models.User
	id := context.Params.ByName("id")
	err := models.GetUserByID(&user, id)

	if err != nil {
	 context.JSON(http.StatusNotFound, user)
	}
	
	context.BindJSON(&user)
	
	err = models.UpdateUser(&user, id)
	
	if err != nil {
	 context.AbortWithStatus(http.StatusNotFound)
	} else {
	 context.JSON(http.StatusOK, user)
	}
	
}

func DeleteUser(context *gin.Context) {
	var user models.User

	id := context.Params.ByName("id")
	
	err := models.DeleteUser(&user, id)
	
	if err != nil {
	 context.AbortWithStatus(http.StatusNotFound)
	} else {
	 context.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}