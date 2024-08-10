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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
	}
	group.POST("/login", lc.Login)
}
