package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	InitRouter(r)

	r.Run(":8777") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
