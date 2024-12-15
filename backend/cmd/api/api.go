package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yikuanzz/social/internal/base/app"
	"github.com/yikuanzz/social/internal/base/conf"
	"github.com/yikuanzz/social/internal/base/data"
	"github.com/yikuanzz/social/internal/base/log"
	"github.com/yikuanzz/social/internal/base/server"
	"github.com/yikuanzz/social/internal/controller"
	"github.com/yikuanzz/social/internal/router"
	"github.com/yikuanzz/social/internal/service"
	"github.com/yikuanzz/social/internal/store"
)

var (
	Name      = "social-api"
	Version   = "0.0.0"
	GoVersion = "1.22"
)

func Main() {
	app, cleanup, err := initApplication()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}

func initApplication() (*app.Application, func(), error) {
	// Connect to mysql server
	dbEngine, err := data.NewDB()
	if err != nil {
		log.Errorf("failed to connect to db: %v", err)
		return nil, nil, err
	}

	// Create data layer
	dataData, cleanup, err := data.NewData(dbEngine)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	// Create store layer
	userStore := store.NewUserStore(dataData)

	// Create service layer
	userService := service.NewUserService(userStore)

	// Create controller layer
	userController := controller.NewUserController(userService)

	// Create router layer
	apiRouter := router.NewSocialApiRouter(userController)

	// Create http server
	ginEngine := newHTTPServer(apiRouter)

	// Create application
	application := newApplication(ginEngine)
	return application, func() { cleanup() }, nil
}

func newApplication(e *gin.Engine) *app.Application {
	return app.NewApp(
		app.WithName(Name),
		app.WithVersion(Version),
		app.WithServer(server.NewServer(e, conf.ServerConfigs.Addr)),
	)
}

func newHTTPServer(
	socialApiRouter *router.SocialApiRouter,
) *gin.Engine {
	// New gin engine and middleware
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })

	// Register api that no need to login
	unAuthV1 := r.Group("/social/api/v1")
	socialApiRouter.RegisterUnAuthorizedRoutes(unAuthV1)

	// Register api that need to login
	authV1 := r.Group("/social/api/v1")
	socialApiRouter.RegisterAuthorizedApiRouter(authV1)
	return r
}
