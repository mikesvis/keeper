package user

import (
	"context"
	"errors"
	"github.com/charmbracelet/ssh"
	"go.uber.org/zap"
	"keeper/internal/storage/user"
)

var ErrUnauthorizedUser = errors.New(`can not authorize user`)
var ErrInternal = errors.New(`failed getting user from storage`)

type ContextKey string

const UserIDContextKey ContextKey = "UserID"

type Service struct {
	storage *user.Storage
	logger  *zap.SugaredLogger
}

func NewService(storage *user.Storage, logger *zap.SugaredLogger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) GetUserIDByKey(ctx context.Context, key string) (uint64, error) {
	s.logger.Info("searching incoming user by public key")
	userID, err := s.storage.GetUserIDByKey(ctx, key)

	if err != nil {
		s.logger.Errorf("error while searching user by public key: %v", err)
		return 0, ErrInternal
	}

	if userID == 0 {
		s.logger.Infof("user is not found by public key")
		return 0, ErrUnauthorizedUser
	}

	s.logger.Infof("found user id %d by public key", userID)
	return userID, nil
}

func (s *Service) IsUserAuthed(c ssh.Context) bool {
	if c.Value(UserIDContextKey) != nil {
		return true
	}

	return false
}
