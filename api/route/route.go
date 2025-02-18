package route

import (
	"time"

	"w3s/go-backend/api/middleware"
	"w3s/go-backend/bootstrap"
	"w3s/go-backend/mongo"

	authRoute "w3s/go-backend/api/route/auth"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {

	// Middleware server
	gin.Use(middleware.ServerMiddleware())

	apiRouter := gin.Group("api")

	// All Public APIs
	publicRouter := apiRouter.Group("")
	{
		// Home
		NewHomeRoute(env, timeout, db, publicRouter)

		// Auth
		authRouter := publicRouter.Group("auth")
		{
			authRoute.NewSignupRouter(env, timeout, db, authRouter)
			authRoute.NewLoginRouter(env, timeout, db, authRouter)
			authRoute.NewRefreshTokenRouter(env, timeout, db, authRouter)
		}
	}

	// All Protected APIs
	protectedRouter := apiRouter.Group("")
	{
		// Middleware to verify AccessToken
		protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

		// Profile
		NewProfileRouter(env, timeout, db, protectedRouter)

		// Task
		NewTaskRouter(env, timeout, db, protectedRouter)
	}

}
