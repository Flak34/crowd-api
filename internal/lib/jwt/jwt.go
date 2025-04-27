package jwt

import (
	"github.com/Flak34/crowd-api/internal/user/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
	"time"
)

const (
	tokenSecret = "KakJeOnGoriachpedjgerjgkr;kgekropgenagnrnteigne"
	tokenTTL    = time.Hour * 24
)

func GenerateToken(user user_model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(tokenTTL).Unix()
	claims["roles"] = lo.Map(user.Roles, func(role user_model.Role, _ int) string {
		return role.Name
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
