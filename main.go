package main

import (
	"fmt"
	"log"
	"os"

	"developer.zopsmart.com/go/backend/zs"
	"github.com/zopping/mock-test/services"
	"github.com/zopping/mock-test/stores"
)

func main() {
	fmt.Println("mocking test")
	var infoLogger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	db := zs.NewMYSQL(infoLogger, zs.MySQLConfig{Port: "3306", HostName: os.Getenv("DB_HOST"), User: os.Getenv("DB_USER"), Password: os.Getenv("DB_PASSWORD"), Database: "organization_service"})

	userStore := stores.New(db)
	userSvc := services.New(userStore)
	out, err := userSvc.Find(1)
	fmt.Println(out, err)
}
