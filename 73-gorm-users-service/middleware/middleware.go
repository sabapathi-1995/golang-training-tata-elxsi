package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"user-service/security"

	"github.com/gin-gonic/gin"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		var tokenString string

		fmt.Println(authHeader)
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			}
		}

		if tokenString == "" {
			ctx.String(http.StatusUnauthorized, "Missing JWT token")
			ctx.Abort()
			return
		}

		claims, err := security.ValidateJWT(tokenString, secret)
		if err != nil {
			slog.Error(err.Error())
			ctx.String(http.StatusUnauthorized, "Invalid or expired token")
			ctx.Abort() // This is very imp for middleware
			return
		}

		ctx.Set("user_name", claims.Username)
		ctx.Set("email", claims.Email)
		ctx.Next() // You let the next handlerfunc to proceed
	}
}
