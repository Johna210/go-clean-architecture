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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
