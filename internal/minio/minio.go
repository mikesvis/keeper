package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"keeper/internal/config"
)

func NewMinio(c *config.Config, logger *zap.SugaredLogger) (*minio.Client, error) {
	logger.Infof("crete new minio client, host: %s key: %s useSsl: %t", c.MinioEndpoint, c.MinioAccessKey, c.MinioUseSSL)

	client, err := minio.New(c.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.MinioAccessKey, c.MinioSecretKey, ""),
		Secure: c.MinioUseSSL,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
