package utils

import (
	"bhstore/bhstore-api/user_api/global"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func removeTopStruct(fields map[string]string) map[string]string {
	resp := make(map[string]string)
	for field, err := range fields {
		resp[field[strings.Index(field, ".")+1:]] = err
	}
	return resp
}

func HandleValidator(err error, c *gin.Context) {
	if err == nil {
		return
	}
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  1,
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}
