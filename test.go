package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func signup(c echo.Context) error {
	indextmp := "C:/Users/user/go/src/Gsmfestival-Master/login.html"
	myemail := "myemail"
	pwd := "pwd"
	fmt.Println(c.FormValue("classnum"))
	fmt.Println(c.FormValue("myname"))
	fmt.Println(c.FormValue(myemail))
	fmt.Println(c.FormValue(pwd))
	fmt.Println(c.FormValue("pwdck"))
	fmt.Println("----------")
	return c.File(indextmp)
}

func login(c echo.Context) error {
	semail := "myemail"
	spwd := "pwd"
	email := "email"
	pwd := "password"
	fmt.Println(c.FormValue("email"))
	fmt.Println(c.FormValue("password"))
	if semail == email && spwd == pwd {
		fmt.Println("로그인 성공")
	}
	return c.File("C:/Users/user/go/src/Gsmfestival-Master/index.html")
}

// func main() {
// 	tmp := "C:/Users/user/go/src/Gsmfestival-Master"
// 	e := echo.New()
// 	e.Use(middleware.Static(tmp))
// 	e.POST("/signup", signup)
// 	e.POST("/login", login)
// 	e.Logger.Fatal(e.Start(":1325"))

// }
