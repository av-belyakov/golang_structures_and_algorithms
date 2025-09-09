package miniosettings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Settings map[string]string

func NewSettings(envFileName string) (*Settings, error) {
	if err := godotenv.Load(envFileName); err != nil {
		return &Settings{}, err
	}

	s := Settings{
		"apiPort":         os.Getenv("GO_MINIO_APIPORT"),
		"apiHost":         os.Getenv("GO_MINIO_APIHOST"),
		"accessKeyID":     os.Getenv("GO_MINIO_ACCESSKEYID"),
		"secretAccessKey": os.Getenv("GO_MINIO_SECRETACCESSKEY"),
	}

	for k, v := range s {
		if v == "" {
			return &s, fmt.Errorf("the value '%s' must not be empty", k)
		}
	}

	return &s, nil
}

func (s *Settings) GetEndpoint() string {
	return fmt.Sprintf("%s:%s", (*s)["apiHost"], (*s)["apiPort"])
}

func (s *Settings) GetAccessKeyID() string {
	return (*s)["accessKeyID"]
}

func (s *Settings) GetSecretAccessKey() string {
	return (*s)["secretAccessKey"]
}
