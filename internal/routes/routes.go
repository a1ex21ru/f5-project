package routes

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "f5-project/docs"

	"f5-project/internal/handlers"
)

type Handler struct {
	*handlers.NoteHandler
}

func NewHandler(handler *handlers.NoteHandler) *Handler {
	return &Handler{handler}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(static.Serve("/static", static.LocalFile("./static", true)))
	router.GET("/", h.GetIndexPage)
	router.GET("/create.html", h.GetCreatePage)
	router.GET("/edit.html", h.GetEditPage)

	note := router.Group("/api/notes")
	{
		note.POST("", h.CreateNote)
		note.DELETE("/:id", h.DeleteNote)
		note.PATCH("/:id", h.UpdateNote)
		note.GET("/:id", h.GetNoteByID)
		note.GET("", h.GetAll)
	}

	router.GET("/healthcheck", h.HealthCheck)

	return router
}
