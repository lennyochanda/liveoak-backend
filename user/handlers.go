package user

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lennyochanda/LiveOak/middleware"
	"github.com/lennyochanda/LiveOak/tokenutil"
)

type CreateUserDTO struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SetUpUserRoutes(us *UserService) {
	router := gin.Default()
	var (
		AccessTokenExpiryHour, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOURS"))
		AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
		// RefreshTokenExpiryHour, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_REFRESH_HOURS"))
		// RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
		FrontendLink = os.Getenv("FRONTEND_LINK")
	)

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
			c.JSON(http.StatusCreated, gin.H{
				"message": "New User Created",
			})
		})
		userHandlersv1.POST("/login", middleware.JWTMiddleware(AccessTokenSecret), func(c *gin.Context) {
			var loginUserForm LoginUserDTO
			if err := c.BindJSON(&loginUserForm); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			found, err := us.GetUserByEmail(loginUserForm.Email)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			
			valid, err := us.CheckPassword(loginUserForm)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			// check if id extracted from token matches id from db
			id, exists := c.Get("liveoak-id")
			if exists {
				fmt.Print("id", id)
				if found.ID == id {
					fmt.Print("same")
					c.Redirect(http.StatusFound, FrontendLink)
					return
				}
			}

			if valid {
				tokenString, err := tokenutil.CreateAccessToken(found, AccessTokenSecret, AccessTokenExpiryHour)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					c.Abort()
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"token": tokenString,
				})
				return
			}
		})
	}
	router.Run()
}
