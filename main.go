package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_articleHttpDeliver "github.com/bxcodec/go-clean-arch/article/delivery/http"
	_articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	_articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	_authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
	_echoMiddleware "github.com/labstack/echo/middleware"

	_isHttpDeliver "github.com/auth/identityserver/delivery/http"
	_isUcase "github.com/auth/identityserver/usecase"
)

func main() {
	//accountStorage := "cgostorage"
	//accessKeyStorage := "OwvEOlzf6e7QwVoV0H75GuSZHpqHxwhYnYL9QbpVPgBRJn+26K26aRJxtZn7Ip5AhaiIkw9kH11xrZSscavXfQ=="
	//dbHost := viper.GetString(`database.host`)
	//dbPort := viper.GetString(`database.port`)
	//dbUser := viper.GetString(`database.user`)
	//dbPass := viper.GetString(`database.pass`)
	//dbName := viper.GetString(`database.name`)
	//dev db

	dbHost := "asparnas.database.windows.net"
	dbPort := 1433
	dbUser := "adminasparnas"
	dbPass := "Standar123"
	dbName := "asparnas"

	////dev IS
	//baseUrlis := "http://identity-server-asparnas.azurewebsites.net"
	////dev URL Forgot Password
	//urlForgotPassword := "http://cgo-web-api-dev.azurewebsites.net/account/change-password"
	//basicAuth := "cm9jbGllbnQ6c2VjcmV0"
	//redirectUrlGoogle := "http://cgo-web-api.azurewebsites.net/account/callback"
	//clientIDGoogle := "422978617473-acff67dn9cgbomorrbvhqh2i1b6icm84.apps.googleusercontent.com"
	//clientSecretGoogle := "z_XfHM4DtamjRmJdpu8q0gQf"

	connection := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		dbHost, dbUser, dbPass, dbPort, dbName)
	dbConn, err := sql.Open(`sqlserver`, connection)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	//middL := middleware.InitMiddleware()
	//e.Use(middL.CORS)
	e.Use(_echoMiddleware.CORSWithConfig(_echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)


	timeoutContext := 30 * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)


	_articleHttpDeliver.NewArticleHandler(e, au)
	log.Fatal(e.Start(":9090"))
}