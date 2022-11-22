package app

import (
	"shake-shake/src/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	router   *gin.Engine
	member   *MemberApp
	vacation *VacationApp
}

func CreateApp() (*App, error) {
	router := gin.Default()
	docs.SwaggerInfo.Title = "shake-shake Server API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Description = "Open API Docs for shake-shake Server."
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	member, err := CreateMemberApp(router)
	vacation, err := CreateVacationApp(router)

	if err != nil {
		return nil, err
	}

	app := &App{
		router:   router,
		member:   member,
		vacation: vacation,
	}
	return app, nil
}

func (a *App) Run() {
	a.router.Run()
}
