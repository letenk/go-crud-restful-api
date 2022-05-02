package middleware

import (
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Check whether the request header `X-API-Key` same with "RAHASIA"
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// Yes, Next process
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// No, resonse error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)

	}
}
