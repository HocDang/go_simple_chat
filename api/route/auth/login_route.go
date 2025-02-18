package route

import (
	"time"

	authController "w3s/go-backend/api/controller/auth"
	"w3s/go-backend/bootstrap"
	"w3s/go-backend/domain"
	"w3s/go-backend/mongo"
	"w3s/go-backend/repository"
	"w3s/go-backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &authController.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
