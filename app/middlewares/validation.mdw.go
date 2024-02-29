package middlewares

import (
	"github.com/gin-gonic/gin"
	form "github.com/yusufocaliskan/tiny-go-mvc/framework/form/validate"
	tinyresponse "github.com/yusufocaliskan/tiny-go-mvc/framework/http/response"
)

// Checking if the coming data valid
func Check4ValidData(data interface{}) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		validate := form.FormValidator{}
		response := tinyresponse.Response{Ctx: ctx}

		//1. Binding the incoming data with the struct
		bindingError := ctx.BindJSON(&data)
		if bindingError != nil {
			response.BadWithAbort(bindingError)
		}

		//2. Check if is validated
		validationError := validate.Check(data)
		if validationError != nil {
			response.BadWithAbort(validationError)
		}

		ctx.Next()
	}
}