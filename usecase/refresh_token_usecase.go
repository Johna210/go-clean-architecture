package usecase

import (
	"context"
	"time"

	"github.com/Johna210/go-clean-architecture/domain"
	tokenutil "github.com/Johna210/go-clean-architecture/internal/tokenUtil"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, id)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
