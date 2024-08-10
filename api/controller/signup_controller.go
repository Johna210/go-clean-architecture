package controller

import "github.com/Johna210/go-clean-architecture/domain"

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env
}
