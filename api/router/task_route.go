package router

import (
	"time"

	"github.com/Johna210/go-clean-architecture/api/controller"
	"github.com/Johna210/go-clean-architecture/bootstrap"
	"github.com/Johna210/go-clean-architecture/domain"
	"github.com/Johna210/go-clean-architecture/repository"
	"github.com/Johna210/go-clean-architecture/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
