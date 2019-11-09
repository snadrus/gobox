package gobox

import "github.com/gin-gonic/gin"

//todo put gobox cmd somewhere.

type ObjFunc func(c *gin.Context) interface{}

//
func StartServer(func(apiRoutes gin.RouterGroup)) {
	// todo further cripple the API to HTML(template, objfunc), JSON(objfunc), group
}

// TODO component registration & use? or just leave it?
