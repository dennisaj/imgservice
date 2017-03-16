# main.go
package main

import (
  "fmt"
  redis "gopkg.in/redis.v4"
  import "gopkg.in/gin-gonic/gin.v1"
)

func main() {
   client := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_URL"),
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    r := gin.Default()
        r.GET("/image", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })
    r.Run() // listen and serve on 0.0.0.0:8080
    // pong, err := client.Ping().Result()
    // fmt.Println(pong, err)
}
