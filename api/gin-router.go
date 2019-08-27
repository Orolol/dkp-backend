package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/orolol/dkp-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var identityKey = "id"

func initRoutes() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin r with default middleware:
	// logger and recovery (crash-free) middleware

	r := gin.Default()

	r.Use(LiberalCORS)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("super secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(utils.AccountApi); ok {
				return jwt.MapClaims{
					identityKey: v.Login,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			var a *utils.AccountApi

			if _, ok := claims[identityKey].(string); ok {

				return claims[identityKey]
			}
			fmt.Println("OKCLAIMS")
			return a
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			var acc utils.Account
			var accApi utils.AccountApi
			db, _ := gorm.Open("mysql", ConnexionString)
			db.First(&acc, "Login = ?", userID)
			errPass := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))

			if errPass != nil {
				fmt.Println("Mauvais password", errPass, acc.Password, password)
				return nil, jwt.ErrFailedAuthentication
			} else if acc.ID == 0 {
				fmt.Println("Mauvais account")
				return nil, jwt.ErrFailedAuthentication
			}

			accApi.ID = acc.ID
			accApi.Login = acc.Login
			return accApi, nil
		},
		Authorizator: func(user interface{}, c *gin.Context) bool {
			// var acc utils.Account
			// db, _ := gorm.Open("mysql", ConnexionString)
			// db.First(&acc)

			// if v, ok := user.(string); ok && v == acc.Login {
			// 	return true
			// } else {
			// 	return false
			// }

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	auth := r.Group("/auth")

	r.POST("/Login", authMiddleware.LoginHandler)
	r.POST("/SignUp", SignUp)
	r.GET("/GetTranslations/:language", GetTranslations)
	r.GET("/GetInfos", GetInfos)
	r.GET("/GetPP", GetPP)
	r.GET("/GetServerInfos", GetServerInfos)
	r.GET("/GetDungeons", GetDungeons)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/RefreshToken", authMiddleware.RefreshHandler)
		auth.GET("/Index", Index)
		auth.POST("/GetProfileInfos", GetProfileInfos)
		auth.POST("/JoinGameAi", JoinGameAi)
		auth.POST("/StartDungeon", StartDungeon)
		auth.GET("/LeaveQueue", LeaveQueue)
		auth.POST("/EditAccount", EditAccount)
		auth.POST("/Actions", Actions)
		auth.GET("/GetEnemyInfos/:id", GetEnemyInfos)

		auth.POST("/GetHistory", GetHistory)
		auth.POST("/GetLeaderBoard", GetLeaderBoard)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(":8081")
	// r.Run(":3000") for a hard coded port

}

func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
