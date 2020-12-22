package rest

import (
	"GSM-Festival-Master-Back/database"
	jwt "GSM-Festival-Master-Back/lib"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/labstack/echo"
)

// type handlerInterface interface {
// 	Login(c echo.Context) error
// 	Signup(c echo.Context) error
// 	updateUser(c echo.Context) error
// 	SignOut(c echo.Context)
// 	deleteUser(c echo.Context) error
// }

//Get main Get
func GetMainPage(c echo.Context) (err error) {
	// return c.String(200, "main page")
	return c.File("/Users/baetaehyeon/go/src/Gsmfestival-Master-Front/index.html")
}

//S
type SignUpParam struct {
	Classnum  string `json:"classnum" form:"classnum" query:"classnum"`
	Name      string `json:"name" form:"name" query:"name"`
	Email     string `json:"email" form:"email" query:"email"`
	Pw        string `json:"pw" form:"pw" query:"pw"`
	IsManager bool   `json:"is_manager" form:"is_mamager" query:"is_manager"`
}

//Get signup Get
func Signup(c echo.Context) (err error) {
	u := new(SignUpParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println(u.Classnum, u.Name, u.Email, u.Pw)
	if u.Classnum == "" || u.Name == "" || u.Email == "" || u.Pw == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "모든 값을 입력해주세요",
		})
	}
	User := &database.User{}

	err = database.DB.Where("classnum = ?", u.Classnum).Find(User).Error
	if err == nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "이미 사용중인 학번입니다.",
		})
	}

	// if e.POST == "/admin" {
	// 	u.IsManager == "true"
	// }
	// u.IsManager = false
	User = &database.User{Classnum: u.Classnum, Name: u.Name, Email: u.Email, Pw: u.Pw, IsManager: u.IsManager}
	err = database.DB.Create(User).Error
	//ismanager not null error

	//pq: column "is_manager" does not exist
	//pq: Could not complete operation in a failed transaction

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "회원가입 실패",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "회원가입이 완료되었습니다",
	})
}

//S
type SignInParam struct {
	Email string `json:"email" form:"email" query:"email"`
	Pw    string `json:"pw" form:"pw" query:"pw"`
}

//Get signin page
func Signin(c echo.Context) (err error) {
	u := new(SignInParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println(u.Email, u.Pw)
	User := &database.User{}
	err = database.DB.Where("email = ? AND pw = ?", u.Email, u.Pw).Find(User).Error
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":       400,
			"message":      "일치하는 회원이 없습니다",
			"refreshToken": "null",
			"accessToken":  "null",
		})
	}
	refreshToken, err := jwt.CreateRefreshToken(User.Email, User.Pw)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       500,
			"message":      "refreshToken 생성 중 에러",
			"refreshToken": "null",
			"accessToken":  "null",
		})
	}
	accessToken, err := jwt.CreateAccessToken(User.Email, User.Pw, User.IsManager)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       500,
			"message":      "accessToken 생성 중 에러",
			"refreshToken": refreshToken,
			"accessToken":  "null",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":       200,
		"message":      "토큰 발급 완료",
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}

//Classroom in use page
func Classroominuse(c echo.Context) (err error) {
	if err != nil {
		return err
	}
	return c.File("/Users/baetaehyeon/go/src/Gsmfestival-Master-Front/admin.html")
}

//Loginpage
func Loginpage(c echo.Context) (err error) {
	if err != nil {
		return err
	}
	return c.File("/Users/baetaehyeon/go/src/Gsmfestival-Master-Front/login.html")
}

//Signuppage
func Signuppage(c echo.Context) (err error) {
	if err != nil {
		return err
	}
	return c.File("/Users/baetaehyeon/go/src/Gsmfestival-Master-Front/signup.html")
}

//Applicationpage
func Applicationpage(c echo.Context) (err error) {
	if err != nil {
		return err
	}
	return c.File("/Users/baetaehyeon/go/src/Gsmfestival-Master-Front/submit.html")
}
