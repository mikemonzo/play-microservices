package api

import (
	"github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/config"
	"github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/docs"
	interceptors "github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/internal/api/interceptors"
	jh "github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/internal/models/job/handler"
	rh "github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/internal/models/report/handler"
	uh "github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/internal/models/user/handler"
	"github.com/mikemonzo/play-microservices/api-gateway/api-gateway-service/pkg/logger"
	"github.com/gin-gonic/gin"

	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	log logger.Logger
	cfg *config.Config
}

func NewRouter(log logger.Logger, cfg *config.Config) *Router {
	return &Router{log: log, cfg: cfg}
}
func (r *Router) Setup(router *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/api/v1/ping", r.Pong)

	userHandler := uh.NewUserHandler(r.log, r.cfg)
	jobHandler := jh.NewJobHandler(r.log, r.cfg)
	reportHandler := rh.NewReportHandler(r.log, r.cfg)

	router.POST("/api/v1/user/create", userHandler.CreateUser)
	router.POST("/api/v1/user/login", userHandler.LoginUser)
	router.POST("/api/v1/user/refresh_token", userHandler.RefreshAccessToken)
	router.POST("/api/v1/user/logout", userHandler.LogOutUser)

	// Apply the middleware to the routes inside the router.Group function
	api := router.Group("/api")
	api.Use(interceptors.AuthenticateUser()) // Apply the middleware here
	{
		api.GET("/v1/user/get", userHandler.GetUser)
		api.GET("/v1/user/list", userHandler.ListUsers)

		api.POST("/v1/job/create", jobHandler.CreateJob)
		api.POST("/v1/job/update", jobHandler.UpdateJob)
		api.GET("/v1/job/get", jobHandler.GetJob)
		api.GET("/v1/job/list", jobHandler.ListJobs)
		api.POST("/v1/job/delete", jobHandler.DeleteJob)

		api.GET("/v1/report/list", reportHandler.ListReports)
	}
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping
// @Schemes
// @Description do ping
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {string} Pong
// @Router /ping [get]
func (s *Router) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}