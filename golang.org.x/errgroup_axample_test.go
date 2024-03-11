package golangorgx

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

//Использование errgroup обеспечивает синхронизацию для группы подпрограмм и
//предоставляет решение для работы с ошибками и общими контекстами.

type Circle struct {
	domain string
}

type Result struct {
	domain, status string
}

func sendReq(url string) (Result, error) {
	result := Result{domain: url}

	res, err := http.Get(url)
	if err != nil {
		return result, err
	}

	result.status = res.Status

	return result, nil
}

// также errgroup можно использовать с context.Context
// func handler(ctx context.Context, circles []Circle) ([]Result, error) {
func handler(circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	//g, ctx := errgroup.WithContext(ctx)
	g := errgroup.Group{}
	for i, elem := range circles {
		i := i
		domain := elem.domain
		g.Go(func() error {
			result, err := sendReq(domain)
			if err != nil {
				return err
			}

			results[i] = result

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func TestHandlerErrgroup(t *testing.T) {
	listCircle := []Circle{
		{domain: "http://www.golang.org/"},
		{domain: "http://www.google.com/"},
		{domain: "http://www.somestupidname.com/"},
		//{domain: "http://my.example.net/"},
	}

	// result, err := handler(context.Background(), listCircle)
	result, err := handler(listCircle)
	fmt.Println(result)

	//если при выполнение функции handler возникнет ошибка, например, из-за ошибок подключения
	// к несуществующему доменному имени то будет прекращена обработка всех выполняемых
	//горутин

	assert.NoError(t, err)
	assert.Equal(t, len(result), 3)
}
