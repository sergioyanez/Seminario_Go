package main

import (
	"tudai2021/internal/services/user"

	"github.com/gin-gonic/gin"
)

func main() { // este main lo hago correr co el comando  -- go run cmd/user/user.go  --
	service, err := user.NewUserService() // con _ ignoro la variable o el error, depende donde lo coloque
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		u := c.Query("user")
		p := c.Query("pass")
		user, _ := service.Login(u, "")
		c.JSON(200, gin.H{
			"message": u + " " + p + " " + user.Name,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
