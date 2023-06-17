package controller

import (
	"crud-api-go/src/configuration/rest_err"
	"crud-api-go/src/controller/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorret filds, error: %s", err))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}
