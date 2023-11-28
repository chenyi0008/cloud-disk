package define

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}
