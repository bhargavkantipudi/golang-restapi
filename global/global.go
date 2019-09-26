package global

import "github.com/go-chi/jwtauth"

var (
	tokenAuth *jwtauth.JWTAuth
	key       string
)
