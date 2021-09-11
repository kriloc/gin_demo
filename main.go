// gin demo from b站
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Recipe struct{
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"published_at"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.Run()

}
