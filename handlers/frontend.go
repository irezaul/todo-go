package handlers

import "github.com/gin-gonic/gin"

func HomeFrontend(c *gin.Context) {
	c.HTML(200, "login.tmpl", nil)
}

func Register(c *gin.Context) {
	c.HTML(200, "register.tmpl", nil)
}
