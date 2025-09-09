package miniogopackage_test

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/miniogopackage/miniosettings"
)

func TestMinIoGoTLS(t *testing.T) {
	const USE_SSL = true

	var (
		minioClient *minio.Client

		err error
	)

	settings, err := miniosettings.NewSettings(".env")
	if err != nil {
		log.Fatalln(err)
	}

	publicCert, err := os.ReadFile("minioimage/certs/public.crt")
	if err != nil {
		log.Fatalln(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(publicCert)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
	}

	t.Run("Тест 1. Проверяем создание подключения", func(t *testing.T) {
		minioClient, err = minio.New(settings.GetEndpoint(), &minio.Options{
			Creds:     credentials.NewStaticV4(settings.GetAccessKeyID(), settings.GetSecretAccessKey(), ""),
			Secure:    USE_SSL,
			Transport: tr,
		})
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Проверяем наличие бакета", func(t *testing.T) {
		isExist, err := minioClient.BucketExists(t.Context(), "mystorage")
		assert.NoError(t, err)
		assert.True(t, isExist)
	})
}
