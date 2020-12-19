package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/qoxogus/GSM-Festival-Master-Back/config"
)

type connectionMethod interface {
	Connect()
}

// DB - 데이터베이스 전역변수
var DB *gorm.DB

// Connect - 데이터베이스 구조 생성, 연결 하는 메서드
func Connect() {
	dbConf := config.Config.DB

	connectOptions := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.Username,
		dbConf.Name,
		dbConf.Password)

	db, err := gorm.Open("postgres", connectOptions)

	if err != nil {
		panic(err)
	}

	DB = db

	log.Print("[DATABASE] 연결 완료")
}

// //DB
// func GetDBCollection() (*mongo.Collection, error) {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// 	// Connect to MongoDB (연결)
// 	fmt.Println("[MongoDB Connect]")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check the connection(연결검증)
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// GPM.users DataBase     MongoDB
// 	collection := client.Database("GPM").Collection("users")

// 	return collection, nil
// }
