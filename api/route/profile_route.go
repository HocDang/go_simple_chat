package route

import (
	"time"

	"w3s/go-backend/api/controller"
	"w3s/go-backend/bootstrap"
	"w3s/go-backend/domain"
	"w3s/go-backend/mongo"
	"w3s/go-backend/repository"
	"w3s/go-backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
