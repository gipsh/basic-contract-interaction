package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gipsh/basic-contract-interaction/controllers"
	"github.com/joho/godotenv"
)

type App struct {
	Router *gin.Engine
}

func New() App {

	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := App{}
	app.Router = gin.New()

	return app

}

func (app *App) Run() {

	pc := controllers.NewProductController()

	app.Router.GET("/product/:id", pc.ProductById)
	app.Router.GET("/products", pc.Products)
	app.Router.POST("/product", pc.CreateProduct)
	app.Router.POST("/product/delegate", pc.DelegateProduct)
	app.Router.POST("/product/accept", pc.AcceptProduct)

	app.Router.Run()

}
