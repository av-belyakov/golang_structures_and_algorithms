package testingpattern

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//Тестовый пакет http (https://pkg.go.dev/net/http/httptest) предоставляет утилиты для
//тестирования HTTP как для клиентов, так и для серверов.

//Первый сценарий использования

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}

func GetIoReader(str ...string) *strings.Reader {
	return strings.NewReader(strings.Join(str, " "))
}

//Обработчик HTTP принимает два аргумента: запрос и способ записи ответа.
//Пакет httptest предоставляет утилиты для обоих вариантов. Для запроса мы можем использовать httptest.
//Новый запрос для создания *http.Request с использованием HTTP-метода, URL-адреса и тела запроса.
//Для получения ответа мы можем использовать httptest.NewRecorder для записи изменений,
//произведенных в обработчике.

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", strings.NewReader("foo"))
	w := httptest.NewRecorder()
	Handler(w, req)

	if got := w.Result().Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, got %s", got)
	}

	body, _ := io.ReadAll(GetIoReader("just test message"))

	if got := string(body); got != "just test message" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	if http.StatusOK != w.Result().StatusCode {
		t.FailNow()
	}
}

//Тестирование обработчика с помощью httptest не проверяет транспорт (часть HTTP).
//Целью теста является прямой вызов обработчика с запросом и способ записи ответа.
//Затем, используя программу записи ответов, мы записываем утверждения для проверки
//HTTP-заголовка, основного текста и кода состояния.

// Выполняем тестирование HTTP-клиента. Клиент отвечает за запрос конечной точки HTTP,
// который вычисляет, сколько времени требуется для перемещения от одной координаты к другой.

func TestHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Compare(string(greeting), "Hello, client") != 0 {
		log.Fatal("responses in not equal")
	}

	fmt.Printf("%s", greeting)
}
