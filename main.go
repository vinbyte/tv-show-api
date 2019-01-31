package main

import (
	"fmt"

	"github.com/501army/go-tv-show/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// var conf = config.ReadConfig()
// var key = conf.ServerKey

// func middleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		bearer := c.Request.Header.Get("Authorization")
// 		if bearer != "" {
// 			strSplit := strings.Split(bearer, " ")
// 			if strSplit[0] == "Bearer" && strSplit[1] != "" {
// 				if strSplit[1] == key {
// 					c.Next()
// 				} else {
// 					c.AbortWithStatus(401)
// 					c.JSON(401, gin.H{
// 						"status":  401,
// 						"message": "unauthorized",
// 					})
// 				}
// 			} else {
// 				c.AbortWithStatus(401)
// 				c.JSON(401, gin.H{
// 					"status":  401,
// 					"message": "unauthorized",
// 				})
// 			}
// 		} else {
// 			c.AbortWithStatus(401)
// 			c.JSON(401, gin.H{
// 				"status":  401,
// 				"message": "unauthorized",
// 			})
// 		}
// 	}
// }

func main() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	doRequest := new(controllers.RequestController)
	// db.Init()
	// peopleController.Create()
	if viper.GetString("mode") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	// router.Use(middleware())
	config := cors.DefaultConfig()
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my API",
		})
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to my API",
			})
		})
		v1.GET("/schedule", doRequest.Schedule)
		v1.GET("/search", doRequest.Search)
	}

	router.Run(":" + viper.GetString("port"))
}
