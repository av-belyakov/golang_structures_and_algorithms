package mispapi

import (
	"fmt"
	"net/url"
	"os"
)

func NewRequest(envMispTokenName, host string, port int) (*ModuleRequest, error) {
	module := &ModuleRequest{
		host: host,
		port: port,
	}

	authKey := os.Getenv(envMispTokenName)
	module.authKey = authKey

	return module, nil
}

// NewClientMISP клиент API MISP
func NewClientMISP(authKey, host string, port int, verify bool) (*ClientMISP, error) {
	urlBase, err := url.Parse(fmt.Sprintf("http://%s:%d/", host, port))
	if err != nil {
		return &ClientMISP{}, err
	}

	return &ClientMISP{
		BaseURL:  urlBase,
		Host:     host,
		AuthHash: authKey,
		Verify:   verify,
	}, nil
}
