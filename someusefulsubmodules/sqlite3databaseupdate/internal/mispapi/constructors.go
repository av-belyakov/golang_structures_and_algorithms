package mispapi

import (
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func NewRequest(host string, port int) (*ModuleRequest, error) {
	module := &ModuleRequest{
		host: host,
		port: port,
	}

	if err := godotenv.Load(".env"); err != nil {
		return module, err
	}

	authKey := os.Getenv("GO_SQLITE3DATABASEUPDATE_MISPTOKEN")
	module.authKey = authKey

	return module, nil
}

// NewClientMISP клиент API MISP
func NewClientMISP(authKey, host string, port int, verify bool) (*ClientMISP, error) {
	urlBase, err := url.Parse(fmt.Sprintf("http://%s:%d/", host, port))
	if err != nil {
		return &ClientMISP{}, err
	}

	fmt.Println("Port=", port)
	fmt.Printf("UrlBase:%+v\n", urlBase)

	return &ClientMISP{
		BaseURL:  urlBase,
		Host:     host,
		AuthHash: authKey,
		Verify:   verify,
	}, nil
}
