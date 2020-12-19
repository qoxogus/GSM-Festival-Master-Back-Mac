package rest

import (
	"github.com/labstack/echo"
	"github.com/qoxogus/GSM-Festival-Master-Back/database"
	jwt "github.com/qoxogus/GSM-Festival-Master-Back/lib"
)

type handlerInterface interface {
	Login(c echo.Context) error
	Signup(c echo.Context) error
	updateUser(c echo.Context) error
	SignOut(c echo.Context)
	deleteUser(c echo.Context) error
}

//Get main Get
func GetMainPage(c echo.Context) (err error) {
	// return c.String(200, "main page")
	return c.File("C:/Users/user/go/src/Gsmfestival-Master-Front/index.html")
}

//S
type SignUpParam struct {
	Classnum  string `json:"classnum" form:"classnum" query:"classnum"`
	Name      string `json:"name" form:"name" query:"name"`
	Email     string `json:"email" form:"email" query:"email"`
	Pw        string `json:"pw" form:"pw" query:"pw"`
	IsManager bool   `json:"manager" form:"mamager" query:"manager"`
}

//Get signup Get
func Signup(c echo.Context) (err error) {
	u := new(SignUpParam)
	if err := c.Bind(u); err != nil {
		return err
	}
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
	User = &database.User{Classnum: u.Classnum, Name: u.Name, Pw: u.Pw, IsManager: u.IsManager}
	err = database.DB.Create(User).Error
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
	// Bind
	// mu := new(model.User) //user 불러오기
	// if err = c.Bind(mu); err != nil {
	// 	return
	// }

	// // // email := bson.M{"email": u.Email}
	// // // password := bson.M{"password": u.Password}
	// filter := bson.M{"token": mu.Token}

	// collection, err := dblayer.GetDBCollection()
	// if err != nil {
	// 	return err
	// 	// return &echo.HTTPError{Code: http.StatusUnauthorized,Message:"invalid email or password"}
	// }

	// err = collection.FindOne(context.TODO(), filter).Decode(&u)

	// _, err = collection.UpdateOne(context.TODO(), filter, &u)

	// defer collection.Database().Client().Disconnect(context.TODO())

	// // Create token
	// token := jwt.New(jwt.SigningMethodHS256)

	// // Set claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["id"] = u.ID
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// // Generate encoded token and send it as response
	// u.Token, err = token.SignedString([]byte("secret"))
	// if err != nil {
	// 	return err
	// }
}

//Loginpage
func Loginpage(c echo.Context) (err error) {
	if err != nil {
		return err
	}
	return c.File("C:/Users/user/go/src/Gsmfestival-Master-Front/login.html")
}

//Signuppage
func Signuppage(c echo.Context) (err error) {
	if err != nil {
		return err
	}
	return c.File("C:/Users/user/go/src/Gsmfestival-Master-Front/signup.html")
}

//Applicationpage
func Applicationpage(c echo.Context) (err error) {
	if err != nil {
		return err
	}
	return c.File("C:/Users/user/go/src/Gsmfestival-Master-Front/submit.html")
}
