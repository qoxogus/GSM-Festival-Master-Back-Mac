package rest

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RunAPI used in main.go
func RunAPI(address string) {
	e := echo.New()
	// e.Use(middleware.Logger())
	tmp := "/Users/baetaehyeon/go/src/Gsmfestival-Master-Front"

	// e.Static("/", "static")

	e.Use(middleware.Static(tmp))

	e.GET("/main", GetMainPage) //서버 테스트용 코드
	e.GET("/loginpage", Loginpage)
	e.GET("/signuppage", Signuppage)
	e.GET("/application", Applicationpage)
	e.GET("/classroominuse", Classroominuse)

	e.POST("/signp", Signup)
	e.POST("/main", Signin)

	e.Logger.Fatal(e.Start(address))
}

//배태현
