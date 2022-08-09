package main

import (
	"fmt"

	"github.com/alifarhan1230/pjk_abt/config"
	"github.com/alifarhan1230/pjk_abt/module"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	db, err := config.InitDatabase()
	if err != nil {
		fmt.Print(err)
	}
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/owners", module.FindOwner)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
