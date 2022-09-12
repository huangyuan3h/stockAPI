package controller


import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8080")
}