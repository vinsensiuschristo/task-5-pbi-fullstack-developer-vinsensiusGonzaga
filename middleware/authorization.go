package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// const (
// 	authorizationHeaderKey  = "authorization"
// 	authorizationTypeBearer = "bearer"
// 	authorizationPayloadKey = "authorization_payload" //Authorization
// )

func AuthorizationMiddleware() gin.HandlerFunc {
	// return func(ctx *gin.Context) {
	// 	// cek auth
	// 	authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
	// 	if len(authorizationHeader) == 0 {
	// 		err := errors.New("authorization header is  not provided")
	// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return

	// 	}

	// 	// cek
	// 	fields := strings.Fields(authorizationHeader)
	// 	if len(fields) < 2 {
	// 		err := errors.New("invalid authorization header format")
	// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	// cek apakah auth = bearer
	// 	authorizationType := strings.ToLower(fields[0])
	// 	if authorizationType != authorizationHeader {
	// 		err := errors.New("unsupoported authorization type")
	// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	// cek token
	// 	accessToken := fields[1]
	// 	payload, err := tokenMaker.VerifyToken(accessToken)
	// 	if err != nil {
	// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	ctx.Set(authorizationPayloadKey, payload)
	// 	ctx.Next()

	// }

	return func(c *gin.Context) {
		// TODO: Implement JWT validation logic

		// JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return os.Getenv("SECRET"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}
		c.Next()
	}
}
