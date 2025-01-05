package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"keeper/internal/domain"
	"keeper/internal/storage/user"
	"keeper/pkg/hash"
	"keeper/pkg/validators"
)

var ErrUnauthorizedUser = errors.New(`can not authorize user`)
var ErrIternal = errors.New(`failed getting user from storage`)
var ErrConflictCreatingUser = errors.New(`user with provided creds already exists`)

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

func (s *Service) Login(ctx context.Context, login string, password string) (*domain.AuthenticatedUser, error) {
	err := validators.LoginValidator(login)
	if err != nil {
		return nil, err
	}

	err = validators.PasswordValidator(password)
	if err != nil {
		return nil, err
	}

	s.logger.Infof("searching user by creds (no password here he-he) login %s", login)
	userID, err := s.storage.GetUserID(ctx, login, hash.Hash([]byte(password)))

	if err != nil {
		s.logger.Errorf("error while searching user by creds with login %s, %v", login, err)
		return nil, ErrIternal
	}

	if userID == uuid.Nil {
		s.logger.Errorf("error because you are a cheater, %s!", login)
		return nil, ErrUnauthorizedUser
	}

	s.logger.Infof("found user by creds login %s id %s", login, userID)
	return &domain.AuthenticatedUser{ID: userID}, nil
}

func (s *Service) Register(ctx context.Context, login string, password string) (*domain.AuthenticatedUser, error) {
	err := validators.LoginValidator(login)
	if err != nil {
		return nil, err
	}

	err = validators.PasswordValidator(password)
	if err != nil {
		return nil, err
	}

	s.logger.Infof("searching existing user by login %s", login)
	exist, err := s.storage.ExistByLogin(ctx, login)
	if err != nil {
		s.logger.Errorf("error while searching existing user by login %s, %v", login, err)
		return nil, ErrIternal
	}

	if exist {
		s.logger.Infof("user with login %s already exists", login)
		return nil, ErrConflictCreatingUser
	}

	s.logger.Infof("creating user with login %s", login)
	userID, err := s.storage.Create(ctx, login, hash.Hash([]byte(password)))
	if err != nil {
		s.logger.Errorf("error while creating new user by login %s, %v", login, err)
		return nil, err
	}

	s.logger.Infof("created user with login %s id %s", login, userID)
	return &domain.AuthenticatedUser{ID: userID}, nil
}
