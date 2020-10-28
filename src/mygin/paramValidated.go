package main

import (
	"time"
	"github.com/go-playground/validator"
	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required" time_format:"2006-01-02"`
}

var dateCheckValidator validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("dateCheck", dateCheckValidator)
	}
	router.GET("/validateParam", getBookAble)
	router.Run(":9001")

}

func getBookAble(context *gin.Context) {
	var b Booking
	if err := context.ShouldBindWith(&b, binding.Query); err == nil {
		context.JSON(http.StatusOK, gin.H{"message": "booking date are valid!"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
