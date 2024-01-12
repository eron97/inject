package validate

import (
	"net/http"

	"github.com/eron97/inject.git/config/error_messages"
	"github.com/eron97/inject.git/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	RequestOk models.CreateUser
	Validate  = validator.New()
	transl    ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateRequest(c *gin.Context) {

	if err := c.ShouldBindJSON(&RequestOk); err != nil {

		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessage := error_messages.NewBadRequestError("Erro de validação")

			for _, fieldError := range ve {
				cause := error_messages.CausesMessages{
					Field:   fieldError.Field(),
					Message: fieldError.Translate(transl),
				}
				errorMessage.Causes = append(errorMessage.Causes, cause)
			}

			c.JSON(http.StatusBadRequest, gin.H{"Erro de validação": errorMessage})
			c.Abort()
			return
		}

		errorMessage := error_messages.NewUnmarshalError("Campos preenchidos não correspondem aos tipos atribuídos")
		c.JSON(http.StatusBadRequest, errorMessage)
		c.Abort()
		return
	}

	c.Set("createUserRequest", &RequestOk)
	c.Next()
}

/*

func ValidateRequest(c *gin.Context) {

	var request models.CreateUser

	if err := c.ShouldBindJSON(&request); err != nil {

		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessage := error_messages.NewBadRequestError("Erro de validação")

			for _, fieldError := range ve {
				cause := error_messages.CausesMessages{
					Field:   fieldError.Field(),
					Message: fieldError.Translate(transl),
				}
				errorMessage.Causes = append(errorMessage.Causes, cause)
			}

			c.JSON(http.StatusBadRequest, gin.H{"Erro de validação": errorMessage})
			c.Abort()
			return
		}

		errorMessage := error_messages.NewUnmarshalError("Campos preenchidos não correspondem aos tipos atribuídos")
		c.JSON(http.StatusBadRequest, errorMessage)
		c.Abort()
		return
	}

	c.Set("createUserRequest", &request)
	c.Next()
}


*/
