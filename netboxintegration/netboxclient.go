package netboxintegration

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Settings настройки для подключения к Netbox
type Settings struct {
	token string
	host  string
	port  int
}

// Client клиент для работы с Netbox
type Client struct {
	client   *http.Client
	settings Settings
}

func NewClient(host string, port int, token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("the token must not be empty")
	}

	if host == "" {
		return nil, errors.New("the host must not be empty")
	}

	return &Client{
		settings: Settings{
			host:  host,
			port:  cmp.Or(port, 8005),
			token: token,
		},
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}, nil
}

// Get реализация HTTP GET запроса
func (api *Client) Get(ctx context.Context, query string) ([]byte, int, error) {
	url := fmt.Sprintf("http://%s:%d%s", api.settings.host, api.settings.port, query)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Add("Authorization", "Bearer nbt_JbhboktIck8d."+api.settings.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; indent=4")

	res, err := api.client.Do(req)
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("status code: %d (%s)", res.StatusCode, res.Status)
	}
	defer CloseHTTPResponse(res)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 500, err
	}

	return resBody, res.StatusCode, nil
}

// CloseHTTPResponse закрытие http соединения с предварительной проверкой
func CloseHTTPResponse(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}
