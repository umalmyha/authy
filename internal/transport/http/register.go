package http

import (
	"net/http"

	"github.com/umalmyha/authy/internal/web/request"
	"github.com/umalmyha/authy/internal/web/response"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := request.BindJSON(r, &req); err != nil {
		response.JSON()
	}
}
