package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserDTO struct {
	UserName string `json:"userName" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SetUpUserRoutes(us *UserService) {
	router := gin.Default()

	userHandlersv1 := router.Group("/v1/user")

	{
		userHandlersv1.POST("/create", func(c *gin.Context) {
			var createUserForm CreateUserDTO
			if err := c.BindJSON(&createUserForm); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			us.CreateUser(createUserForm)
			c.JSON(http.StatusOK, gin.H{
				"status": "New User Created",
			})
		})
	}
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	router.Run()
}