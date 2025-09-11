package miniogopackage_test

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"slices"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/miniogopackage/miniosettings"
	"github.com/av-belyakov/golang_structures_and_algorithms/supportfunctions"
)

func TestMinIoGoTLS(t *testing.T) {
	const USE_SSL = true

	var (
		testBucketName string   = "my-test-bucket"
		testObjectDir  string   = "./testfiles"
		testObjects    []string = []string{
			"file_1.txt",
			"file_2.txt",
		}
		newTestObjects []string
		minioClient    *minio.Client

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

	t.Run("Тест 2. Проверяем наличие bucket", func(t *testing.T) {
		isExist, err := minioClient.BucketExists(t.Context(), "mystorage")
		assert.NoError(t, err)
		assert.True(t, isExist)
	})

	t.Run("Тест 3. Проверяем наличие тестового bucket и если его нет, создаём", func(t *testing.T) {
		isExist, err := minioClient.BucketExists(t.Context(), testBucketName)
		assert.NoError(t, err)

		if !isExist {
			err := minioClient.MakeBucket(t.Context(), testBucketName, minio.MakeBucketOptions{Region: "rus-Moscow"})
			assert.NoError(t, err)
		}

		//список buckets
		listBuckets, err := minioClient.ListBuckets(t.Context())
		assert.NoError(t, err)
		assert.NotEmpty(t, listBuckets)

		fmt.Println("List buckets:")
		for k, v := range listBuckets {
			fmt.Printf("%d. %s\n", k, v)
		}
	})

	t.Run("Тест 4. Добавляем объекты в созданный bucket", func(t *testing.T) {
		for _, v := range testObjects {
			info, err := minioClient.FPutObject(t.Context(), testBucketName, v, path.Join(testObjectDir, v), minio.PutObjectOptions{})
			assert.NoError(t, err)

			fmt.Printf("Put object information:\n\tBucket:'%s'\n\tFileName:'%s'\n\tSize:%d\n", info.Bucket, info.Key, info.Size)
		}
	})

	t.Run("Тест 5. Получаем информацию по загруженным объектам", func(t *testing.T) {
		chInfo := minioClient.ListObjects(t.Context(), testBucketName, minio.ListObjectsOptions{
			Prefix:    "file",
			Recursive: true,
		})

		for objInfo := range chInfo {
			assert.NoError(t, objInfo.Err)

			fmt.Printf("Object information:\n\tName:%s\n\tETag:%s\n", objInfo.Key, objInfo.ETag)
		}
	})

	t.Run("Тест 6. Скачиваем объекты", func(t *testing.T) {
		for _, v := range testObjects {
			err := minioClient.FGetObject(t.Context(), testBucketName, v, fmt.Sprintf("./%s/upload_%s", testObjectDir, v), minio.GetObjectOptions{})
			assert.NoError(t, err)
		}

		items, err := os.ReadDir(testObjectDir)
		assert.NoError(t, err)

		for _, object := range testObjects {
			newTestObjects = append(newTestObjects, fmt.Sprintf("upload_%s", object))

			assert.True(t, slices.ContainsFunc(items, func(v os.DirEntry) bool {
				return v.Name() == fmt.Sprintf("upload_%s", object)
			}))
		}
	})

	t.Run("Тест 7. Копирование объектов", func(t *testing.T) {
		//в рамках одного bucket
		for _, v := range testObjects {
			uploadInfo, err := minioClient.CopyObject(
				t.Context(),
				minio.CopyDestOptions{
					Bucket:          testBucketName,
					Object:          fmt.Sprintf("copy_%s", v),
					ReplaceMetadata: true,
				},
				minio.CopySrcOptions{
					Bucket: testBucketName,
					Object: v,
				})
			assert.NoError(t, err)

			fmt.Printf("UploadInfo:%#v\n", uploadInfo)
		}

		bucketTmp := "my-test-tmp-storage"
		//с одного bucket в другой тестовый bucket
		err := minioClient.MakeBucket(t.Context(), bucketTmp, minio.MakeBucketOptions{
			Region: "nowhere",
		})
		assert.NoError(t, err)

		for _, v := range testObjects {
			_, err := minioClient.CopyObject(
				t.Context(),
				minio.CopyDestOptions{
					Bucket:          bucketTmp,
					Object:          fmt.Sprintf("file_%s", v),
					ReplaceMetadata: true,
				},
				minio.CopySrcOptions{
					Bucket: testBucketName,
					Object: v,
				})
			assert.NoError(t, err)
		}

		chInfo := minioClient.ListObjects(t.Context(), bucketTmp, minio.ListObjectsOptions{})
		result := []minio.ObjectInfo{}

		//fmt.Println("Object info:")
		for info := range chInfo {
			result = append(result, info)

			//fmt.Printf("\t%+v\n", result)
		}
		assert.NotEqual(t, len(result), 0)

		//удаляем временный bucket
		for _, obj := range result {
			assert.NoError(t, minioClient.RemoveObject(t.Context(), bucketTmp, obj.Key, minio.RemoveObjectOptions{}))
		}
		assert.NoError(t, minioClient.RemoveBucket(t.Context(), bucketTmp))
		isExist, err := minioClient.BucketExists(t.Context(), bucketTmp)
		assert.NoError(t, err)
		assert.False(t, isExist)
	})

	//t.Run("", func(t *testing.T) {
	//})

	t.Cleanup(func() {
		//удаление скачанных файловых объектов
		//************************************
		if err := supportfunctions.DeleteFilesInDirectory(testObjectDir, newTestObjects); err != nil {
			t.Logf("error delete files in directory: %s", err.Error())
		}

		//ctx := context.Background()

		//удаление всех объектов находящихся в bucket
		//*******************************************
		//chInfo := minioClient.ListObjects(ctx, testBucketName, minio.ListObjectsOptions{
		//	Recursive: true,
		//})
		//for obj := range chInfo {
		//	if err := minioClient.RemoveObject(ctx, testBucketName, obj.Key, minio.RemoveObjectOptions{}); err != nil {
		//		t.Logf("error remove objects: %s", obj.Err.Error())
		//	}
		//}
		//или так
		//chErr := minioClient.RemoveObjects(ctx, testBucketName, chInfo, minio.RemoveObjectsOptions{})
		//for obj := range chErr {
		//	if obj.Err != nil {
		//		t.Logf("error remove objects: %s", obj.Err.Error())
		//	}
		//}

		//удаление сомого bucket
		//**********************
		//if err := minioClient.RemoveBucket(ctx, testBucketName); err != nil {
		//	t.Logf("error bucket: %+v", err)
		//}
	})
}
