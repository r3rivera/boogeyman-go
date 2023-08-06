package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/r3rivera/boogeyman/services/udetails"
)

type UserAddress struct {
	Street1 string `json:"street1" binding:"required"`
	Street2 string `json:"street2" binding:"required"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	Zip     string `json:"zip" binding:"required"`
}

type UserInfo struct {
	Firstname string      `json:"firstname" binding:"required"`
	Lastname  string      `json:"lastname" binding:"required"`
	Dob       string      `json:"dob"`
	Address   UserAddress `json:"address"`
}

func CreateUserProfile(c *gin.Context) {
	var userInfo UserInfo

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request"})
		return
	}

	claims := ExtractClaims(c)
	email := claims["sub"]
	err := createUserDetails(email.(string), &userInfo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
	})
}

func createUserDetails(email string, userInfo *UserInfo) error {

	var bdate time.Time
	if userInfo.Dob != "" {
		dob, err := time.Parse("01-02-06", userInfo.Dob)
		if err != nil {
			return errors.New("Invalid Date Format")
		}
		bdate = dob
	}

	uDetail := &udetails.UserDetail{
		Firstname: userInfo.Firstname,
		Lastname:  userInfo.Lastname,
		Dob:       bdate,
	}
	uAddress := &udetails.UserAddress{
		Street1: userInfo.Address.Street1,
		Street2: userInfo.Address.Street2,
		City:    userInfo.Address.City,
		State:   userInfo.Address.State,
		Zip:     userInfo.Address.Zip,
	}

	uInfo := udetails.NewUserDetail(email, *uDetail, *uAddress)
	return uInfo.WriteDB()
}

func GetUserDetails(c *gin.Context) {
	claims := ExtractClaims(c)
	email := claims["sub"].(string)

	userDetail := udetails.NewUserDetail(email, udetails.UserDetail{}, udetails.UserAddress{})
	user, err := userDetail.ReadDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}
