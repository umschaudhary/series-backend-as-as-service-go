package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "blog/infrastructure"
)

func main() {
  router := gin.Default() //new gin router initialization
  infrastructure.NewDatabase()
  router.GET("/", func(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})    
  }) // first endpoint returns Hello World
  router.Run(":8000") //running application, Default port is 8080
}
