package files

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"keeper/internal/config"
)

type Storage struct {
	bucketName string
	client     *minio.Client
}

func NewStorage(c *config.Config, client *minio.Client) *Storage {
	return &Storage{
		bucketName: c.MinioBucketName,
		client:     client,
	}
}

func (s *Storage) Store(ctx context.Context, objectID uuid.UUID, data []byte) error {
	reader := bytes.NewReader(data)

	_, err := s.client.PutObject(
		ctx,
		s.bucketName,
		objectID.String(),
		reader,
		-1,
		minio.PutObjectOptions{},
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Get(ctx context.Context, objectID uuid.UUID) ([]byte, error) {
	object, err := s.client.GetObject(
		ctx,
		s.bucketName,
		objectID.String(),
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}
	defer object.Close()

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(object)
	if err != nil {
		if err.Error() != "The specified key does not exist" {
			return nil, nil
		}
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *Storage) Delete(ctx context.Context, objectID uuid.UUID) error {
	return s.client.RemoveObject(ctx, s.bucketName, objectID.String(), minio.RemoveObjectOptions{})
}
