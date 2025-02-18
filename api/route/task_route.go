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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
	group.GET("/task/:id", tc.GetByID)
}
