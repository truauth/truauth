package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	lib "github.com/truauth/truauth/pkg/auth-lib"
)

// CheckResponseError checks if a response error is present and handles it
func CheckResponseError(
	ctx *gin.Context,
	err *lib.AuthErrorResponse,
	redirectURI string,
	state string,
) bool {
	if err != nil {
		err.State = state

		if redirectURI == "" {
			ctx.JSON(http.StatusUnauthorized, err)

			return true
		}

		redirect := fmt.Sprintf("%s?error=%s&error_description=%s&state=%s",
			redirectURI,
			err.Error,
			err.ErrorDescription,
			err.State)

		http.Redirect(ctx.Writer, ctx.Request, redirect, http.StatusFound)

		return true
	}

	return false
}
