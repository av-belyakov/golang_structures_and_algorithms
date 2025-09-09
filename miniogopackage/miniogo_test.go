package miniogopackage_test

import (
	"log"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/miniogopackage/miniosettings"
)

func TestMinIoGo(t *testing.T) {
	settings, err := miniosettings.NewSettings(".env")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = minio.New(settings.GetEndpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(settings.GetAccessKeyID(), settings.GetSecretAccessKey(), ""),
		Secure: false,
	})
	assert.NoError(t, err)
}
