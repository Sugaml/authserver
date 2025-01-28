package http

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	sloggin "github.com/samber/slog-gin"
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
	v1 := h.router.Group("api/v1")
	v1.POST("/token/connect", func(c *gin.Context) {
		err := h.srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	})
	h.User(v1)
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
	secret := v1.Group("/secret")
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
