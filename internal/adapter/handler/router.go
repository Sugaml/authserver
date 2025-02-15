package http

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	sloggin "github.com/samber/slog-gin"
	"github.com/sugaml/authserver/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter creates a new HTTP router
func (h *Handler) NewRouter() error {
	// Disable debug mode in production
	if h.config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	h.router.Use(CORSMiddleware())
	h.router.Use(sloggin.New(slog.Default()), gin.Recovery())

	// Custom validators
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		if err := v.RegisterValidation("user_role", userRoleValidator); err != nil {
			return err
		}
	}
	// Swagger
	h.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := h.router.Group("api/v1/auth")
	// Set Swagger
	setupSwagger(v1)
	v1.POST("/connect/token", func(c *gin.Context) {
		err := h.srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	})

	h.User(v1)
	h.Customer(v1)
	h.Application(v1)
	h.Client(v1)
	h.Secret(v1)

	return nil
}

func (h *Handler) User(v1 *gin.RouterGroup) {
	user := v1.Group("/users")
	{
		user.POST("/register", h.Register)

		authUser := user.Group("/").Use(authMiddleware(h.token))
		{
			authUser.GET("/", h.ListUsers)
			authUser.GET("/:id", h.GetUser)

			admin := authUser.Use(adminMiddleware())
			{
				admin.PUT("/:id", h.UpdateUser)
				admin.DELETE("/:id", h.DeleteUser)
			}
		}

	}
}

// Customer Endpoint
func (h *Handler) Customer(v1 *gin.RouterGroup) {
	client := v1.Group("/customer")
	{
		client.POST("", h.CreateCustomer)
		client.GET("/:id", h.GetCustomer)
		client.PUT("/:id", h.UpdateCustomer)
		client.DELETE("/:id", h.DeleteCustomer)
	}
}

// Application Endpoint
func (h *Handler) Application(v1 *gin.RouterGroup) {
	client := v1.Group("/application")
	{
		client.POST("", h.CreateApplication)
		client.GET("/:id", h.GetApplication)
		client.PUT("/:id", h.UpdateApplication)
		client.DELETE("/:id", h.DeleteApplication)
	}
}

// Client Endpoint
func (h *Handler) Client(v1 *gin.RouterGroup) {
	client := v1.Group("/client")
	{
		client.POST("", h.CreateClient)
		client.GET("/:id", h.GetClient)
		client.PUT("/:id", h.UpdateClient)
		client.DELETE("/:id", h.DeleteClient)
	}
}

// Secret Endpoint
func (h *Handler) Secret(v1 *gin.RouterGroup) {
	secret := v1.Group("/client-secret")
	{
		secret.POST("", h.CreateClientSecret)
		secret.GET("/:id", h.GetClientSecret)
		secret.PUT("/:id", h.UpdateClientSecret)
		secret.DELETE("/:id", h.DeleteClientSecret)
	}
}

// Serve starts the HTTP server
func (h *Handler) Serve(listenAddr string) error {
	err := h.NewRouter()
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}
	return h.router.Run(listenAddr)
}

// Swagger host path and basepath configuration
func setupSwagger(v1 *gin.RouterGroup) {
	hostPath := os.Getenv("HOST_PATH")
	if hostPath == "" {
		hostPath = "localhost:8080"
	}
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "/api/v1/parking"
	}
	docs.SwaggerInfo.Host = hostPath
	docs.SwaggerInfo.BasePath = basePath
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
