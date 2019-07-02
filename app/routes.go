package routes

import (
	"fmt"
	"time"

	// "strconv"
	// "io/ioutil"
	"net/http"
	// "path/filepath"
	"log"

	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson"
	"encoding/json"
	"meg-bot-backend/app/models"
	utils "meg-bot-backend/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SetupRoutes ...
func SetupRoutes(router *gin.Engine) *gin.Engine {

	router.SecureJsonPrefix(")]}',\n")
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// Statics
	// router.Static("/public", "../public")
	api := router.Group("/api")
	{
		// create ObjectID from string
		id, err := primitive.ObjectIDFromHex("5abf4f07396b6036a0391383")
		if err != nil {
			log.Fatal(err)
		}

		ruan := models.User{id, "Ruand", "34", "Cape Town", time.Now(), time.Now()}
		// fmt.Println("User =", ruan.GetDisplay())
		// fmt.Println("User =",ruan.NamePhoto())
		api.GET("/ping", func(c *gin.Context) {
			fmt.Println("=================================")
			fmt.Println("c.Request.URL", c.Request.URL)
			fmt.Println("c.Params", c.Params)
			// pageNoAndsize := models.PageNoAndsize{}
			page := models.GetDefaultPageNoAndsize(c)
			// StructToMap
			var inInterface map[string]interface{}
			inrec, _ := json.Marshal(ruan)
			json.Unmarshal(inrec, &inInterface)
			// ddd := []interface{}{
			// 	inInterface,
			// }
			page.CountList = 100

			utils.Respond(c, http.StatusOK, utils.MessageOne(true, "pass", inInterface))
			// utils.Respond(c, http.StatusOK, utils.MessageMany(true, "pass", page, ddd))

			// c.JSON(http.StatusOK,  utils.MessageMany(true,"pass",page,ddd))
		})
	}
	return router
}
