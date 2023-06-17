package controller

import (
	"crud-api-go/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("Teste")
	c.JSON(err.Code, err)
}

func FindUserByEmai(c *gin.Context) {

}
