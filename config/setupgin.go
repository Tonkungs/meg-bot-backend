package setupgin

import (
	"io"
	routes "meg-bot-backend/app"
	middlewares "meg-bot-backend/middlewares"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init ...
func Init() *gin.Engine {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	// ใส่ cors เพื่อให้เค้าได้จากทุก url ที่ร้องขอมา
	router.Use(cors.Default())
	// กัน hack helmet
	router.Use(helmet.Default())

	// Middleware
	router.Use(middlewares.ErrorHandler)
	// เพื่อม routes
	routerss := routes.SetupRoutes(router)

	return routerss

}
