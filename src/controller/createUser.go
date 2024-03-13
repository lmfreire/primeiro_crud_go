package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lmfreire/primeiro_crud_go/src/configuration/validation"
	"github.com/lmfreire/primeiro_crud_go/src/controller/model/request"
	"github.com/lmfreire/primeiro_crud_go/src/controller/model/response"
)

func CreateUser(c *gin.Context) {
	
	var userRequest request.UserRequest
	
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are some incorrect fields, error=%s", err.Error())
		
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID: "teste",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	c.JSON(200, response)

}