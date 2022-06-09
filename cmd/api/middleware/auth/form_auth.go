package auth

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

func FormMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		formMiddlewareImpl(c)
	}
}

func formMiddlewareImpl(c *gin.Context) {
	claims, err := getClaimsFromJWT(c)
	if err != nil {
		unauthorized(c, http.StatusUnauthorized, err.Error())
		return
	}

	if claims["exp"] == nil {
		unauthorized(c, http.StatusBadRequest, errors.New("missing exp field").Error())
		return
	}

	if _, ok := claims["exp"].(float64); !ok {
		unauthorized(c, http.StatusBadRequest, errors.New("exp must be float64 format").Error())
		return
	}

	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		unauthorized(c, http.StatusUnauthorized, errors.New("token is expired").Error())
		return
	}

	c.Set("JWT_PAYLOAD", claims)
	//identity := Config.IdentityHandler(c)
	//
	//if identity != nil {
	//	c.Set(Config.IdentityKey, identity)
	//}
	//
	//if !Config.Authorizator(identity, c) {
	//	unauthorized(c, http.StatusForbidden, errors.New("you don't have permission to access this resource").Error())
	//	return
	//}
	c.Next()
}

func getClaimsFromJWT(c *gin.Context) (jwt.MapClaims, error) {
	token, err := parseFormToken(c)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	for key, value := range token.Claims.(jwt4.MapClaims) {
		claims[key] = value
	}
	return claims, nil
}

func parseFormToken(c *gin.Context) (*jwt4.Token, error) {
	var token string
	var err error

	methods := strings.Split(Config.TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "query":
			token, err = jwtFromQuery(c, v)
		case "form":
			token, err = jwtFromForm(c, v)
		}
	}
	if err != nil {
		return nil, err
	}

	if Config.KeyFunc != nil {
		return jwt4.Parse(token, Config.KeyFunc)
	}

	return jwt4.Parse(token, func(t *jwt4.Token) (interface{}, error) {
		if jwt4.GetSigningMethod(Config.SigningAlgorithm) != t.Method {
			return nil, jwt.ErrInvalidSigningAlgorithm
		}
		// save token string if valid
		c.Set("JWT_TOKEN", token)
		return Config.Key, nil
	})
}

func jwtFromForm(c *gin.Context, key string) (string, error) {
	token := c.PostForm(key)

	if token == "" {
		return "", jwt.ErrEmptyParamToken
	}

	return token, nil
}

func jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", jwt.ErrEmptyQueryToken
	}

	return token, nil
}

func unauthorized(c *gin.Context, code int, message string) {
	c.Abort()
	Config.Unauthorized(c, code, message)
}
