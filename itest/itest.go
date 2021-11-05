package itest

import (
	"dance/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleSetContactInfo(c *core.Context, client *http.Client) {
	c.JSON(http.StatusOK, gin.H{"aa": "bb"})
}

func init() {
	core.Engine.GET("test", func(context *gin.Context) {
		//name := context.Param("name")
		//action := context.Param("action")

		//fmt.Println(action)
		////  截取/
		////action = strings.Trim(action, "/")
		//context.String(http.StatusOK, name+" is "+action)

		context.JSON(http.StatusOK, gin.H{"aa": "bb"})
	})
}
