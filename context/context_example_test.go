package context_test

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestContextExample(t *testing.T) {
	t.Run("Test 1. Example context.WithTimeout", func(t *testing.T) {
		t.Run("Test 1.1. Дождались отмены по таймауту", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
			defer cancel()

			log.Println("проверяем установлен ли таймаут в контексте")
			if deadline, ok := ctx.Deadline(); ok {
				log.Printf("да, таймаут в контексте установлен, deadline в '%s'\n", deadline.String())
				log.Printf("до истечении deadline осталось %s\n", time.Until(deadline).String())
			}

			log.Println("ждем отмены через 3 секунды")
			<-ctx.Done()
			log.Println("отмена спустя 3 секунды")
		})

		t.Run("Test 1.2. Отменяем контекст раньше чем истечет время", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)

			log.Println("ждем отмены через 3 секунды")
			cancel()

			<-ctx.Done()
			log.Println("отмена выполнена сразу")
		})
	})

	t.Run("Test 2. Example context.WithCancel", func(t *testing.T) {
		log.Println("каскадная отмена всей цепочки контекстов")

		ctx, cancel := context.WithCancel(t.Context())

		for i := range 3 {
			go func(ctx context.Context, i int) {
				newCtx, newCancel := context.WithCancel(ctx)
				defer newCancel()

				log.Printf("гороутина %d, ждём отмены\n", i)
				<-newCtx.Done()
				log.Printf("гороутина %d, отмена контекста\n", i)
			}(ctx, i)
		}

		time.Sleep(1 * time.Second)

		log.Println("отменяем контекст для всех гороутин")
		cancel()

		time.Sleep(3 * time.Second)
	})

	t.Run("Test 3. Example context.WithValue", func(t *testing.T) {
		t.Run("Test 3.1. Создаем контекст с ключом userIdType", func(t *testing.T) {
			type userIdType string
			userId := "12345"

			log.Printf("создаём контекст с ключём 'userIdType' равным '%s'\n", userId)

			ctx := context.WithValue(t.Context(), userIdType("userId"), userId)

			uid := ctx.Value(userIdType("userId")).(string)
			if uid == userId {
				log.Printf("получаем userId='%s' из контекста", uid)
			}
		})
	})

	t.Run("Test 4. Example context.WithoutCancel", func(t *testing.T) {
		type tokenType string
		token := "ydw78f8g3r7tr8r8ufguyedgf73fr73"

		t.Run("Test 4.1. Создаем контекст без отмены", func(t *testing.T) {
			ctx, cancel := context.WithCancel(t.Context())
			ctxVal := context.WithValue(ctx, tokenType("token"), token)

			ctxWoc := context.WithoutCancel(ctxVal)

			log.Println("отменяем контекст")
			cancel()

			ttoken := ctxWoc.Value(tokenType("token")).(string)

			if ttoken == token {
				log.Printf("успешно получаем токен '%s' даже после отмены контекста", ttoken)
			}
		})
	})

	t.Run("Test 5. Example context.AfterFunc", func(t *testing.T) {
		t.Run("Test 5.1. Контекст отменяется функция в AfterFunc не вызывается", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), time.Duration(time.Second*5))

			stop1 := context.AfterFunc(ctx, func() {
				log.Println("___ вызывается функция AfterFunc в Test 5.1")
			})

			log.Println("работа завершается до того как истекает context.WithTimeout, отменяем context.AfterFunc")

			stop1()
			cancel()
		})

		t.Run("Test 5.2. Контекст не отменяется функция в AfterFunc вызывается", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), time.Duration(time.Millisecond*5))
			defer cancel()

			context.AfterFunc(ctx, func() {
				log.Println("___ вызывается функция AfterFunc в Test 5.2")
			})

			log.Println("context.WithTimeout истекает до завершения работы, вызывается функция в context.AfterFunc")

			time.Sleep(1 * time.Second)
		})
	})
}
