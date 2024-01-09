package middlewares

import (
	"my-task-app/app/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(configs.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

// Generate token jwt
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	//yang tulisannya hijau adalah data payload yang disisipkan ke sebuah token
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.JWT_SECRET))

}

// extract token jwt
func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		//yang tulisannya hijau adalah data payload yang diextract dari sebuah token
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
