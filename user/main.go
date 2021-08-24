package main

import (
	"net/http"

	gormModel "github.com/SeijiOmi/user/gorm/model"
	"github.com/SeijiOmi/user/graph"
	"github.com/SeijiOmi/user/graph/generated"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	e := echo.New()

	dsn := "root:@tcp(127.0.0.1:3306)/dataloader?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:               dsn,
			DefaultStringSize: 256, // default size for string fields
		},
	), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&gormModel.User{})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	srv :=
		graph.Middleware(db, handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}})))

	e.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":4001"))
}
