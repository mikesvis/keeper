package secret

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"keeper/internal/domain"
	"keeper/internal/storage/files"
	"keeper/internal/storage/secret"
)

type Service struct {
	secretStorage *secret.Storage
	fileStorage   *files.Storage
	logger        *zap.SugaredLogger
}

func NewService(secretStorage *secret.Storage, fileStorage *files.Storage, logger *zap.SugaredLogger) *Service {
	return &Service{
		secretStorage: secretStorage,
		fileStorage:   fileStorage,
		logger:        logger,
	}
}

func (s *Service) Create(ctx context.Context, userID uuid.UUID, secretType domain.UserSecretType, name string, data *domain.UserSecretData) error {
	newSecret := domain.NewUserSecret(userID, secretType, name, data)

	createdSecret, err := s.secretStorage.CreateSecret(ctx, newSecret)
	if err != nil {
		return err
	}

	fileData, err := interface{}(*data).(domain.UserSecretData).GetData()
	if err != nil {
		return err
	}

	err = s.fileStorage.Store(ctx, createdSecret.ID, fileData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUserSecrets(ctx context.Context, userID uuid.UUID) ([]*domain.UserSecret, error) {
	secrets, err := s.secretStorage.GetAllForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, oneSecret := range secrets {
		d, err := s.fileStorage.Get(ctx, oneSecret.ID)
		if err != nil {
			return nil, err
		}

		data, err := domain.MakeUserSecretData(oneSecret.Type, d)
		if err != nil {
			return nil, err
		}
		oneSecret.Data = &data
	}

	return secrets, nil
}

func (s *Service) Delete(ctx context.Context, secretId uuid.UUID) error {
	err := s.secretStorage.Delete(ctx, secretId)
	if err != nil {
		return err
	}

	err = s.fileStorage.Delete(ctx, secretId)
	if err != nil {
		return err
	}

	return nil
}
