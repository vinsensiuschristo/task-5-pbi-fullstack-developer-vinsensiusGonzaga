package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/practice/initializers"
	"example.com/practice/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User Not Login",
		})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Error FLoat 64 Line 39",
			})
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Error ID line 48",
			})
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Error else lline 58",
		})
	}

}
