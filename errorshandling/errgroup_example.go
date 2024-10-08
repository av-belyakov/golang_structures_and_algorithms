// Шаблоны группировки произвольного количества ошибок
package errorshandling

import (
	"net/http"

	"golang.org/x/sync/errgroup"
)

// Circle некий цикл
type Circle struct {
	domain string
}

// Result некий результат
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

// ErrorGroup использует errgroup для группировки некоторого количества ошибок,
// также errgroup можно использовать с context.Context
// func handler(ctx context.Context, circles []Circle) ([]Result, error) {
func ErrorGroup(circles []Circle) ([]Result, error) {
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
