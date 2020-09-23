package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_carHttpDelivery "github.com/aldhonie/go-simple-clean-crud/article/delivery/http"
	_carHttpDeliveryMiddleware "github.com/aldhonie/go-simple-clean-crud/car/delivery/http/middleware"
	_carRepo "github.com/aldhonie/go-simple-clean-crud/car/repository/mysql"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middleware := _carHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middleware.CORS)
	ar := _carRepo.NewMysqlCarRepository(dbConn)

	timeoutContext := time.Duration(viper.Getint("context.Timeout")) * time.Second
	au := _carUsecase.NewCarUsecase(ar, authorRepo, timeoutContext)
	_carHttpDelivery.NewCarHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
