package user

import (
	"context"

	"github.com/tmolyakov/go-api-xmp/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO
	return
}
