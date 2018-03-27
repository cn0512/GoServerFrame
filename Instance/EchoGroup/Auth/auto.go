package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cn0512/GoServerFrame/Config"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
}

func Auth(c echo.Context) error {
	u := &User{}
	return c.JSON(http.StatusOK, u)
}
func main() {
	e := echo.New()
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(Config.JWT_rs256_public))
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	r := e.Group("/Auth")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:       func(echo.Context) bool { return false },
		SigningMethod: "RS256",
		ContextKey:    "user",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		SigningKey:    pubKey,
	}))
	r.GET("/Token", Auth)
	//save routers to file
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("routes.json", data, 0644)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
