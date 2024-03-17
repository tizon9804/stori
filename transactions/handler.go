package transactions

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TransactionFileHandler
// @Summary upload csv file of transactions
// @Description upload transactions for a user
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param file formData file true "File to upload"
// @Param  email formData string true "email"
// @Success 204
// @failure 400 {object} error
// @response default {object} string
// @Header 200 {string} Location
// @Router /Transaction/upload [post]
func TransactionFileHandler(transaction Transaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Request.FormValue("email")
		if email == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("email is required"))
			return
		}
		file, _, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		defer file.Close()

		if err = transaction.ProcessFile(file, email); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, nil)

	}
}
